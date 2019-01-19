package market

import (
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
	// "time"
)

func calcPrice() *big.Int {
	rate, _ := deployed.ParameterizerContract.GetConversionRate(nil)
	slopeD, _ := deployed.ParameterizerContract.GetConversionSlopeDenominator(nil)
	slopeN, _ := deployed.ParameterizerContract.GetConversionSlopeNumerator(nil)
	reserve, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.MarketTokenAddress)
	// cuz golang big int math...
	var x big.Int
	y := x.Mul(slopeN, reserve)
	z := x.Add(rate, x.Div(y, slopeD))
	return z
}

func TestGetInvestmentPrice(t *testing.T) {
	// use the getListingHash func to make a faux dataHash
	// dataHash, _ := deployed.MarketContract.GetListingHash(nil, "FooData")

	price, _ := deployed.MarketContract.GetInvestmentPrice(nil)
	calc := calcPrice()

	if price.Cmp(calc) != 0 {
		t.Fatalf("Expected investment price to be %v, got: %v", calc, price)
	}
}

// TODO the invest method would require that a member have approved ether token.
func TestInvest(t *testing.T) {
	// member3 is not a listing owner
	// isOwner, _ := deployed.MarketContract.IsListingOwner(nil, context.AuthMember3.From)
	// if isOwner != false {
	// t.Fatalf("Expected isOwner to be false, got: %v", isOwner)
	// }

	// member 3 will need to deposit funds into the ether token
	// _, depErr := deployed.EtherTokenContract.Deposit(&bind.TransactOpts{
	// From:     context.AuthMember3.From,
	// Signer:   context.AuthMember3.Signer,
	// GasPrice: big.NewInt(ONE_GWEI * 2),
	// GasLimit: 1000000,
	// Value:    big.NewInt(ONE_WEI * 2),
	// }) // deposit 2 tokens

	// if depErr != nil {
	// t.Fatalf("Error transferring tokens to member: %v", depErr)
	// }

	// context.Blockchain.Commit()

	// allow the ether token to spend on user3's behalf
	// _, approveErr := deployed.EtherTokenContract.Approve(&bind.TransactOpts{
	// From:     context.AuthMember3.From,
	// Signer:   context.AuthMember3.Signer,
	// GasPrice: big.NewInt(ONE_GWEI * 2),
	// GasLimit: 1000000,
	// }, deployed.MarketAddress, big.NewInt(ONE_WEI*2)) // up to 2 tokens

	// if approveErr != nil {
	// t.Fatalf("Error approving market contract to spend: %v", approveErr)
	// }

	// context.Blockchain.Commit()

	// note our balances right now as investing will change them
	// memberBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember3.From)
	// marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.MarketAddress)
	// totalInvested, _ := deployed.MarketContract.GetInvested(nil)

	// remove
	// t.Log(memberBal)
	// t.Log(marketBal)

	// member 3 is not on council yet
	// inCouncil, _ := deployed.VotingContract.InCouncil(nil, context.AuthMember3.From)

	// if inCouncil != false {
	// t.Fatalf("Expected inCouncil to be false, got: %v", inCouncil)
	// }

	// member 3 also not an investor yet
	// isInvestor, _ := deployed.MarketContract.IsInvestor(nil, context.AuthMember3.From)

	// if isInvestor != false {
	// t.Fatalf("Expected isInvestor to be false, got: %v", isInvestor)
	// }

	// here we will just offer the suggested price
	// 	price, _ := deployed.MarketContract.GetInvestmentPrice(nil)
	// t.Logf("Investment price: %v", price)
	// what is user3's ether token bal?
	// member3Bal, _ := deployed.EtherTokenContract.BalanceOf(&bind.CallOpts{From: context.AuthMember3.From}, context.AuthMember3.From)
	// t.Logf("Authmember3 ether token balance: %v", member3Bal)
	// what is the allowance that ethertoken has for user3-market
	// member3Allowed, _ := deployed.EtherTokenContract.Allowance(nil, context.AuthMember3.From, deployed.MarketAddress)
	// t.Logf("Authmember3 allowance to market from ether token: %v", member3Allowed)

	// _, investErr := deployed.MarketContract.Invest(&bind.TransactOpts{
	// From:     context.AuthMember3.From,
	// Signer:   context.AuthMember3.Signer,
	// GasPrice: big.NewInt(ONE_GWEI * 2),
	// GasLimit: 1000000,
	// }, price)

	// if investErr != nil {
	// t.Fatalf("Error investing in market: %v", investErr)
	// }

	// context.Blockchain.Commit()

	// should be an investor now
	// isInvestorNow, _ := deployed.MarketContract.IsInvestor(nil, context.AuthMember3.From)

	// if isInvestorNow != true {
	// t.Fatalf("Expected isInvestor to be true, got: %v", isInvestorNow)
	// }

	// should be a council member now
	// inCouncilNow, _ := deployed.VotingContract.InCouncil(nil, context.AuthMember3.From)

	// if inCouncilNow != true {
	// t.Fatalf("Expected inCouncil to be true, got: %v", inCouncilNow)
	// }

	// inspect the investor
	// invested, minted, _ := deployed.MarketContract.GetInvestor(nil, context.AuthMember3.From)

	// if invested.Cmp(price) != 0 {
	// t.Fatalf("Expected %v to be %v", invested, price)
	// }

	// t.Log(invested)
	// t.Log(minted)
	// t.Log(totalInvested)
}
