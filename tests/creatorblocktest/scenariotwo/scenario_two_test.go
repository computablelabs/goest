package scenariotwo

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func TestTransferToReserveThenSupport(t *testing.T) {
	_, transError := deployed.EtherTokenContract.Transfer(test.GetTxOpts(
		context.AuthOwner, nil, big.NewInt(test.ONE_GWEI*2), 100000),
		deployed.ReserveAddress, oneHundredEth())
	test.IfNotNil(t, transError, "Error transferring token")
	context.Blockchain.Commit()

	ownerEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthOwner.From)
	resEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.ReserveAddress)

	// has 1 ETH left
	if ownerEthBal.Cmp(big.NewInt(test.ONE_ETH)) != 0 {
		t.Errorf("Expected ether token balance of 1 ETH, got: %v", ownerEthBal)
	}

	// got the 100
	if resEthBal.Cmp(oneHundredEth()) != 0 {
		t.Errorf("Expected reserve of 100 Eth, got: %v", resEthBal)
	}

	t.Logf("Current Reserve balance: %v", test.Commafy(resEthBal))

	var supporters [3](*bind.TransactOpts)
	supporters[0] = context.AuthUser1
	supporters[1] = context.AuthUser2
	supporters[2] = context.AuthUser3

	for ind, supporter := range supporters {

		// supporter deposits eth in the ethToken
		etBal, _ := deployed.EtherTokenContract.BalanceOf(nil, supporter.From)
		if etBal.Cmp(big.NewInt(0)) != 0 {
			t.Errorf("Expected ether token balance of 0, got: %v", etBal)
		}
		_, depErr := deployed.EtherTokenContract.Deposit(test.GetTxOpts(supporter,
			oneHundredEth(), big.NewInt(test.ONE_GWEI*2), 100000))
		test.IfNotNil(t, depErr, "Error depositing ETH")
		context.Blockchain.Commit()

		// snapshot current balances
		etBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, supporter.From)
		if etBalNow.Cmp(oneHundredEth()) != 0 {
			t.Errorf("Expected ether token balance of 100 eth, got: %v", etBalNow)
		}
		mtBal, _ := deployed.MarketTokenContract.BalanceOf(nil, supporter.From)
		if mtBal.Cmp(big.NewInt(0)) != 0 {
			t.Errorf("Expected market token balance of 0, got: %v", mtBal)
		}
		t.Logf("Supporter %d current market token balance: %v", (ind + 1), mtBal)

		// support price currently
		sPrice, _ := deployed.ReserveContract.GetSupportPrice(nil)
		t.Logf("Current support price: %v", test.Commafy(sPrice))

		// supporter wants to offer 100 ETH. Must approve res contract first...
		_, approveErr := deployed.EtherTokenContract.Approve(test.GetTxOpts(supporter, nil,
			big.NewInt(test.ONE_GWEI*2), 100000), deployed.ReserveAddress, oneHundredEth()) // up to 100 ETH
		test.IfNotNil(t, approveErr, fmt.Sprintf("Error approving market contract to spend: %v", approveErr))
		context.Blockchain.Commit()

		// reserve has that allowance now
		allowed, _ := deployed.EtherTokenContract.Allowance(nil, supporter.From, deployed.ReserveAddress)
		if allowed.Cmp(oneHundredEth()) != 0 {
			t.Errorf("Expected allowance of 100 ETH, got: %v", allowed)
		}

		_, sErr := deployed.ReserveContract.Support(test.GetTxOpts(supporter, nil,
			big.NewInt(test.ONE_GWEI*2), 150000), oneHundredEth())
		test.IfNotNil(t, sErr, "Error supporting")
		context.Blockchain.Commit()

		// check current market token balance
		mtBalNow, _ := deployed.MarketTokenContract.BalanceOf(nil, supporter.From)
		if mtBalNow.Cmp(mtBal) != 1 {
			t.Errorf("Expected %v to be > %v", mtBalNow, mtBal)
		}
		t.Logf("Supporter %d market token balance post 100 ETH support: %v", (ind + 1), test.Commafy(mtBalNow))

		// market token total supply should be updated
		mtSup, _ := deployed.MarketTokenContract.TotalSupply(nil)
		t.Logf("Market token total supply post supporter %d support of 100 ETH: %v", (ind + 1), test.Commafy(mtSup))

		// Get new reserve balance
		resEthBal, _ = deployed.EtherTokenContract.BalanceOf(nil, deployed.ReserveAddress)
		// reserve should be updated
		t.Logf("Reserve balance post supporter %d support of 100 ETH: %v", (ind + 1), test.Commafy(resEthBal))

		sPriceNow, _ := deployed.ReserveContract.GetSupportPrice(nil)
		t.Logf("Support Price post supporter %d support of 100 ETH: %v", (ind + 1), test.Commafy(sPriceNow))
	}
}
