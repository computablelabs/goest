package plcrvoting

import (
	"math/big"
	"os"
	"testing"
)

// variables decalred here have package scope
var context *ctx
var deployed *dep
var deployedError error

func TestDeployPLCRVoting(t *testing.T) {
	t.Log("Voting contract should deploy correctly")

	if deployedError != nil {
		t.Fatalf("Failed to deploy the Voting contract or Market token: %v", deployedError)
	}

	if len(deployed.TokenAddress.Bytes()) == 0 {
		t.Error("Expected a valid token deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.VotingAddress.Bytes()) == 0 {
		t.Error("Expected a valid voting deployment address to be returned from deploy, got empty byte array instead")
	}
}

// we'll locate our testmain in these deploy_foo_test files as a pattern.
// NOTE the test main is run once per package, therefore
// the ctx and dep vars will be avail to the other tests in the package
func TestMain(m *testing.M) {
	// see ./helpers#context
	context = SetupBlockchain(big.NewInt(8000000))
	// see ./helpers#deployed
	deployed, deployedError = Deploy(big.NewInt(1000), context)
	code := m.Run()
	os.Exit(code)
}
