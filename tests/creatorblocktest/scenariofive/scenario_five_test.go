package scenariofive

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
	//"time"
)

func oneHundredOneEth() *big.Int {
	one := big.NewInt(test.ONE_WEI)
	oneHundred := one.Mul(one, big.NewInt(100))
	return oneHundred.Add(oneHundred, big.NewInt(test.ONE_WEI))
}

func oneHundredEth() *big.Int {
	one := big.NewInt(test.ONE_WEI)
	return one.Mul(one, big.NewInt(100))
}

func TestFullSimulation(t *testing.T) {
	// The creator of the market starts by transfering 100 ETH to reserve
	_, transError := deployed.EtherTokenContract.Transfer(test.GetTxOpts(
		context.AuthOwner, nil, big.NewInt(test.ONE_GWEI*2), 100000),
		deployed.InvestingAddress, oneHundredEth())
	test.IfNotNil(t, transError, "Error transferring token")
	context.Blockchain.Commit()

	ownerEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthOwner.From)
	resEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.InvestingAddress)

	// has 1 ETH left
	if ownerEthBal.Cmp(big.NewInt(test.ONE_WEI)) != 0 {
		t.Errorf("Expected ether token balance of 1 ETH, got: %v", ownerEthBal)
	}

	// got the 100
	if resEthBal.Cmp(oneHundredEth()) != 0 {
		t.Errorf("Expected reserve of 100 Eth, got: %v", resEthBal)
	}

	t.Logf("Current Reserve balance: %v", test.Commafy(resEthBal))
	// Original market token supply
	mtSup, _ := deployed.MarketTokenContract.TotalSupply(nil)
	t.Logf("Original Market token total supply post creation: %v", test.Commafy(mtSup))

	t.Logf("Number Makers: %d", len(makers))
	// Let's have the maker submit a listing
	for name, maker := range makers {
		t.Logf("Submitting listing for %s", name)
		listingHash := test.GenBytes32(name)

		_, listErr := deployed.ListingContract.List(test.GetTxOpts(maker, nil,
			big.NewInt(test.ONE_GWEI*2), 250000), listingHash)
		test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
		context.Blockchain.Commit()
	}
}
