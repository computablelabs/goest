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

	// cast a vote, one will suffice as we only have one council member here
	_, voteErr := deployed.VotingContract.Vote(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 100000,
	}, paramHash, big.NewInt(1))

	if voteErr != nil {
		t.Fatalf("Error voting for candidate: %v", voteErr)
	}

	context.Blockchain.Commit()

	// move time forward so the poll is closed
	context.Blockchain.AdjustTime(20 * time.Second)
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
