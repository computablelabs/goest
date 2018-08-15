package token

import (
	"math/big"
	"testing"
)

// Expect to deploy correctly
func TestDeployConstructableBurnable(t *testing.T) {
	// use the test_helpers setup
	auth, blockchain := Setup(big.NewInt(1000000000))

	// do the actual deploy. args go [auth, bc, ...args to the contructor]
	address, _, contract, err := DeployConstructableBurnable(
		auth,
		blockchain,
		auth.From,
		big.NewInt(1000000),
	)

	// do stuff
	blockchain.Commit()

	if err != nil {
		t.Fatalf("Failed to deploy the burnable token contract: %v", err)
	}

	if len(address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address to be returned from deploy, got empty byte array instead")
	}

	// the supply should be the passed in initial balance, use Cmp to compare big ints
	if supply, _ := contract.TotalSupply(nil); supply.Cmp(big.NewInt(1000000)) != 0 {
		t.Errorf("Expected total supply to equal initial balance, got %v", supply)
	}
}
