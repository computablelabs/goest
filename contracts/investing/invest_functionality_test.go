package investing

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
	// "time"
)

func buyPrice() uint64 {
	rate, _ := deployed.ParameterizerContract.GetConversionRate(nil)
	rate64 := rate.Uint64()
	investD, _ := deployed.ParameterizerContract.GetInvestDenominator(nil)
	investD64 := investD.Uint64()
	investN, _ := deployed.ParameterizerContract.GetInvestNumerator(nil)
	investN64 := investN.Uint64()
	reserve, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.MarketTokenAddress)
	reserve64 := reserve.Uint64()
	tot, _ := deployed.MarketTokenContract.TotalSupply(nil)
	tot64 := tot.Uint64()

	if tot64 < 1000000000000000000 {
		return rate64 + investN64*reserve64/investD64
	} else {
		return rate64 + (investN64*reserve64*1000000000000000000)/(investD64*tot64)
	}
}

func sellPrice(addr common.Address) *big.Int {
	bal, _ := deployed.MarketTokenContract.BalanceOf(nil, addr)
	res, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.InvestingAddress)
	tot, _ := deployed.MarketTokenContract.TotalSupply(nil)
	var x big.Int
	y := x.Mul(bal, res)
	z := x.Div(y, tot)
	return z
}

func TestGetInvestmentPrice(t *testing.T) {
	price, _ := deployed.InvestingContract.GetInvestmentPrice(nil)
	price64 := price.Uint64()
	calc := buyPrice()

	if price64 != calc {
		t.Fatalf("Expected investment price to be %v, got: %v", calc, price64)
	}
}

func TestInvest(t *testing.T) {
	var x big.Int
	// member 3 will need to deposit funds into the ether token
	_, depErr := deployed.EtherTokenContract.Deposit(&bind.TransactOpts{
		From:     context.AuthMember3.From,
		Signer:   context.AuthMember3.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
		Value:    big.NewInt(ONE_WEI * 2),
	}) // deposit 2 tokens

	if depErr != nil {
		t.Fatalf("Error transferring tokens to member: %v", depErr)
	}

	context.Blockchain.Commit()

	// allow the ether token to spend on user3's behalf
	_, approveErr := deployed.EtherTokenContract.Approve(&bind.TransactOpts{
		From:     context.AuthMember3.From,
		Signer:   context.AuthMember3.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, deployed.InvestingAddress, big.NewInt(ONE_WEI*2)) // up to 2 tokens

	if approveErr != nil {
		t.Fatalf("Error approving market contract to spend: %v", approveErr)
	}

	context.Blockchain.Commit()

	// snapshot this user's market token bal as investing will increase by what was minted
	bal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember3.From)
	// here we will just offer the suggested price
	price, _ := deployed.InvestingContract.GetInvestmentPrice(nil)
	// t.Logf("Investment price: %v", price)

	// investing the asking price will always yield one market token gwei
	_, investErr := deployed.InvestingContract.Invest(&bind.TransactOpts{
		From:     context.AuthMember3.From,
		Signer:   context.AuthMember3.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, price.Mul(price, big.NewInt(10))) // 1e9 * 10

	if investErr != nil {
		t.Fatalf("Error investing in market: %v", investErr)
	}

	context.Blockchain.Commit()

	// if offering the price, you'll always get 1 gwei worth of market token
	// this will be reflected in the market token balance for this investor
	newBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember3.From)
	expectedBal := x.Add(bal, big.NewInt(ONE_GWEI*10)) // this is true in a market that had no prior reserve (like this spec)

	if newBal.Cmp(expectedBal) != 0 {
		t.Fatalf("Expected %v to be %v", newBal, expectedBal)
	}
}

func TestGetDivestmentProceeds(t *testing.T) {
	// user should have a non-zero market token balance
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember3.From)
	if marketBal.Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("Expected user to have a market token balance, got: %v", marketBal)
	}

	// generate a sale price and compare to the cantract method
	expectedProceeds := sellPrice(context.AuthMember3.From)
	proceeds, _ := deployed.InvestingContract.GetDivestmentProceeds(nil, context.AuthMember3.From)

	if proceeds.Cmp(expectedProceeds) != 0 {
		t.Fatalf("Expected %v to be %v", proceeds, expectedProceeds)
	}
}

// this is the divest-path-for-investor. for converted maker divest we will use the
// respective convert spec...
func TestDivest(t *testing.T) {
	// the total market balance atm
	tot, _ := deployed.MarketTokenContract.TotalSupply(nil)
	// user's ether token bal will go up by the sell price after divesting
	etherBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthMember3.From)
	expectedProceeds := sellPrice(context.AuthMember3.From)

	_, divestErr := deployed.InvestingContract.Divest(&bind.TransactOpts{
		From:     context.AuthMember3.From,
		Signer:   context.AuthMember3.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	})

	if divestErr != nil {
		t.Fatalf("Error divesting: %v", divestErr)
	}

	context.Blockchain.Commit()

	newTot, _ := deployed.MarketTokenContract.TotalSupply(nil)
	// user should have a non-zero market token balance
	newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember3.From)
	// user's ether token bal will go up by the sell price after divesting
	newEtherBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthMember3.From)

	if newTot.Cmp(tot) != -1 {
		t.Fatalf("Expected market total %v to be less than %v", newTot, tot)
	}

	if newMarketBal.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Expected users market token balance to be 0, got: %v", newMarketBal)
	}

	expectedEtherBal := etherBal.Add(etherBal, expectedProceeds)

	// the additional funds should match the expected sell price
	if newEtherBal.Cmp(expectedEtherBal) != 0 {
		t.Fatalf("Expected ether token balance %v to be %v", newEtherBal, expectedEtherBal)
	}
}
