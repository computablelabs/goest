package burnable

import (
	"math/big"
	"os"
	"testing"
)

// variables decalred here have package scope
var ctx *context
var dep *deployed
var depErr error

func TestDeployConstructableBurnable(t *testing.T) {
	if depErr != nil {
		t.Fatalf("Failed to deploy the burnable token contract: %v", depErr)
	}

	if len(dep.address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address to be returned from deploy, got empty byte array instead")
	}

	// the supply should be the passed in initial balance, use Cmp to compare big ints
	if supply, _ := dep.contract.TotalSupply(nil); supply.Cmp(big.NewInt(1000000)) != 0 {
		t.Errorf("Expected total supply to equal initial balance, got %v", supply)
	}
}

// we'll locate our testmain in these deploy_foo_test files as a pattern.
// NOTE the test main is run once per package, therefore
// the ctx and dep vars will be avail to the other tests in the package
func TestMain(m *testing.M) {
	// see ./helpers#context
	ctx = SetupBlockchain(big.NewInt(1000000000))
	// see ./helpers#deployed
	dep, depErr = Deploy(big.NewInt(1000000), ctx)
	code := m.Run()
	os.Exit(code)
}
