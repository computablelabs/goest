package networktoken

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
)

func TestTransferFrom(t *testing.T) {
	ownerBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthFactory.From}, context.AuthFactory.From)

	// if we wanted to check the owner's account bal
	// t.Logf("owner balance: %v", context.Alloc[context.AuthFactory.From].Balance)

	// transfer from owner to user
	_, err := deployed.Contract.Transfer(&bind.TransactOpts{
		From:     context.AuthFactory.From,
		Signer:   context.AuthFactory.Signer,
		Value:    nil,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 100000,
	}, context.AuthMember.From, big.NewInt(10)) // 10 tokens

	if err != nil {
		t.Fatalf("Error transferring funds from owner to user: %v", err)
	}

	// gas cost looking to be around ~50k consistently
	// t.Logf("owner to user transfer gas cost: %v", tx.Cost())

	context.Blockchain.Commit()

	// transfer from user to other user
	_, err2 := deployed.Contract.Transfer(&bind.TransactOpts{
		From:     context.AuthMember.From,
		Signer:   context.AuthMember.Signer,
		Value:    nil,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 100000,
	}, context.OtherUser, big.NewInt(5))

	if err2 != nil {
		t.Fatalf("Error transferring funds from user to other user: %v", err2)
	}

	context.Blockchain.Commit()

	// owner should have 10 subtracted
	expectedBal := ownerBal.Sub(ownerBal, big.NewInt(10))
	newOwnerBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthFactory.From}, context.AuthFactory.From)
	if newOwnerBal.Cmp(expectedBal) != 0 {
		t.Errorf("Expected owner balance of %v, got %v", expectedBal, newOwnerBal)
	}

	// user should have had 5 subtracted from his 10
	userBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthMember.From}, context.AuthMember.From)
	if userBal.Cmp(big.NewInt(5)) != 0 {
		t.Errorf("Expected user balance of 5, got %v", userBal)
	}

	otherBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.OtherUser}, context.OtherUser)
	expectedOtherBal := otherBal.Add(otherBal, big.NewInt(5))
	if otherBal.Cmp(expectedOtherBal) != 0 {
		t.Errorf("Expected other user balance of %v, got %v", expectedOtherBal, otherBal)
	}
}

func TestApprove(t *testing.T) {
	_, err := deployed.Contract.Approve(&bind.TransactOpts{
		From:     context.AuthMember.From,
		Signer:   context.AuthMember.Signer,
		Value:    nil,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 100000,
	}, context.OtherContract, big.NewInt(4))

	if err != nil {
		t.Fatalf("Error approving spender: %v", err)
	}

	context.Blockchain.Commit()
}

func TestAllowance(t *testing.T) {
	allowed, _ := deployed.Contract.Allowance(&bind.CallOpts{From: context.AuthMember.From}, context.AuthMember.From, context.OtherContract)
	if allowed.Cmp(big.NewInt(4)) != 0 {
		t.Errorf("Expected spender to be approved for 4, got %v", allowed)
	}
}

func TestDecreaseApproval(t *testing.T) {
	_, err := deployed.Contract.DecreaseApproval(&bind.TransactOpts{
		From:     context.AuthMember.From,
		Signer:   context.AuthMember.Signer,
		Value:    nil,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 100000,
	}, context.OtherContract, big.NewInt(1))

	if err != nil {
		t.Errorf("Error decreasing approval for spender: %v", err)
	}

	context.Blockchain.Commit()

	allowed, _ := deployed.Contract.Allowance(&bind.CallOpts{From: context.AuthMember.From}, context.AuthMember.From, context.OtherContract)
	if allowed.Cmp(big.NewInt(3)) != 0 {
		t.Errorf("Expected spender to have been decreased to 3, got %v", allowed)
	}
}

func TestIncreaseApproval(t *testing.T) {
	_, err := deployed.Contract.IncreaseApproval(&bind.TransactOpts{
		From:     context.AuthMember.From,
		Signer:   context.AuthMember.Signer,
		Value:    nil,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 100000,
	}, context.OtherContract, big.NewInt(2))

	if err != nil {
		t.Errorf("Error increasing approval for spender: %v", err)
	}

	context.Blockchain.Commit()

	allowed, _ := deployed.Contract.Allowance(&bind.CallOpts{From: context.AuthMember.From}, context.AuthMember.From, context.OtherContract)
	if allowed.Cmp(big.NewInt(5)) != 0 {
		t.Errorf("Expected spender to have been increased to 5, got %v", allowed)
	}
}
