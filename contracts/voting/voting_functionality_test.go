package voting

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
	"time"
)

func TestVote(t *testing.T) {
	t.Log("Voting contract should allow a council member to vote")

	// we need a council member
	_, councilErr := deployed.VotingContract.AddToCouncil(&bind.TransactOpts{
		From:     context.AuthMarket.From,
		Signer:   context.AuthMarket.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 100000,
	}, context.AuthMember1.From)

	if councilErr != nil {
		t.Fatalf("Error adding council member: %v", councilErr)
	}

	// we need a candidate
	bytes := genBytes32("iLoveListing.com")

	_, candidateErr := deployed.VotingContract.AddCandidate(&bind.TransactOpts{
		From:     context.AuthMarket.From,
		Signer:   context.AuthMarket.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 150000,
	}, "application", bytes, big.NewInt(20)) // numbers smaller that 10 can be erratic?

	if candidateErr != nil {
		t.Fatalf("Error adding candidate: %v", candidateErr)
	}

	context.Blockchain.Commit()

	// cast a vote
	_, voteErr := deployed.VotingContract.Vote(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 100000,
	}, bytes)

	if voteErr != nil {
		t.Fatalf("Error voting for candidate: %v", voteErr)
	}

	context.Blockchain.Commit()

	// has been recorded in the candidate's votes
	_, _, votes, _ := deployed.VotingContract.GetCandidate(nil, bytes)

	if votes.Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("Expected number of votes to be > 0, got: %v", votes)
	}
}

func TestDidVote(t *testing.T) {
	t.Log("Voting contract has recorded that a member cast a vote")

	bytes := genBytes32("iLoveListing.com")
	voted, _ := deployed.VotingContract.DidVote(nil, bytes, context.AuthMember1.From)

	if voted != true {
		t.Fatalf("Expected member didVote to be true, got: %v", voted)
	}
}

func TestPollClosed(t *testing.T) {
	t.Log("Voting contract correcly identifies that polls are closed")
	// should still be open
	bytes := genBytes32("iLoveListing.com")
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
	t.Log("Voting contract correctly states if a cadidate passed a vote")

	bytes := genBytes32("iLoveListing.com")
	// with only one council member, the one vote will do it
	passed, _ := deployed.VotingContract.DidPass(nil, bytes, big.NewInt(50))

	if passed != true {
		t.Fatalf("Expected didPass to be true, got: %v", passed)
	}
}
