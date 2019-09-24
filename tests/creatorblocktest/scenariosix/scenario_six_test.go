package scenariosix

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

func zeroEth() *big.Int {
	return big.NewInt(0)
}

func oneEth() *big.Int {
	return big.NewInt(test.ONE_ETH)
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
	if mtSup.Cmp(oneEth()) != 0 { // need this file to change
		t.Errorf("Expected market token supply of 1 whole, got: %v", mtSup)
	}

	t.Logf("Current Market Token total supply: %v", test.Commafy(mtSup))
}

// This differs from test in scenario-two since it has more supporters
func TestTransferToReserveThenSupport(t *testing.T) {

	resEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.ReserveAddress)

	t.Logf("Current reserve balance: %v", resEthBal)
	if resEthBal.Cmp(zeroEth()) != 0 {
		t.Errorf("Expected reserve of 0 Eth, got: %v", resEthBal)
	}

	t.Logf("Current Reserve balance: %v", test.Commafy(resEthBal))
	t.Logf("Number supporters: %d", len(supporters))

	for name, supporter := range supporters {

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
		t.Logf("%s current market token balance: %v", name, mtBal)

		sPrice, _ := deployed.ReserveContract.GetSupportPrice(nil)
		t.Logf("Current support price: %v", test.Commafy(sPrice))

		_, approveErr := deployed.EtherTokenContract.Approve(test.GetTxOpts(supporter, nil,
			big.NewInt(test.ONE_GWEI*2), 100000), deployed.ReserveAddress, oneHundredEth()) // up to 100 ETH
		test.IfNotNil(t, approveErr, fmt.Sprintf("Error approving market contract to spend: %v", approveErr))
		context.Blockchain.Commit()

		allowed, _ := deployed.EtherTokenContract.Allowance(nil, supporter.From, deployed.ReserveAddress)
		if allowed.Cmp(oneHundredEth()) != 0 {
			t.Errorf("Expected allowance of 100 ETH, got: %v", allowed)
		}

		// _, sErr := deployed.ReserveContract.Support(test.GetTxOpts(supporter, nil,
		// big.NewInt(test.ONE_GWEI*2), 150000), oneHundredEth())
		// test.IfNotNil(t, sErr, "Error supporting")
		// context.Blockchain.Commit()

		// check current market token balance
		// mtBalNow, _ := deployed.MarketTokenContract.BalanceOf(nil, supporter.From)
		// if mtBalNow.Cmp(mtBal) != 1 {
		// t.Errorf("Expected %v to be > %v", mtBalNow, mtBal)
		// }
		// t.Logf("%s market token balance post 100 ETH support: %v", name, test.Commafy(mtBalNow))

		// market token total supply should be updated
		// mtSup, _ := deployed.MarketTokenContract.TotalSupply(nil)
		// t.Logf("Market token total supply post %s support of 100 ETH: %v", name, test.Commafy(mtSup))

		// Get new reserve balance
		// resEthBal, _ = deployed.EtherTokenContract.BalanceOf(nil, deployed.ReserveAddress)
		// reserve should be updated
		// t.Logf("Reserve balance post %s support of 100 ETH: %v", name, test.Commafy(resEthBal))

		// sPriceNow, _ := deployed.ReserveContract.GetSupportPrice(nil)
		// t.Logf("Support Price post %s support of 100 ETH: %v", name, test.Commafy(sPriceNow))
	}
}
