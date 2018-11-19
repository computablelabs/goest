package networktoken

import (
	"math/big"
	"os"
	"testing"
)

// variables decalred here have package scope
var context *ctx
var deployed *dep
var deployedError error

func TestDeployNetworkToken(t *testing.T) {
	t.Log("Network token contract should deploy correctly")

	if deployedError != nil {
		t.Fatalf("Failed to deploy the network token contract: %v", deployedError)
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
	context = SetupBlockchain(big.NewInt(1000000000000000000)) // 1ETH in wei
	// see ./helpers#deployed
	deployed, deployedError = Deploy(big.NewInt(1000), context) // 1000 token balance
	code := m.Run()
	os.Exit(code)
}
