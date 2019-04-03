package datatrusttest

import (
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"os"
	"testing"
)

// variables decalred here have package scope
var context *test.Ctx
var deployed *test.Dep
var deployedError error

func TestDeployDatatrust(t *testing.T) {
	if deployedError != nil {
		t.Fatalf("Failed to deploy the Datatrust contract or a dependency: %v", deployedError)
	}

	if len(deployed.VotingAddress.Bytes()) == 0 {
		t.Fatal("Expected a valid voting deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.ParameterizerAddress.Bytes()) == 0 {
		t.Error("Expected a valid parameterizer deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.DatatrustAddress.Bytes()) == 0 {
		t.Error("Expected a valid Datatrust deployment address to be returned from deploy, got empty byte array instead")
	}
}

func TestDatatrustSetPrivileged(t *testing.T) {
	list, _ := deployed.DatatrustContract.GetPrivileged(nil)

	if list != deployed.ListingAddress {
		t.Fatalf("Expected Listing address of %v but got %v", deployed.ListingAddress, list)
	}
}

func TestMain(m *testing.M) {
	context = test.GetContext(big.NewInt(test.ONE_WEI * 3))                    // users have 3 ETH
	deployed, deployedError = test.Deploy(big.NewInt(test.ONE_WEI*6), context) // 6 tokens in wei
	code := m.Run()
	os.Exit(code)
}
