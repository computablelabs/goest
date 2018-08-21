package basic

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestTotalSupply(t *testing.T) {
	t.Log("Basic token should fetch total supply on demand")
	// the supply should have been passed in as initial balance in the deploy spec
	if supply, _ := deployed.Contract.TotalSupply(nil); supply.Cmp(big.NewInt(10)) != 0 {
		t.Errorf("Expected total supply to equal initial balance, got %v", supply)
	}
}

func TestTransfer(t *testing.T) {
	t.Log("Basic token should transfer funds between addresses")

	// should be able to feed HexToAddress a short string and get back a full length address...
	user := common.HexToAddress("0xabc")

	// NOTE: if we want to view the transaction itself, it would be the first return arg
	_, err := deployed.Contract.Transfer(&bind.TransactOpts{
		From:   context.Auth.From,
		Signer: context.Auth.Signer,
		Value:  nil,
	}, user, big.NewInt(3))

	if err != nil {
		t.Fatalf("Error transferring funds to another account: %v", err)
	}

	context.Blockchain.Commit()
}

func TestBalanceOf(t *testing.T) {
	t.Log("Basic token should fetch the balance of a given address")

	user := common.HexToAddress("0xabc")

	// the user should have been transfered 100000 whatever
	userBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: user}, user)

	if userBal.Cmp(big.NewInt(3)) != 0 {
		t.Errorf("Expected user balance of 3, got %v", userBal)
	}

	// that 100000 should have been subtracted from the original owner, Auth.From in this case
	ownerBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.Auth.From}, context.Auth.From)

	if ownerBal.Cmp(big.NewInt(7)) != 0 {
		t.Errorf("Expected owner balance of 7, got %v", ownerBal)
	}
}
