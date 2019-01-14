package markettoken

import (
	"math/big"
	"os"
	"testing"
)

// variables decalred here have package scope
var context *ctx
var deployed *dep
var deployedError error

func TestDeployMarketToken(t *testing.T) {
	if deployedError != nil {
		t.Fatalf("Failed to deploy the market token contract: %v", deployedError)
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
	context = SetupBlockchain(big.NewInt(ONE_WEI))
	// see ./helpers#deployed
	deployed, deployedError = Deploy(big.NewInt(ONE_WEI*5), context) // deployed with a 5 token balance
	code := m.Run()
	os.Exit(code)
}
