package markettokentest

import (
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
	ownerBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthFactory.From)

	if ownerBal.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected owner balance to be greater than 0, got %v", ownerBal)
	}
}

func TestMint(t *testing.T) {
	// only a privileged contract may call for mint, have a user pose as Listing
	_, privErr := deployed.MarketTokenContract.SetPrivileged(test.GetTxOpts(context.AuthFactory, nil, big.NewInt(test.ONE_GWEI*2),
		100000), context.AuthUser3.From, deployed.InvestingAddress)
	test.IfNotNil(t, privErr, "Error setting Market Token privileged contracts")

	context.Blockchain.Commit()

	supply, _ := deployed.MarketTokenContract.TotalSupply(nil)
	ownerBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthFactory.From)
	listBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser3.From) // our imposter
	// minting will add to the balances...
	_, mintErr := deployed.MarketTokenContract.Mint(test.GetTxOpts(context.AuthUser3, nil, big.NewInt(test.ONE_GWEI*2),
		100000), big.NewInt(test.ONE_WEI*2))
	test.IfNotNil(t, mintErr, "Error minting tokens")

	context.Blockchain.Commit()

	// our imposter
	expectedListBal := listBal.Add(listBal, big.NewInt(test.ONE_WEI*2))
	if listBalNow, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser3.From); listBalNow.Cmp(expectedListBal) != 0 {
		t.Errorf("Expected %v to be %v", expectedListBal, listBalNow)
	}

	expectedSupply := supply.Add(supply, big.NewInt(test.ONE_WEI*2))
	if newSupply, _ := deployed.MarketTokenContract.TotalSupply(nil); newSupply.Cmp(expectedSupply) != 0 {
		t.Errorf("Expected %v to be %v", newSupply, expectedSupply)
	}

	// owner bal does not increase in a mint operation
	if newOwnerBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthFactory.From); newOwnerBal.Cmp(ownerBal) != 0 {
		t.Errorf("Expected %v to be %v", newOwnerBal, ownerBal)
	}
}

func TestStopMinting(t *testing.T) {
	// only a privileged contract may call to stop minting TODO
}
