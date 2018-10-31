package market

import (
	"math/big"
	"os"
	"testing"
)

// variables decalred here have package scope
var context *ctx
var deployed *dep
var deployedError error

func TestDeployMarket(t *testing.T) {
	t.Log("Market contract should deploy correctly")

	if deployedError != nil {
		t.Fatalf("Failed to deploy the Market contract or a dependency: %v", deployedError)
	}

	if len(deployed.NetworkTokenAddress.Bytes()) == 0 {
		t.Error("Expected a valid network token deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.MarketTokenAddress.Bytes()) == 0 {
		t.Error("Expected a valid market token deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.VotingAddress.Bytes()) == 0 {
		t.Error("Expected a valid voting deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.ParameterizerAddress.Bytes()) == 0 {
		t.Error("Expected a valid parameterizer deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.MarketAddress.Bytes()) == 0 {
		t.Error("Expected a valid market deployment address to be returned from deploy, got empty byte array instead")
	}

	// ntr, _ := context.Blockchain.TransactionReceipt(nil, deployed.NetworkTokenTransaction.Hash())
	// mtr, _ := context.Blockchain.TransactionReceipt(nil, deployed.MarketTokenTransaction.Hash())
	// pr, _ := context.Blockchain.TransactionReceipt(nil, deployed.ParameterizerTransaction.Hash())
	// vr, _ := context.Blockchain.TransactionReceipt(nil, deployed.VotingTransaction.Hash())
	// mr, _ := context.Blockchain.TransactionReceipt(nil, deployed.MarketTransaction.Hash())

	// t.Log("******************** Costs to deploy all-the-things *************************")
	// t.Logf("Gas used to deploy Network Token %v", ntr.GasUsed)
	// t.Logf("Gas used to deploy Market Token %v", mtr.GasUsed)
	// t.Logf("Gas used to deploy Parameterizer %v", pr.GasUsed)
	// t.Logf("Gas used to deploy Voting %v", vr.GasUsed)
	// t.Logf("Gas used to deploy Market %v", mr.GasUsed)
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
