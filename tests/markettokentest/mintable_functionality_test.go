package markettokentest

import (
	"github.com/computablelabs/goest/contracts/markettoken"
	"github.com/computablelabs/goest/tests/test"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestTotalSupply(t *testing.T) {
	// the supply should be truthy (return of 1 if > 0)
	if supply, _ := deployed.MarketTokenContract.TotalSupply(nil); supply.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected total supply to be greater than 0, got %v", supply)
	}
}

func TestBalanceOf(t *testing.T) {
	user := common.HexToAddress("0xabc")

	// should have a 0 bal atm
	userBal, _ := deployed.MarketTokenContract.BalanceOf(nil, user)

	if userBal.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected user balance of 0, got %v", userBal)
	}

	// owner has a truthy bal.
	ownerBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthOwner.From)

	if ownerBal.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected owner balance to be greater than 0, got %v", ownerBal)
	}
}

func TestMint(t *testing.T) {
	// Market Token: { consumes: [none], privileged: [listing, reserve] }
	mtBal := big.NewInt(test.ONE_WEI * 5)
	_, _, marketTokenCont, marketTokenErr := markettoken.DeployMarketToken(
		context.AuthOwner,
		context.Blockchain,
		context.AuthOwner.From,
		mtBal,
	)
	test.IfNotNil(t, marketTokenErr, "Error deploying markettoken")
	context.Blockchain.Commit()

	// only a privileged contract may call for mint, have a user pose as Listing
	_, privErr := marketTokenCont.SetPrivileged(test.GetTxOpts(context.AuthOwner, nil, big.NewInt(test.ONE_GWEI*2),
		100000), context.AuthUser3.From, deployed.ReserveAddress)
	test.IfNotNil(t, privErr, "Error setting Market Token privileged contracts")

	context.Blockchain.Commit()

	supply, _ := marketTokenCont.TotalSupply(nil)
	ownerBal, _ := marketTokenCont.BalanceOf(nil, context.AuthOwner.From)
	listBal, _ := marketTokenCont.BalanceOf(nil, context.AuthUser3.From) // our imposter
	// minting will add to the balances...
	_, mintErr := marketTokenCont.Mint(test.GetTxOpts(context.AuthUser3, nil, big.NewInt(test.ONE_GWEI*2),
		100000), big.NewInt(test.ONE_WEI*2))
	test.IfNotNil(t, mintErr, "Error minting tokens")

	context.Blockchain.Commit()

	// our imposter
	expectedListBal := listBal.Add(listBal, big.NewInt(test.ONE_WEI*2))
	if listBalNow, _ := marketTokenCont.BalanceOf(nil, context.AuthUser3.From); listBalNow.Cmp(expectedListBal) != 0 {
		t.Errorf("Expected %v to be %v", expectedListBal, listBalNow)
	}

	expectedSupply := supply.Add(supply, big.NewInt(test.ONE_WEI*2))
	if newSupply, _ := marketTokenCont.TotalSupply(nil); newSupply.Cmp(expectedSupply) != 0 {
		t.Errorf("Expected %v to be %v", newSupply, expectedSupply)
	}

	// owner bal does not increase in a mint operation
	if newOwnerBal, _ := marketTokenCont.BalanceOf(nil, context.AuthOwner.From); newOwnerBal.Cmp(ownerBal) != 0 {
		t.Errorf("Expected %v to be %v", newOwnerBal, ownerBal)
	}
}
