package plcrvoting

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
)

func TestCreatePoll(t *testing.T) {
	t.Log("Voting contract should create a poll on demand")
	// voteQuorum, commitDuration, revealDuration are args to startPoll
	_, err := deployed.VotingContract.StartPoll(&bind.TransactOpts{
		From:   context.AuthOwner.From,
		Signer: context.AuthOwner.Signer,
	}, big.NewInt(51), big.NewInt(60), big.NewInt(60))

	if err != nil {
		t.Fatalf("Error starting poll: %v", err)
	}

	context.Blockchain.Commit()

	// we should have at least one poll now
	exists, _ := deployed.VotingContract.PollExists(&bind.CallOpts{
		From: context.AuthOwner.From,
	}, big.NewInt(1))

	if exists != true {
		t.Fatalf("Expected poll exists to be true, got : %v", err)
	}
}
