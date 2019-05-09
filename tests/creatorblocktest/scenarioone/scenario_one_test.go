package scenarioone

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
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

func tenEth() *big.Int {
	one := big.NewInt(test.ONE_WEI)
	return one.Mul(one, big.NewInt(10))
}

func TestInitialBalance(t *testing.T) {
	// owner has all the eth atm
	etBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthOwner.From)
	// all market tokens that exist atm
	mtSup, _ := deployed.MarketTokenContract.TotalSupply(nil)

	if etBal.Cmp(oneHundredOneEth()) != 0 {
		t.Errorf("Expected ether token balance of %v, got: %v", oneHundredOneEth(), etBal)
	}
	if mtSup.Cmp(big.NewInt(test.ONE_WEI)) != 0 {
		t.Errorf("Expected market token supply of 1wei, got: %v", mtSup)
	}

	t.Logf("Current Market Token total supply: %v", test.Commafy(mtSup))
}

func TestTransferToReserve(t *testing.T) {
	_, transError := deployed.EtherTokenContract.Transfer(test.GetTxOpts(
		context.AuthOwner, nil, big.NewInt(test.ONE_GWEI*2), 100000),
		deployed.InvestingAddress, oneHundredEth())
	test.IfNotNil(t, transError, "Error transferring token")
	context.Blockchain.Commit()

	ownerEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthOwner.From)
	resEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.InvestingAddress)

	// has one wei left
	if ownerEthBal.Cmp(big.NewInt(test.ONE_WEI)) != 0 {
		t.Errorf("Expected ether token balance of 0, got: %v", ownerEthBal)
	}

	// got the 100
	if resEthBal.Cmp(oneHundredEth()) != 0 {
		t.Errorf("Expected reserve of 100 Eth, got: %v", resEthBal)
	}

	t.Logf("Current Reserve balance: %v", test.Commafy(resEthBal))
}

func TestOneInvestor(t *testing.T) {
	// user1 deposit eth in the ethToken
	etBal1, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser1.From)
	if etBal1.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected ether token balance of 0, got: %v", etBal1)
	}
	_, depErr := deployed.EtherTokenContract.Deposit(test.GetTxOpts(context.AuthUser1,
		tenEth(), big.NewInt(test.ONE_GWEI*2), 100000))
	test.IfNotNil(t, depErr, "Error depositing ETH")
	context.Blockchain.Commit()
	// snapshot current balances
	etBal1Now, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser1.From)
	if etBal1Now.Cmp(tenEth()) != 0 {
		t.Errorf("Expected ether token balance of 10 eth, got: %v", etBal1Now)
	}
	mtBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser1.From)
	if mtBal.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected market token balance of 0, got: %v", mtBal)
	}
	t.Logf("User1 current market token balance: %v", mtBal)
	// invest price currently
	invPrice, _ := deployed.InvestingContract.GetInvestmentPrice(nil)
	t.Logf("Current invest price: %v", test.Commafy(invPrice))

	// user1 wants to invest 10 ETH. Must approve inv contract first...
	_, approveErr := deployed.EtherTokenContract.Approve(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), deployed.InvestingAddress, tenEth()) // up to 10 ETH
	test.IfNotNil(t, approveErr, fmt.Sprintf("Error approving market contract to spend: %v", approveErr))
	context.Blockchain.Commit()

	// investing has that allowance now
	allowed, _ := deployed.EtherTokenContract.Allowance(nil, context.AuthUser1.From, deployed.InvestingAddress)
	if allowed.Cmp(tenEth()) != 0 {
		t.Errorf("Expected allowance of 10 ETH, got: %v", allowed)
	}

	// the actual investment (now that we know we can)
	_, invErr := deployed.InvestingContract.Invest(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), tenEth())
	test.IfNotNil(t, invErr, "Error investing")
	context.Blockchain.Commit()
	// check current market token balance
	mtBalNow, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser1.From)
	if mtBalNow.Cmp(mtBal) != 1 {
		t.Errorf("Expected %v to be > %v", mtBalNow, mtBal)
	}
	t.Logf("User1 market token balance post 10 ETH investment: %v", test.Commafy(mtBalNow))
	// market token total supply should be updated
	mtSup, _ := deployed.MarketTokenContract.TotalSupply(nil)
	t.Logf("Market token total supply post user1 investment of 10 ETH: %v", test.Commafy(mtSup))
	// reserve should be updated
	resEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.InvestingAddress)
	t.Logf("Reserve balance post user1 investment of 10 ETH: %v", test.Commafy(resEthBal))
	// invest price should change
	invPriceNow, _ := deployed.InvestingContract.GetInvestmentPrice(nil)
	t.Logf("Investment Price post user1 investment of 10 ETH: %v", test.Commafy(invPriceNow))
}