package voting

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
	"time"
)

func TestVote(t *testing.T) {
	// auth member one will need some token
	memBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember1.From)
	if memBal.Cmp(big.NewInt(ONE_GWEI)) == -1 {
		// transfer one
		_, transErr := deployed.MarketTokenContract.Transfer(&bind.TransactOpts{
			From:     context.AuthFactory.From,
			Signer:   context.AuthFactory.Signer,
			GasPrice: big.NewInt(ONE_GWEI * 2),
			GasLimit: 1000000,
		}, context.AuthMember1.From, big.NewInt(ONE_GWEI))

		if transErr != nil {
			t.Fatalf("Error transferring tokens to member: %v", transErr)
		}

		context.Blockchain.Commit()
	}

	// member will need to have approved the voting contract to spend
	allowed, _ := deployed.MarketTokenContract.Allowance(nil, context.AuthMember1.From, deployed.VotingAddress)
	if !(allowed.Cmp(big.NewInt(ONE_GWEI)) >= 0) {
		_, approveErr := deployed.MarketTokenContract.Approve(&bind.TransactOpts{
			From:     context.AuthMember1.From,
			Signer:   context.AuthMember1.Signer,
			GasPrice: big.NewInt(ONE_GWEI * 2),
			GasLimit: 1000000,
		}, deployed.VotingAddress, big.NewInt(ONE_GWEI))

		if approveErr != nil {
			t.Fatalf("Error approving market contract to spend: %v", approveErr)
		}

		context.Blockchain.Commit()
	}

	// we need a candidate
	bytes := GenBytes32("iLoveListing.com")

	_, candidateErr := deployed.VotingContract.AddCandidate(&bind.TransactOpts{
		From:     context.AuthParameterizer.From,
		Signer:   context.AuthParameterizer.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 150000,
	}, bytes, APPLICATION, context.AuthMember1.From, big.NewInt(ONE_GWEI), big.NewInt(20)) // numbers smaller that 10 can be erratic?

	if candidateErr != nil {
		t.Fatalf("Error adding candidate: %v", candidateErr)
	}

	context.Blockchain.Commit()

	// should be no votes atm
	_, _, _, _, preYea, _, _ := deployed.VotingContract.GetCandidate(nil, bytes)

	if preYea.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Expected number of votes to be 0, got: %v", preYea)
	}

	// cast a yea vote
	_, voteErr := deployed.VotingContract.Vote(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 150000,
	}, bytes, big.NewInt(1))

	if voteErr != nil {
		t.Fatalf("Error voting for candidate: %v", voteErr)
	}

	context.Blockchain.Commit()

	// has been recorded in the candidate's votes
	_, _, _, _, yea, _, _ := deployed.VotingContract.GetCandidate(nil, bytes)

	if yea.Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("Expected number of votes to be > 0, got: %v", yea)
	}
}

func TestPollClosed(t *testing.T) {
	// should still be open
	bytes := GenBytes32("iLoveListing.com")
	closed, _ := deployed.VotingContract.PollClosed(nil, bytes)

	if closed != false {
		t.Fatalf("Expected pollClosed to be false, got: %v", closed)
	}

	// now move time forward so that voteBy has elapsed
	context.Blockchain.AdjustTime(20 * time.Second)
	context.Blockchain.Commit()

	updated, _ := deployed.VotingContract.PollClosed(nil, bytes)

	if updated != true {
		t.Fatalf("Expected pollClosed to be true, got: %v", updated)
	}
}

func TestDidPass(t *testing.T) {
	bytes := GenBytes32("iLoveListing.com")
	// the one total vote will do it
	passed, _ := deployed.VotingContract.DidPass(nil, bytes, big.NewInt(50))

	if passed != true {
		t.Fatalf("Expected didPass to be true, got: %v", passed)
	}
}
