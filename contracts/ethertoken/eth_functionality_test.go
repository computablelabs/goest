package ethertoken

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
)

func testDeposit(t *testing.T) {
	var x big.Int
	// the balance our member currently has
	memberBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthMember.From}, context.AuthMember.From)
	tokenBal, _ := deployed.Contract.TotalSupply(nil)

	// make a deposit in ETH, resulting in a 1:1 ethertoken balance
	_, depErr := deployed.Contract.Deposit(&bind.TransactOpts{
		From:     context.AuthMember.From,
		Signer:   context.AuthMember.Signer,
		Value:    big.NewInt(ONE_WEI), // depositing one eth in wei
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 100000,
	})

	if depErr != nil {
		t.Fatalf("Error depositing funds from member to ether token: %v", depErr)
	}

	// balances should be +1 eth now
	memberExpected := x.Add(memberBal, big.NewInt(ONE_WEI))
	newMemberBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthMember.From}, context.AuthMember.From)
	tokenExpected := x.Add(tokenBal, big.NewInt(ONE_WEI))
	newTokenBal, _ := deployed.Contract.TotalSupply(nil)

	if newMemberBal.Cmp(memberExpected) != 0 {
		t.Fatalf("Expected %v to be %v", newMemberBal, memberExpected)
	}

	if newTokenBal.Cmp(tokenExpected) != 0 {
		t.Fatalf("Expected %v to be %v", newTokenBal, tokenExpected)
	}
}

func testWithdrawal(t *testing.T) {
	var x big.Int

	memberBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthMember.From}, context.AuthMember.From)
	tokenBal, _ := deployed.Contract.TotalSupply(nil)

	// make a withdrawal in ETH
	_, withErr := deployed.Contract.Withdraw(&bind.TransactOpts{
		From:     context.AuthMember.From,
		Signer:   context.AuthMember.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 100000,
	}, big.NewInt(ONE_WEI/2)) // taking out .5 token

	if withErr != nil {
		t.Fatalf("Error withdrawing funds from ether token: %v", withErr)
	}

	// balances should be -.5 eth now
	memberExpected := x.Sub(memberBal, big.NewInt(ONE_WEI/2))
	newMemberBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthMember.From}, context.AuthMember.From)
	tokenExpected := x.Sub(tokenBal, big.NewInt(ONE_WEI/2))
	newTokenBal, _ := deployed.Contract.TotalSupply(nil)

	if newMemberBal.Cmp(memberExpected) != 0 {
		t.Fatalf("Expected %v to be %v", newMemberBal, memberExpected)
	}

	if newTokenBal.Cmp(tokenExpected) != 0 {
		t.Fatalf("Expected %v to be %v", newTokenBal, tokenExpected)
	}
}
