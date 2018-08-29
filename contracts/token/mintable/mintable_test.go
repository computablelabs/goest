package mintable

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestTotalSupply(t *testing.T) {
	t.Log("Mintable token should fetch total supply on demand")
	// the supply should have been passed in as initial balance in the deploy spec
	if supply, _ := deployed.Contract.TotalSupply(nil); supply.Cmp(big.NewInt(1000)) != 0 {
		t.Errorf("Expected total supply to equal initial balance, got %v", supply)
	}
}

func TestBalanceOf(t *testing.T) {
	t.Log("Mintable token should fetch the balance of a given address")

	user := common.HexToAddress("0xabc")

	// should have a 0 bal atm
	userBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: user}, user)

	if userBal.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected user balance of 0, got %v", userBal)
	}

	// owner has all the initial bal
	ownerBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.Auth.From}, context.Auth.From)

	if ownerBal.Cmp(big.NewInt(1000)) != 0 {
		t.Errorf("Expected owner balance of 1000, got %v", ownerBal)
	}
}

func TestMint(t *testing.T) {
	t.Log("Mintable token should mint on demand")

	user := common.HexToAddress("0xabc")

	// minting will give user the minted amount
	_, err := deployed.Contract.Mint(&bind.TransactOpts{
		From:     context.Auth.From,
		Signer:   context.Auth.Signer,
		GasLimit: 200000,
		Value:    nil,
	}, user, big.NewInt(200))

	if err != nil {
		t.Fatalf("Error minting tokens: %v", err)
	}

	context.Blockchain.Commit()

	userBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: user}, user)

	if userBal.Cmp(big.NewInt(200)) != 0 {
		t.Errorf("Expected user balance of 200, got %v", userBal)
	}

	// will have increased the total supply by the minted amount
	if supply, _ := deployed.Contract.TotalSupply(nil); supply.Cmp(big.NewInt(1200)) != 0 {
		t.Errorf("Expected total supply to equal initial balance + minted, got %v", supply)
	}
}
