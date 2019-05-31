package markettokentest

import (
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
)

func TestBurn(t *testing.T) {
	// MarketToken.Burn is only available to privileged contracts. Have user 3 pose as Listing...
	_, privErr := deployed.MarketTokenContract.SetPrivileged(test.GetTxOpts(context.AuthOwner, nil, big.NewInt(test.ONE_GWEI*2),
		100000), context.AuthUser3.From, deployed.ReserveAddress)
	test.IfNotNil(t, privErr, "Error setting Market Token privileged contracts")

	context.Blockchain.Commit()

	// check for any balance first, minted token bal is held under the listing contract...
	bal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser3.From)
	// minting will add to that balance
	_, mintErr := deployed.MarketTokenContract.Mint(test.GetTxOpts(context.AuthUser3, nil, big.NewInt(test.ONE_GWEI*2), 100000),
		big.NewInt(test.ONE_WEI*4))
	test.IfNotNil(t, mintErr, "Error minting Market Token")

	context.Blockchain.Commit()

	// burn 2 and check balance
	_, burnErr := deployed.MarketTokenContract.Burn(test.GetTxOpts(context.AuthUser3, nil, big.NewInt(test.ONE_GWEI*2),
		100000), big.NewInt(test.ONE_WEI*2))
	test.IfNotNil(t, burnErr, "Error burning Market Token")

	context.Blockchain.Commit()

	expectedBal := bal.Add(bal, big.NewInt(test.ONE_WEI*2)) // minted 4, burned 2
	newBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser3.From)
	if newBal.Cmp(expectedBal) != 0 {
		t.Errorf("Expected new balance to be %v, got: %v", expectedBal, newBal)
	}

	// reset the actual listing privileged address
	_, privErr2 := deployed.MarketTokenContract.SetPrivileged(test.GetTxOpts(context.AuthOwner, nil, big.NewInt(test.ONE_GWEI*2),
		100000), deployed.ListingAddress, deployed.ReserveAddress)
	test.IfNotNil(t, privErr2, "Error resetting Market Token privileged contracts")
}
