package parameterizer

import (
	// "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	// "github.com/ethereum/go-ethereum/crypto/sha3"
	"math/big"
	"testing"
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

	inCouncil, _ := deployed.VotingContract.InCouncil(nil, context.AuthVoter.From)

	if inCouncil != true {
		t.Fatalf("Expected voter to be council member")
	}

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

func TestResolveReparam(t *testing.T) {
	// the old voteby
}
