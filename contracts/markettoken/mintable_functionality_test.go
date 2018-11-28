package markettoken

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestTotalSupply(t *testing.T) {
	t.Log("Market token should fetch supply on demand")
	// the supply should be truthy (return of 1 if > 0)
	if supply, _ := deployed.Contract.TotalSupply(nil); supply.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected total supply to be greater than 0, got %v", supply)
	}
}

func TestBalanceOf(t *testing.T) {
	t.Log("Market token should fetch the balance of a given address")

	user := common.HexToAddress("0xabc")

	// should have a 0 bal atm
	userBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: user}, user)

	if userBal.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected user balance of 0, got %v", userBal)
	}

	// owner has a truthy bal. TODO this will change after we implement the "initial token table" later...
	ownerBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthFactory.From}, context.AuthFactory.From)

	if ownerBal.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected owner balance to be greater than 0, got %v", ownerBal)
	}
}

// TODO test that calling mint from non-market address throws (or that it does nothing)

func TestMint(t *testing.T) {
	t.Log("Market token should mint on demand")

	// the starting supply at this point
	supply, _ := deployed.Contract.TotalSupply(nil)
	// owner's current token holdings TODO this check may change?
	ownerBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthFactory.From}, context.AuthFactory.From)
	marketBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthMarket.From}, context.AuthMarket.From)

	// minting will give market the minted amount and add to the total supply
	_, err := deployed.Contract.Mint(&bind.TransactOpts{
		From:     context.AuthMarket.From, // only the market can call for mint
		Signer:   context.AuthMarket.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 100000,
	}, big.NewInt(200))

	if err != nil {
		t.Fatalf("Error minting tokens: %v", err)
	}

	context.Blockchain.Commit()

	newMarketBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthMarket.From}, context.AuthMarket.From)
	expectedMarketBal := marketBal.Add(marketBal, big.NewInt(200)) // we just minted 200

	if newMarketBal.Cmp(expectedMarketBal) != 0 {
		t.Errorf("Expected market balance of %v, got %v", expectedMarketBal, newMarketBal)
	}

	// will have increased the total supply by the minted amount
	expectedSupply := supply.Add(supply, big.NewInt(200))

	if newSupply, _ := deployed.Contract.TotalSupply(nil); newSupply.Cmp(expectedSupply) != 0 {
		t.Errorf("Expected total supply to equal %v, got %v", expectedSupply, newSupply)
	}

	// total supply is increased but the balance of the initial holder did not - again, this may change TODO
	ownerBalCheck, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthFactory.From}, context.AuthFactory.From)
	if ownerBalCheck.Cmp(ownerBal) != 0 {
		t.Errorf("Expected owner balance of %v, got %v", ownerBal, ownerBalCheck)
	}
}

func TestStopMinting(t *testing.T) {
	t.Log("Market token can stop minting on demand (by market)")

	// the starting supply at this point
	supply, _ := deployed.Contract.TotalSupply(nil)

	// minting stopped is false by default
	if stopped, _ := deployed.Contract.MintingStopped(&bind.CallOpts{From: context.AuthMarket.From}); stopped != false {
		t.Errorf("Expected minting stopped to be false, got %v", stopped)
	}
	// stop and (re)check
	_, err := deployed.Contract.StopMinting(&bind.TransactOpts{
		From:     context.AuthMarket.From,
		Signer:   context.AuthMarket.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 100000,
	})

	if err != nil {
		t.Fatalf("Error stopping minting: %v", err)
	}

	context.Blockchain.Commit()

	if stopped, _ := deployed.Contract.MintingStopped(&bind.CallOpts{From: context.AuthFactory.From}); stopped != true {
		t.Errorf("Expected minting stopped to be true, got %v", stopped)
	}

	// no more minting can be done
	_, noMint := deployed.Contract.Mint(&bind.TransactOpts{
		From:     context.AuthMarket.From,
		Signer:   context.AuthMarket.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 100000,
	}, big.NewInt(100))

	if noMint != nil {
		t.Fatal("Error checking that no minting occurred")
	}

	// did not actually mint anything
	if newSupply, _ := deployed.Contract.TotalSupply(nil); newSupply.Cmp(supply) != 0 {
		t.Errorf("Expected total supply to remain %v, got %v", supply, newSupply)
	}
}
