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
		big.NewInt(test.ONE_GWEI*2), 100000), context.AuthUser1.From, big.NewInt(test.ONE_GWEI*3))
	test.IfNotNil(t, err, fmt.Sprintf("Error transferring funds from owner to user: %v", err))

	// gas cost looking to be around ~50k consistently
	// t.Logf("owner to user transfer gas cost: %v", tx.Cost())

	context.Blockchain.Commit()

	// transfer from user to other user
	_, err2 := deployed.EtherTokenContract.Transfer(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), context.AuthUser2.From, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, err2, fmt.Sprintf("Error transferring funds from user to other user: %v", err2))

	context.Blockchain.Commit()

	// owner should have less...
	expectedBal := ownerBal.Sub(ownerBal, big.NewInt(test.ONE_GWEI*3))
	newOwnerBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthOwner.From)
	if newOwnerBal.Cmp(expectedBal) != 0 {
		t.Errorf("Expected owner balance of %v, got %v", expectedBal, newOwnerBal)
	}

	// user should have had 1 subtracted from his 3
	userBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser1.From)
	if userBal.Cmp(big.NewInt(test.ONE_GWEI*2)) != 0 {
		t.Errorf("Expected user balance of 1 GWEI, got %v", userBal)
	}

	otherBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser2.From)
	// also has one GWEI
	if otherBal.Cmp(big.NewInt(test.ONE_GWEI)) != 0 {
		t.Errorf("Expected other user balance of one GWEI, got %v", otherBal)
	}
}

func TestApprove(t *testing.T) {
	_, err := deployed.EtherTokenContract.Approve(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), context.AuthUser2.From, big.NewInt(test.ONE_ETH*2)) // 2 tokens
	test.IfNotNil(t, err, fmt.Sprintf("Error approving spender: %v", err))

	context.Blockchain.Commit()
}

func TestAllowance(t *testing.T) {
	allowed, _ := deployed.EtherTokenContract.Allowance(nil, context.AuthUser1.From, context.AuthUser2.From)
	if allowed.Cmp(big.NewInt(test.ONE_ETH*2)) != 0 {
		t.Errorf("Expected spender to be approved for 2 tokens, got %v", allowed)
	}
}

func TestDecreaseAllowance(t *testing.T) {
	_, err := deployed.EtherTokenContract.DecreaseAllowance(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), context.AuthUser2.From, big.NewInt(test.ONE_ETH))
	test.IfNotNil(t, err, fmt.Sprintf("Error decreasing approval for spender: %v", err))

	context.Blockchain.Commit()

	allowed, _ := deployed.EtherTokenContract.Allowance(nil, context.AuthUser1.From, context.AuthUser2.From)
	if allowed.Cmp(big.NewInt(test.ONE_ETH)) != 0 {
		t.Errorf("Expected spender to have been decreased to 1, got %v", allowed)
	}
}

func TestIncreaseAllowance(t *testing.T) {
	_, err := deployed.EtherTokenContract.IncreaseAllowance(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), context.AuthUser2.From, big.NewInt(test.ONE_ETH))
	test.IfNotNil(t, err, fmt.Sprintf("Error increasing approval for spender: %v", err))

	context.Blockchain.Commit()

	allowed, _ := deployed.EtherTokenContract.Allowance(nil, context.AuthUser1.From, context.AuthUser2.From)
	if allowed.Cmp(big.NewInt(test.ONE_ETH*2)) != 0 {
		t.Errorf("Expected spender to have been increased to 2, got %v", allowed)
	}
}
