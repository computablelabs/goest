// Package for testing the Computable Market Token Contract
package markettokentest

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

func TestDeployMarketToken(t *testing.T) {
	if deployedError != nil {
		t.Fatalf("Failed to deploy the market token contract: %v", deployedError)
	}

	if len(deployed.MarketTokenAddress.Bytes()) == 0 {
		t.Fatal("Expected a valid deployment address to be returned from deploy, got empty byte array instead")
	}
}

func TestMain(m *testing.M) {
	context = test.GetContext(big.NewInt(test.ONE_WEI))
	deployed, deployedError = test.Deploy(big.NewInt(test.ONE_WEI*5), context) // deployed with a 5 token balance
	code := m.Run()
	os.Exit(code)
}
