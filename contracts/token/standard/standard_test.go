package standard

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
)

func TestTransferFrom(t *testing.T) {
	t.Log("Standard token should be able to transfer from one address to another")

	// if we wanted to check the owner's account bal
	// t.Logf("owner balance: %v", context.Alloc[context.AuthOwner.From].Balance)

	// transfer from owner to user
	_, err := deployed.Contract.Transfer(&bind.TransactOpts{
		From:   context.AuthOwner.From,
		Signer: context.AuthOwner.Signer,
		Value:  nil,
	}, context.AuthUser.From, big.NewInt(10)) // 10 tokens

	if err != nil {
		t.Fatalf("Error transferring funds from owner to user: %v", err)
	}

	// gas cost looking to be around ~50k consistently
	// t.Logf("owner to user transfer gas cost: %v", tx.Cost())

	context.Blockchain.Commit()

	// transfer from user to other user
	_, err2 := deployed.Contract.Transfer(&bind.TransactOpts{
		From:   context.AuthUser.From,
		Signer: context.AuthUser.Signer,
		Value:  nil,
	}, context.OtherUser, big.NewInt(5))

	if err2 != nil {
		t.Fatalf("Error transferring funds from user to other user: %v", err2)
	}

	context.Blockchain.Commit()

	// owner should have 10 subtracted from his 100
	ownerBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthOwner.From}, context.AuthOwner.From)
	if ownerBal.Cmp(big.NewInt(90)) != 0 {
		t.Errorf("Expected owner balance of 90, got %v", ownerBal)
	}

	// user should have had 5 subtracted from his 10
	userBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthUser.From}, context.AuthUser.From)
	if userBal.Cmp(big.NewInt(5)) != 0 {
		t.Errorf("Expected user balance of 5, got %v", userBal)
	}

	otherBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.OtherUser}, context.OtherUser)
	if otherBal.Cmp(big.NewInt(5)) != 0 {
		t.Errorf("Expected other user balance of 5, got %v", otherBal)
	}
}

func TestApprove(t *testing.T) {
	t.Log("Standard token should allow a user to approve a spender for some amount")

	_, err := deployed.Contract.Approve(&bind.TransactOpts{
		From:   context.AuthUser.From,
		Signer: context.AuthUser.Signer,
		Value:  nil,
	}, context.OtherContract, big.NewInt(4))

	if err != nil {
		t.Fatalf("Error approving spender: %v", err)
	}

	context.Blockchain.Commit()
}

func TestAllowance(t *testing.T) {
	t.Log("Standard token should be able to check the spending allowance of a given address for a given user")

	allowed, _ := deployed.Contract.Allowance(&bind.CallOpts{From: context.AuthUser.From}, context.AuthUser.From, context.OtherContract)
	if allowed.Cmp(big.NewInt(4)) != 0 {
		t.Errorf("Expected spender to be approved for 4, got %v", allowed)
	}
}

func TestDecreaseApproval(t *testing.T) {
	t.Log("Standard token should be able to decrease the spending allowance of a given address for a given user")

	_, err := deployed.Contract.DecreaseApproval(&bind.TransactOpts{
		From:   context.AuthUser.From,
		Signer: context.AuthUser.Signer,
		Value:  nil,
	}, context.OtherContract, big.NewInt(1))

	if err != nil {
		t.Errorf("Error decreasing approval for spender: %v", err)
	}

	context.Blockchain.Commit()

	allowed, _ := deployed.Contract.Allowance(&bind.CallOpts{From: context.AuthUser.From}, context.AuthUser.From, context.OtherContract)
	if allowed.Cmp(big.NewInt(3)) != 0 {
		t.Errorf("Expected spender to have been decreased to 3, got %v", allowed)
	}
}

func TestIncreaseApproval(t *testing.T) {
	t.Log("Standard token should be able to increase the spending allowance of a given address for a given user")

	_, err := deployed.Contract.IncreaseApproval(&bind.TransactOpts{
		From:   context.AuthUser.From,
		Signer: context.AuthUser.Signer,
		Value:  nil,
	}, context.OtherContract, big.NewInt(2))

	if err != nil {
		t.Errorf("Error increasing approval for spender: %v", err)
	}

	context.Blockchain.Commit()

	allowed, _ := deployed.Contract.Allowance(&bind.CallOpts{From: context.AuthUser.From}, context.AuthUser.From, context.OtherContract)
	if allowed.Cmp(big.NewInt(5)) != 0 {
		t.Errorf("Expected spender to have been increased to 5, got %v", allowed)
	}
}
