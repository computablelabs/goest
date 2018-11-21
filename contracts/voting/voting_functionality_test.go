package voting

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

// we'll include council functionality in here first
func TestAddToCouncil(t *testing.T) {
	t.Log("Voting contract should add a council member")

	member := common.HexToAddress("0xfgh")

	_, err := deployed.VotingContract.AddToCouncil(&bind.TransactOpts{
		From:     context.AuthMarket.From,
		Signer:   context.AuthMarket.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 100000,
	}, member)

	if err != nil {
		t.Fatalf("Error adding council member: %v", err)
	}

	context.Blockchain.Commit()
}

func TestInCouncil(t *testing.T) {
	t.Log("Voting contract should identify a council member")

	member := common.HexToAddress("0xfgh")
	res, _ := deployed.VotingContract.InCouncil(nil, member)

	if res != true {
		t.Fatalf("Expected inCouncil call to be true, got: %v", res)
	}
}

func TestRemoveFromCouncil(t *testing.T) {

}
