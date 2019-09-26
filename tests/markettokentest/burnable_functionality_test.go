package markettokentest

import (
	"github.com/computablelabs/goest/contracts/markettoken"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
)

func TestBurn(t *testing.T) {
	// Deploy custom Market Token
	mtBal := big.NewInt(test.ONE_ETH * 5)
	_, _, marketTokenCont, marketTokenErr := markettoken.DeployMarketToken(
		context.AuthOwner,
		context.Blockchain,
		context.AuthOwner.From,
		mtBal,
	)
	test.IfNotNil(t, marketTokenErr, "Error deploying markettoken")
	context.Blockchain.Commit()

	// MarketToken.Burn is only available to privileged contracts. Have user 3 pose as Listing...
	_, privErr := marketTokenCont.SetPrivileged(test.GetTxOpts(context.AuthOwner, nil, big.NewInt(test.ONE_GWEI*2),
		100000), deployed.ReserveAddress, context.AuthUser3.From)
	test.IfNotNil(t, privErr, "Error setting Market Token privileged contracts")

	context.Blockchain.Commit()

	// check for any balance first, minted token bal is held under the listing contract...
	bal, _ := marketTokenCont.BalanceOf(nil, context.AuthUser3.From)
	// minting will add to that balance
	_, mintErr := marketTokenCont.Mint(test.GetTxOpts(context.AuthUser3, nil, big.NewInt(test.ONE_GWEI*2), 100000),
		big.NewInt(test.ONE_ETH*4))
	test.IfNotNil(t, mintErr, "Error minting Market Token")

	context.Blockchain.Commit()

	newBal, _ := marketTokenCont.BalanceOf(nil, context.AuthUser3.From)
	if newBal.Cmp(bal) != 1 {
		t.Errorf("Expected %v to be > %v", newBal, bal)
	}

	// burn 2 and check balance
	_, burnErr := marketTokenCont.Burn(test.GetTxOpts(context.AuthUser3, nil, big.NewInt(test.ONE_GWEI*2),
		100000), big.NewInt(test.ONE_ETH*2))
	test.IfNotNil(t, burnErr, "Error burning Market Token")

	context.Blockchain.Commit()

	expectedBal := newBal.Sub(newBal, big.NewInt(test.ONE_ETH*2)) // burned 2
	newerBal, _ := marketTokenCont.BalanceOf(nil, context.AuthUser3.From)
	if newerBal.Cmp(expectedBal) != 0 {
		t.Errorf("Expected new balance to be %v, got: %v", expectedBal, newBal)
	}

	// Test that attempting to reset the listing privileged address doesn't work
	_, privErr2 := marketTokenCont.SetPrivileged(test.GetTxOpts(context.AuthOwner, nil, big.NewInt(test.ONE_GWEI*2),
		100000), deployed.ListingAddress, deployed.ReserveAddress)
	test.IfNotNil(t, privErr2, "Error attempting to reset Market Token privileged contracts")

	// Check that the listing address hasn't changed
	_, actualListingAddress, getErr := marketTokenCont.GetPrivileged(test.GetCallOpts(context.AuthOwner))
	test.IfNotNil(t, getErr, "Error attempting to get privileged contracts")
	if actualListingAddress != context.AuthUser3.From {
		t.Errorf("Listing address has been changed incorrectly")
	}
}
