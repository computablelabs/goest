package parameterizer

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
	"time"
)

func TestParameterize(t *testing.T) {
	t.Log("Parameterizer can create a new reparam proposal")

	// we need a council member
	_, councilErr := deployed.VotingContract.AddToCouncil(&bind.TransactOpts{
		From:     context.AuthMarket.From,
		Signer:   context.AuthMarket.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 100000,
	}, context.AuthVoter.From)

	if councilErr != nil {
		t.Fatalf("Error adding council member: %v", councilErr)
	}

	context.Blockchain.Commit()

	_, err := deployed.ParameterizerContract.Reparameterize(&bind.TransactOpts{
		From:     context.AuthVoter.From,
		Signer:   context.AuthVoter.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 200000,
	}, "voteBy", big.NewInt(25))

	if err != nil {
		t.Fatalf("Error creating proposal: %v", err)
	}

	context.Blockchain.Commit()

	// we should see a candidate now
	candidates, _ := deployed.VotingContract.GetCandidates(nil)

	if len(candidates) < 1 {
		t.Fatalf("Expected number of candidates to be >= 1, got: %v", len(candidates))
	}
}

func TestGetReparam(t *testing.T) {
	paramHash, _ := deployed.ParameterizerContract.GetParamHash(nil, "voteBy", big.NewInt(25))
	proposer, name, val, _ := deployed.ParameterizerContract.GetReparam(nil, paramHash)

	if proposer != context.AuthVoter.From {
		t.Fatalf("Expected proposer to be %v, got: %v", context.AuthVoter.From, proposer)
	}

	if name != "voteBy" {
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
	paramHash, _ := deployed.ParameterizerContract.GetParamHash(nil, "voteBy", big.NewInt(25))

	// cast a vote, one will suffice as we only have one council member here
	_, voteErr := deployed.VotingContract.Vote(&bind.TransactOpts{
		From:     context.AuthVoter.From,
		Signer:   context.AuthVoter.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 100000,
	}, paramHash)

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
		From:     context.AuthVoter.From,
		Signer:   context.AuthVoter.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
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
	candidates, _ := deployed.VotingContract.GetCandidates(nil)

	if len(candidates) != 0 {
		t.Fatalf("Expected no candidates, got: %v", len(candidates))
	}

	// should have cleaned up the proposal
	proposer, name, val, _ := deployed.ParameterizerContract.GetReparam(nil, paramHash)

	if proposer != common.HexToAddress("0x0") {
		t.Fatalf("Expected proposer to be falsy, got: %v", proposer)
	}

	if name != "" {
		t.Fatalf("Expected name to be falsy, got: %v", name)
	}

	if val.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Expected value to be 0, got: %v", val)
	}
}
