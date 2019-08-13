package scenarioone

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
)

func oneHundredOneEth() *big.Int {
	one := big.NewInt(test.ONE_ETH)
	oneHundred := one.Mul(one, big.NewInt(100))
	return oneHundred.Add(oneHundred, big.NewInt(test.ONE_ETH))
}

func oneHundredEth() *big.Int {
	one := big.NewInt(test.ONE_ETH)
	return one.Mul(one, big.NewInt(100))
}

func tenEth() *big.Int {
	one := big.NewInt(test.ONE_ETH)
	return one.Mul(one, big.NewInt(10))
}

// as we start with 0 supply and balances, we'll need authOwner to deposit some
func TestDeposit(t *testing.T) {
	_, depErr := deployed.EtherTokenContract.Deposit(test.GetTxOpts(context.AuthOwner, oneHundredOneEth(),
		big.NewInt(test.ONE_GWEI*2), 100000))
	test.IfNotNil(t, depErr, fmt.Sprintf("Error depositing funds from owner to ether token: %v", depErr))

	context.Blockchain.Commit()

	ownerBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthOwner.From)
	if ownerBal.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected %v to be > 0", ownerBal)
	}
}

func TestInitialBalance(t *testing.T) {
	// owner has all the eth atm
	etBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthOwner.From)
	// all market tokens that exist atm
	mtSup, _ := deployed.MarketTokenContract.TotalSupply(nil)

	if etBal.Cmp(oneHundredOneEth()) != 0 {
		t.Errorf("Expected ether token balance of %v, got: %v", oneHundredOneEth(), etBal)
	}
	if mtSup.Cmp(big.NewInt(test.ONE_ETH)) != 0 {
		t.Errorf("Expected market token supply of 1wei, got: %v", mtSup)
	}

	t.Logf("Current Market Token total supply: %v", test.Commafy(mtSup))
}

func TestTransferToReserve(t *testing.T) {
	_, transError := deployed.EtherTokenContract.Transfer(test.GetTxOpts(
		context.AuthOwner, nil, big.NewInt(test.ONE_GWEI*2), 100000),
		deployed.ReserveAddress, oneHundredEth())
	test.IfNotNil(t, transError, "Error transferring token")
	context.Blockchain.Commit()

	ownerEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthOwner.From)
	resEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.ReserveAddress)

	// has one wei left
	if ownerEthBal.Cmp(big.NewInt(test.ONE_ETH)) != 0 {
		t.Errorf("Expected ether token balance of 0, got: %v", ownerEthBal)
	}

	// got the 100
	if resEthBal.Cmp(oneHundredEth()) != 0 {
		t.Errorf("Expected reserve of 100 Eth, got: %v", resEthBal)
	}

	t.Logf("Current Reserve balance: %v", test.Commafy(resEthBal))
}

func TestOneSupporter(t *testing.T) {
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
	// support price currently
	invPrice, _ := deployed.ReserveContract.GetSupportPrice(nil)
	t.Logf("Current support price: %v", test.Commafy(invPrice))

	// user1 wants to place 10 ETH. Must approve inv contract first...
	_, approveErr := deployed.EtherTokenContract.Approve(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), deployed.ReserveAddress, tenEth()) // up to 10 ETH
	test.IfNotNil(t, approveErr, fmt.Sprintf("Error approving market contract to spend: %v", approveErr))
	context.Blockchain.Commit()

	// reserve has that allowance now
	allowed, _ := deployed.EtherTokenContract.Allowance(nil, context.AuthUser1.From, deployed.ReserveAddress)
	if allowed.Cmp(tenEth()) != 0 {
		t.Errorf("Expected allowance of 10 ETH, got: %v", allowed)
	}

	// the actual support (now that we know we can)
	_, invErr := deployed.ReserveContract.Support(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), tenEth())
	test.IfNotNil(t, invErr, "Error supporting")
	context.Blockchain.Commit()
	// check current market token balance
	mtBalNow, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser1.From)
	if mtBalNow.Cmp(mtBal) != 1 {
		t.Errorf("Expected %v to be > %v", mtBalNow, mtBal)
	}
	t.Logf("User1 market token balance post 10 ETH support: %v", test.Commafy(mtBalNow))
	// market token total supply should be updated
	mtSup, _ := deployed.MarketTokenContract.TotalSupply(nil)
	t.Logf("Market token total supply post user1 support of 10 ETH: %v", test.Commafy(mtSup))
	// reserve should be updated
	resEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.ReserveAddress)
	t.Logf("Reserve balance post user1 support of 10 ETH: %v", test.Commafy(resEthBal))
	// support price should change
	invPriceNow, _ := deployed.ReserveContract.GetSupportPrice(nil)
	t.Logf("Support Price post user1 support of 10 ETH: %v", test.Commafy(invPriceNow))
}
