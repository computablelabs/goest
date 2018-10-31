package markettoken

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
)

func TestBurn(t *testing.T) {
	t.Log("Market token contract should burn tokens on demand")
	// check for any market balance first...
	bal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthMarket.From}, context.AuthMarket.From)

	t.Logf("Market token balance, %v", bal)

	// mint 10 more and add to whatever bal was
	_, mintErr := deployed.Contract.Mint(&bind.TransactOpts{
		From:     context.AuthMarket.From,
		Signer:   context.AuthMarket.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 100000,
	}, big.NewInt(10))

	if mintErr != nil {
		t.Fatalf("Error minting tokens: %v", mintErr)
	}

	context.Blockchain.Commit()

	bal.Add(bal, big.NewInt(10)) // update the balance that we know about

	// burn 5 of them, check bal
	_, burnErr := deployed.Contract.Burn(&bind.TransactOpts{
		From:     context.AuthMarket.From,
		Signer:   context.AuthMarket.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 100000,
	}, big.NewInt(5))

	if burnErr != nil {
		t.Fatalf("Error burning tokens: %v", burnErr)
	}

	context.Blockchain.Commit()

	expectedBal := bal.Sub(bal, big.NewInt(5)) // we burned 5
	newBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthMarket.From}, context.AuthMarket.From)

	if newBal.Cmp(expectedBal) != 0 {
		t.Errorf("Expected market balance of %v, got %v", expectedBal, newBal)
	}
}
