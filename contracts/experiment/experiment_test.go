package experiment

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
)

func TestGetInvestmentPrice(t *testing.T) {
	// reserve is 0 atm, so the price should == the default conversion rate of 1e9
	price, _ := deployed.Contract.GetInvestmentPrice(nil)
	expectedPrice := big.NewInt(1e9)
	t.Logf("Investment price with 0 reserve (is the conversion rate): %v", price)

	if price.Cmp(expectedPrice) != 0 {
		t.Fatalf("Expected %v to be %v", price, expectedPrice)
	}

	// let's say some stuff has happened and 1 token in wei has accumulated in reserve
	_, setResErr := deployed.Contract.SetReserve(&bind.TransactOpts{
		From:     context.AuthFactory.From,
		Signer:   context.AuthFactory.Signer,
		GasPrice: big.NewInt(1e9 * 2),
		GasLimit: 100000,
	}, big.NewInt(1e18))

	if setResErr != nil {
		t.Fatalf("Error setting reserve balance: %v", setResErr)
	}

	context.Blockchain.Commit()

	// reserve should be truthy now
	reserve, _ := deployed.Contract.GetReserve(nil)
	expectedReserve := big.NewInt(1e18)

	if reserve.Cmp(expectedReserve) != 0 {
		t.Fatalf("Expected %v to be %v", price, expectedReserve)
	}

	newPrice, _ := deployed.Contract.GetInvestmentPrice(nil)
	t.Logf("Investment price with a 1 ether reserve balance: %v", newPrice)

	// the new price will be higher than the old one, use log above to see what it is
	if newPrice.Cmp(expectedPrice) != 1 {
		t.Fatalf("Expected %v to be %v", newPrice, expectedPrice)
	}
}
