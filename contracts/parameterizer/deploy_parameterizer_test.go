package parameterizer

import (
	"math/big"
	"os"
	"testing"
)

// variables decalred here have package scope
var context *ctx
var deployed *dep
var deployedError error

func TestDeployParameterizer(t *testing.T) {
	t.Log("Parameterizer contract should deploy correctly")

	if deployedError != nil {
		t.Fatalf("Failed to deploy the Parameterizer contract or a dependency: %v", deployedError)
	}

	if len(deployed.TokenAddress.Bytes()) == 0 {
		t.Error("Expected a valid token deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.VotingAddress.Bytes()) == 0 {
		t.Error("Expected a valid voting deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.ParameterizerAddress.Bytes()) == 0 {
		t.Error("Expected a valid parameterizer deployment address to be returned from deploy, got empty byte array instead")
	}
}

// we'll locate our testmain in these deploy_foo_test files as a pattern.
// NOTE the test main is run once per package, therefore
// the ctx and dep vars will be avail to the other tests in the package
func TestMain(m *testing.M) {
	// see ./helpers#context
	context = SetupBlockchain(big.NewInt(1000000000000000000)) // 1 ETH in wei
	// see ./helpers#deployed
	deployed, deployedError = Deploy(big.NewInt(1000), context)
	code := m.Run()
	os.Exit(code)
}
