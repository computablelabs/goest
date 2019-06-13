package ethertokentest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
)

func testDeposit(t *testing.T) {
	var x big.Int
	// the balance our member currently has
	memberBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser1.From)
	tokenBal, _ := deployed.EtherTokenContract.TotalSupply(nil)

	// make a deposit in ETH, resulting in a 1:1 ethertoken balance
	_, depErr := deployed.EtherTokenContract.Deposit(test.GetTxOpts(context.AuthUser1, big.NewInt(test.ONE_ETH),
		big.NewInt(test.ONE_GWEI*2), 100000))
	test.IfNotNil(t, depErr, fmt.Sprintf("Error depositing funds from member to ether token: %v", depErr))

	// balances should be +1 eth now
	memberExpected := x.Add(memberBal, big.NewInt(test.ONE_ETH))
	newMemberBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser1.From)
	tokenExpected := x.Add(tokenBal, big.NewInt(test.ONE_ETH))
	newTokenBal, _ := deployed.EtherTokenContract.TotalSupply(nil)

	if newMemberBal.Cmp(memberExpected) != 0 {
		t.Errorf("Expected %v to be %v", newMemberBal, memberExpected)
	}

	if newTokenBal.Cmp(tokenExpected) != 0 {
		t.Errorf("Expected %v to be %v", newTokenBal, tokenExpected)
	}
}

func testWithdrawal(t *testing.T) {
	var x big.Int

	memberBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser1.From)
	tokenBal, _ := deployed.EtherTokenContract.TotalSupply(nil)

	// make a withdrawal in ETH
	_, withErr := deployed.EtherTokenContract.Withdraw(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), big.NewInt(test.ONE_ETH/2)) // taking out .5 token
	test.IfNotNil(t, withErr, fmt.Sprintf("Error withdrawing funds from ether token: %v", withErr))

	// balances should be -.5 eth now
	memberExpected := x.Sub(memberBal, big.NewInt(test.ONE_ETH/2))
	newMemberBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser1.From)
	tokenExpected := x.Sub(tokenBal, big.NewInt(test.ONE_ETH/2))
	newTokenBal, _ := deployed.EtherTokenContract.TotalSupply(nil)

	if newMemberBal.Cmp(memberExpected) != 0 {
		t.Errorf("Expected %v to be %v", newMemberBal, memberExpected)
	}

	if newTokenBal.Cmp(tokenExpected) != 0 {
		t.Errorf("Expected %v to be %v", newTokenBal, tokenExpected)
	}
}
