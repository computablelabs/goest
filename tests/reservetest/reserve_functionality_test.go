package reservetest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func buyPrice() uint64 {
	floor, _ := deployed.ParameterizerContract.GetPriceFloor(nil)
	floor64 := floor.Uint64()
	var d64 uint64 = 100
	spread, _ := deployed.ParameterizerContract.GetSpread(nil)
	spread64 := spread.Uint64()
	reserve, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.MarketTokenAddress)
	reserve64 := reserve.Uint64()
	tot, _ := deployed.MarketTokenContract.TotalSupply(nil)
	tot64 := tot.Uint64()

	if tot64 < 1000000000000000000 {
		return floor64 + ((spread64 * reserve64) / d64)
	} else {
		return floor64 + ((spread64 * reserve64 * 1000000000000000000) / (d64 * tot64))
	}
}

func sellPrice(addr common.Address) *big.Int {
	bal, _ := deployed.MarketTokenContract.BalanceOf(nil, addr)
	res, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.ReserveAddress)
	tot, _ := deployed.MarketTokenContract.TotalSupply(nil)
	var x big.Int
	y := x.Mul(bal, res)
	z := x.Div(y, tot)
	return z
}

func TestGetSupportPrice(t *testing.T) {
	price, _ := deployed.ReserveContract.GetSupportPrice(nil)
	price64 := price.Uint64()
	calc := buyPrice()

	if price64 != calc {
		t.Errorf("Expected support price to be %v, got: %v", calc, price64)
	}
}

func TestSupport(t *testing.T) {
	var x big.Int
	// member 3 will need to deposit funds into the ether token
	_, depErr := deployed.EtherTokenContract.Deposit(test.GetTxOpts(context.AuthUser3,
		big.NewInt(test.ONE_ETH*2), big.NewInt(test.ONE_GWEI*2), 1000000)) // deposit 2 tokens
	test.IfNotNil(t, depErr, fmt.Sprintf("Error depositing to ethertoken: %v", depErr))
	context.Blockchain.Commit()

	// allow the ether token to spend on user3's behalf
	_, approveErr := deployed.EtherTokenContract.Approve(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), deployed.ReserveAddress, big.NewInt(test.ONE_ETH*2)) // up to 2 tokens
	test.IfNotNil(t, approveErr, fmt.Sprintf("Error approving market contract to spend: %v", approveErr))
	context.Blockchain.Commit()

	// snapshot this user's market token bal as reserve will increase by what was minted
	bal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser3.From)
	// here we will just offer the suggested price
	price, _ := deployed.ReserveContract.GetSupportPrice(nil)

	_, sErr := deployed.ReserveContract.Support(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), price.Mul(price, big.NewInt(10))) // 1e9 * 10
	test.IfNotNil(t, sErr, fmt.Sprintf("Error supporting market: %v", sErr))
	context.Blockchain.Commit()

	// if offering the price, you'll always get 1 gwei worth of market token
	// this will be reflected in the market token balance for this supporter
	newBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser3.From)
	expectedBal := x.Add(bal, big.NewInt(test.ONE_GWEI*10)) // this is true in a market that had no prior reserve (like this spec)

	if newBal.Cmp(expectedBal) != 0 {
		t.Errorf("Expected %v to be %v", newBal, expectedBal)
	}
}

func TestGetWithdrawalProceeds(t *testing.T) {
	// user should have a non-zero market token balance
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser3.From)
	if marketBal.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected user to have a market token balance, got: %v", marketBal)
	}

	// generate a sale price and compare to the contract method
	expectedProceeds := sellPrice(context.AuthUser3.From)
	proceeds, _ := deployed.ReserveContract.GetWithdrawalProceeds(nil, context.AuthUser3.From)

	if proceeds.Cmp(expectedProceeds) != 0 {
		t.Errorf("Expected %v to be %v", proceeds, expectedProceeds)
	}
}

// this is the withdrawal-path-for-supporter
func TestWithdraw(t *testing.T) {
	// the total market balance atm
	tot, _ := deployed.MarketTokenContract.TotalSupply(nil)
	// user's ether token bal will go up by the sell price after withdrawal
	etherBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser3.From)
	expectedProceeds := sellPrice(context.AuthUser3.From)

	_, wErr := deployed.ReserveContract.Withdraw(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000))
	test.IfNotNil(t, wErr, fmt.Sprintf("Error withdrawing: %v", wErr))
	context.Blockchain.Commit()

	newTot, _ := deployed.MarketTokenContract.TotalSupply(nil)
	// user should have a non-zero market token balance
	newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser3.From)
	newEtherBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser3.From)

	if newTot.Cmp(tot) != -1 {
		t.Errorf("Expected market total %v to be less than %v", newTot, tot)
	}

	if newMarketBal.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected users market token balance to be 0, got: %v", newMarketBal)
	}

	expectedEtherBal := etherBal.Add(etherBal, expectedProceeds)

	// the additional funds should match the expected sell price
	if newEtherBal.Cmp(expectedEtherBal) != 0 {
		t.Errorf("Expected ether token balance %v to be %v", newEtherBal, expectedEtherBal)
	}
}
