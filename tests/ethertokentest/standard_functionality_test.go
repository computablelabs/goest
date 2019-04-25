package ethertokentest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
)

func TestTransferFrom(t *testing.T) {
	ownerBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthOwner.From)

	// if we wanted to check the owner's account bal
	// t.Logf("owner balance: %v", context.Alloc[context.AuthOwner.From].Balance)

	// transfer from owner to user
	_, err := deployed.EtherTokenContract.Transfer(test.GetTxOpts(context.AuthOwner, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), context.AuthUser1.From, big.NewInt(test.ONE_WEI*4))
	test.IfNotNil(t, err, fmt.Sprintf("Error transferring funds from owner to user: %v", err))

	// gas cost looking to be around ~50k consistently
	// t.Logf("owner to user transfer gas cost: %v", tx.Cost())

	context.Blockchain.Commit()

	// transfer from user to other user
	_, err2 := deployed.EtherTokenContract.Transfer(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), context.AuthUser2.From, big.NewInt(test.ONE_WEI*2))
	test.IfNotNil(t, err2, fmt.Sprintf("Error transferring funds from user to other user: %v", err2))

	context.Blockchain.Commit()

	// owner should have 4 subtracted
	expectedBal := ownerBal.Sub(ownerBal, big.NewInt(test.ONE_WEI*4))
	newOwnerBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthOwner.From)
	if newOwnerBal.Cmp(expectedBal) != 0 {
		t.Errorf("Expected owner balance of %v, got %v", expectedBal, newOwnerBal)
	}

	// user should have had 2 subtracted from his 4
	userBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser1.From)
	if userBal.Cmp(big.NewInt(test.ONE_WEI*2)) != 0 {
		t.Errorf("Expected user balance of 2 tokens in wei, got %v", userBal)
	}

	otherBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser2.From)
	expectedOtherBal := otherBal.Add(otherBal, big.NewInt(test.ONE_WEI*2))
	if otherBal.Cmp(expectedOtherBal) != 0 {
		t.Errorf("Expected other user balance of %v, got %v", expectedOtherBal, otherBal)
	}
}

func TestApprove(t *testing.T) {
	_, err := deployed.EtherTokenContract.Approve(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), context.AuthUser2.From, big.NewInt(test.ONE_WEI*2)) // 2 tokens
	test.IfNotNil(t, err, fmt.Sprintf("Error approving spender: %v", err))

	context.Blockchain.Commit()
}

func TestAllowance(t *testing.T) {
	allowed, _ := deployed.EtherTokenContract.Allowance(nil, context.AuthUser1.From, context.AuthUser2.From)
	if allowed.Cmp(big.NewInt(test.ONE_WEI*2)) != 0 {
		t.Errorf("Expected spender to be approved for 2 tokens, got %v", allowed)
	}
}

func TestDecreaseApproval(t *testing.T) {
	_, err := deployed.EtherTokenContract.DecreaseApproval(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), context.AuthUser2.From, big.NewInt(test.ONE_WEI))
	test.IfNotNil(t, err, fmt.Sprintf("Error decreasing approval for spender: %v", err))

	context.Blockchain.Commit()

	allowed, _ := deployed.EtherTokenContract.Allowance(nil, context.AuthUser1.From, context.AuthUser2.From)
	if allowed.Cmp(big.NewInt(test.ONE_WEI)) != 0 {
		t.Errorf("Expected spender to have been decreased to 1, got %v", allowed)
	}
}

func TestIncreaseApproval(t *testing.T) {
	_, err := deployed.EtherTokenContract.IncreaseApproval(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), context.AuthUser2.From, big.NewInt(test.ONE_WEI))
	test.IfNotNil(t, err, fmt.Sprintf("Error increasing approval for spender: %v", err))

	context.Blockchain.Commit()

	allowed, _ := deployed.EtherTokenContract.Allowance(nil, context.AuthUser1.From, context.AuthUser2.From)
	if allowed.Cmp(big.NewInt(test.ONE_WEI*2)) != 0 {
		t.Errorf("Expected spender to have been increased to 2, got %v", allowed)
	}
}
