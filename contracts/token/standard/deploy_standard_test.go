package standard

import (
	"math/big"
	"os"
	"testing"
)

// variables decalred here have package scope
var context *ctx
var deployed *dep
var deployedError error

func TestDeployConstructableStandard(t *testing.T) {
	t.Log("Standard token contract should deploy correctly")

	if deployedError != nil {
		t.Fatalf("Failed to deploy the basic token contract: %v", deployedError)
	}

	if len(deployed.Address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address to be returned from deploy, got empty byte array instead")
	}

}

// we'll locate our testmain in these deploy_foo_test files as a pattern.
// NOTE the test main is run once per package, therefore
// the ctx and dep vars will be avail to the other tests in the package
func TestMain(m *testing.M) {
	// see ./helpers#context
	context = SetupBlockchain(big.NewInt(1000000)) // 1M
	// see ./helpers#deployed
	deployed, deployedError = Deploy(big.NewInt(100), context) // 100 token balance
	code := m.Run()
	os.Exit(code)
}
