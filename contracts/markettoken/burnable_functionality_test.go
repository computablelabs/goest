package markettoken

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
)

func TestBurn(t *testing.T) {
	// check for any market balance first...
	bal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthListing.From}, context.AuthListing.From)

	// mint 4 more and add to whatever bal was
	_, mintErr := deployed.Contract.Mint(&bind.TransactOpts{
		From:     context.AuthListing.From,
		Signer:   context.AuthListing.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 100000,
	}, big.NewInt(ONE_WEI*4))

	if mintErr != nil {
		t.Fatalf("Error minting tokens: %v", mintErr)
	}

	context.Blockchain.Commit()

	bal.Add(bal, big.NewInt(ONE_WEI*4)) // update the balance that we know about

	// burn 2 of them, check bal
	_, burnErr := deployed.Contract.Burn(&bind.TransactOpts{
		From:     context.AuthListing.From,
		Signer:   context.AuthListing.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 100000,
	}, big.NewInt(ONE_WEI*2))

	if burnErr != nil {
		t.Fatalf("Error burning tokens: %v", burnErr)
	}

	context.Blockchain.Commit()

	expectedBal := bal.Sub(bal, big.NewInt(ONE_WEI*2)) // we burned 2
	newBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthListing.From}, context.AuthListing.From)

	if newBal.Cmp(expectedBal) != 0 {
		t.Errorf("Expected market balance of %v, got %v", expectedBal, newBal)
	}
}
