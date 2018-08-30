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

	// total supply is increased but the balance of the initial holder did not
	ownerBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.Auth.From}, context.Auth.From)

	if ownerBal.Cmp(big.NewInt(1000)) != 0 {
		t.Errorf("Expected owner balance of 1000, got %v", ownerBal)
	}
}

func TestStopMinting(t *testing.T) {
	t.Log("Mintable token can stop minting on demand (by owner)")
	// minting stopped is false by default
	if stopped, _ := deployed.Contract.MintingStopped(&bind.CallOpts{From: context.Auth.From}); stopped != false {
		t.Errorf("Expected minting stopped to be false, got %v", stopped)
	}
	// stop and (re)check
	_, err := deployed.Contract.StopMinting(&bind.TransactOpts{
		From:   context.Auth.From,
		Signer: context.Auth.Signer,
	})

	if err != nil {
		t.Fatalf("Error stopping minting: %v", err)
	}

	context.Blockchain.Commit()

	if stopped, _ := deployed.Contract.MintingStopped(&bind.CallOpts{From: context.Auth.From}); stopped != true {
		t.Errorf("Expected minting stopped to be true, got %v", stopped)
	}

	// no more minting can be done
	_, noMint := deployed.Contract.Mint(&bind.TransactOpts{
		From:     context.Auth.From,
		Signer:   context.Auth.Signer,
		GasLimit: 200000,
	}, common.HexToAddress("0xdef"), big.NewInt(100))

	if noMint != nil {
		t.Fatal("Error checking that no minting occurred")
	}

	// did not actually mint anything
	if supply, _ := deployed.Contract.TotalSupply(nil); supply.Cmp(big.NewInt(1200)) != 0 {
		t.Errorf("Expected total supply to remain 1200, got %v", supply)
	}
}
