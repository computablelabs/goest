package experiment

import (
	"math/big"
	"os"
	"testing"
)

// variables decalred here have package scope
var context *ctx
var deployed *dep
var deployedError error

func TestDeployExperiment(t *testing.T) {
	if deployedError != nil {
		t.Fatalf("Failed to deploy the experiment contract: %v", deployedError)
	}

	if len(deployed.Address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address to be returned from deploy, got empty byte array instead")
	}

}

func TestMain(m *testing.M) {
	// see ./helpers#context
	context = SetupBlockchain(big.NewInt(1e18 * 2)) // 2 ETH in wei
	// see ./helpers#deployed
	deployed, deployedError = Deploy(context) // 9 tokens in wei
	code := m.Run()
	os.Exit(code)
}
