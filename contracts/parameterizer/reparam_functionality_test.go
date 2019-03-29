package parameterizer

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
	"time"
)

func TestParameterize(t *testing.T) {
	_, err := deployed.ParameterizerContract.Reparameterize(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 200000,
	}, VOTE_BY, big.NewInt(25))

	if err != nil {
		t.Fatalf("Error creating proposal: %v", err)
	}

	context.Blockchain.Commit()

	// we should see a candidate now
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(25))
	isCan, _ := deployed.VotingContract.IsCandidate(nil, paramHash)

	if !isCan {
		t.Fatalf("Expected isCandidate to be true, got: %v", isCan)
	}
}

func TestGetReparam(t *testing.T) {
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(25))
	name, val, _ := deployed.ParameterizerContract.GetReparam(nil, paramHash)

	if name.Cmp(VOTE_BY) != 0 {
		t.Fatalf("Expected name to be 'voteBy', got: %v", name)
	}

	if val.Cmp(big.NewInt(25)) != 0 {
		t.Fatalf("Expected value to be 25, got: %v", val)
	}
}

func TestResolveReparam(t *testing.T) {
	oldVoteBy, _ := deployed.ParameterizerContract.GetVoteBy(nil)

	// make sure its not the proposed candidate's amount (25)
	if oldVoteBy.Cmp(big.NewInt(25)) == 0 {
		t.Fatalf("Expected the old voteBy param to not be 25, got: %v", oldVoteBy)
	}

	// how we will fetch this proposal
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(25))

	// cast a vote, member may need funding...
	memBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember3.From)
	if memBal.Cmp(big.NewInt(ONE_GWEI)) == -1 {
		// transfer one
		_, transErr := deployed.MarketTokenContract.Transfer(&bind.TransactOpts{
			From:     context.AuthFactory.From,
			Signer:   context.AuthFactory.Signer,
			GasPrice: big.NewInt(ONE_GWEI * 2),
			GasLimit: 1000000,
		}, context.AuthMember3.From, big.NewInt(ONE_GWEI))

		if transErr != nil {
			t.Fatalf("Error transferring tokens to member: %v", transErr)
		}

		context.Blockchain.Commit()
	}

	// member will need to have approved the voting contract to spend
	allowed, _ := deployed.MarketTokenContract.Allowance(nil, context.AuthMember3.From, deployed.VotingAddress)
	if !(allowed.Cmp(big.NewInt(ONE_GWEI)) >= 0) {
		_, approveErr := deployed.MarketTokenContract.Approve(&bind.TransactOpts{
			From:     context.AuthMember3.From,
			Signer:   context.AuthMember3.Signer,
			GasPrice: big.NewInt(ONE_GWEI * 2),
			GasLimit: 1000000,
		}, deployed.VotingAddress, big.NewInt(ONE_GWEI))

		if approveErr != nil {
			t.Fatalf("Error approving market contract to spend: %v", approveErr)
		}

		context.Blockchain.Commit()
	}

	_, voteErr := deployed.VotingContract.Vote(&bind.TransactOpts{
		From:     context.AuthMember3.From,
		Signer:   context.AuthMember3.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 150000,
	}, paramHash, big.NewInt(1))

	if voteErr != nil {
		t.Fatalf("Error voting for candidate: %v", voteErr)
	}

	context.Blockchain.Commit()

	_, _, _, _, yea, _, _ := deployed.VotingContract.GetCandidate(nil, paramHash)
	t.Log(yea)

	if yea.Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("Expected number of votes to be > 0, got: %v", yea)
	}

	// move time forward so the poll is closed
	context.Blockchain.AdjustTime(40 * time.Second)
	context.Blockchain.Commit()

	// make sure its closed now
	closed, _ := deployed.VotingContract.PollClosed(nil, paramHash)

	if closed != true {
		t.Fatalf("Expected pollClosed to be true, got: %v", closed)
	}

	// resolve it
	_, resolveErr := deployed.ParameterizerContract.ResolveReparam(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 150000,
	}, paramHash)

	if resolveErr != nil {
		t.Fatalf("Error resolving proposal: %v", resolveErr)
	}

	context.Blockchain.Commit()

	// voteBy should be changed now
	newVoteBy, _ := deployed.ParameterizerContract.GetVoteBy(nil)

	if newVoteBy.Cmp(big.NewInt(25)) != 0 {
		t.Fatalf("Expected the new voteBy param to be 25, got: %v", newVoteBy)
	}

	// should have cleaned up the voting candidate
	isCanNow, _ := deployed.VotingContract.IsCandidate(nil, paramHash)

	if isCanNow {
		t.Fatalf("Expected isCandidate to be false, got: %v", isCanNow)
	}

	// should have cleaned up the proposal
	name, val, _ := deployed.ParameterizerContract.GetReparam(nil, paramHash)

	if name.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Expected name to be falsy, got: %v", name)
	}

	if val.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Expected value to be 0, got: %v", val)
	}
}
