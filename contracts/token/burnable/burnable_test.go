package burnable

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
)

func TestBurn(t *testing.T) {
	t.Log("Burnable token contract should burn tokens on demand")
	// owner should have 1000 tokens atm
	bal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthOwner.From}, context.AuthOwner.From)
	if bal.Cmp(big.NewInt(1000)) != 0 {
		t.Errorf("Expected owner balance of 1000, got %v", bal)
	}

	// burn 100 of them, check bal
	_, err := deployed.Contract.Burn(&bind.TransactOpts{
		From:   context.AuthOwner.From,
		Signer: context.AuthOwner.Signer,
		Value:  nil,
	}, big.NewInt(100))

	if err != nil {
		t.Fatalf("Error burning funds: %v", err)
	}

	context.Blockchain.Commit()

	newBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthOwner.From}, context.AuthOwner.From)
	if newBal.Cmp(big.NewInt(900)) != 0 {
		t.Errorf("Expected owner balance of 900, got %v", newBal)
	}
}

// BurnFrom is used to adjust an allowance and funds balance
func TestBurnFrom(t *testing.T) {
	t.Log("Burnable token contract can use burnFrom to adjust both funds and allowance")
	// first, transfer some tokens to the user
	_, err := deployed.Contract.Transfer(&bind.TransactOpts{
		From:   context.AuthOwner.From,
		Signer: context.AuthOwner.Signer,
		Value:  nil,
	}, context.AuthUser.From, big.NewInt(100)) // 10 tokens

	if err != nil {
		t.Fatalf("Error transferring funds from owner to user: %v", err)
	}

	context.Blockchain.Commit()

	// user should have 100 tokens now
	bal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthUser.From}, context.AuthUser.From)
	if bal.Cmp(big.NewInt(100)) != 0 {
		t.Errorf("Expected user balance of 100, got %v", bal)
	}

	// user approves other contract to create an allowance
	_, err2 := deployed.Contract.Approve(&bind.TransactOpts{
		From:   context.AuthUser.From,
		Signer: context.AuthUser.Signer,
		Value:  nil,
	}, context.AuthOther.From, big.NewInt(50))

	if err2 != nil {
		t.Fatalf("Error approving spender: %v", err)
	}

	context.Blockchain.Commit()

	// allowance for other contract is 50 atm
	allowed, _ := deployed.Contract.Allowance(&bind.CallOpts{From: context.AuthUser.From}, context.AuthUser.From, context.AuthOther.From)
	if allowed.Cmp(big.NewInt(50)) != 0 {
		t.Errorf("Expected spender to be approved for 50, got %v", allowed)
	}

	// burnFrom should originate from the holder of an allowance
	_, err3 := deployed.Contract.BurnFrom(&bind.TransactOpts{
		From:     context.AuthOther.From,
		Signer:   context.AuthOther.Signer,
		Value:    nil,
		GasLimit: 100000, // minus a limit, gas may exceed threshold
	}, context.AuthUser.From, big.NewInt(10))

	// t.Logf("burnFrom gas cost: %v", tx.Cost())

	if err3 != nil {
		t.Fatalf("Error burning funds from other contract: %v", err3)
	}

	context.Blockchain.Commit()

	// burn from has decreased the allowance...
	newAllowed, _ := deployed.Contract.Allowance(&bind.CallOpts{From: context.AuthUser.From}, context.AuthUser.From, context.AuthOther.From)
	if newAllowed.Cmp(big.NewInt(40)) != 0 {
		t.Errorf("Expected adjusted allowance of 40, got %v", newAllowed)
	}

	// as well as burned the actual tokens
	newBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthUser.From}, context.AuthUser.From)
	if newBal.Cmp(big.NewInt(90)) != 0 {
		t.Errorf("Expected new user balance of 90, got %v", newBal)
	}
}
