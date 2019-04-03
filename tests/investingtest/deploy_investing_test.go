package investingtest

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

func TestDeployInvesting(t *testing.T) {
	if deployedError != nil {
		t.Fatalf("Failed to deploy the Investing contract or a dependency: %v", deployedError)
	}
}

func TestMain(m *testing.M) {
	// see ./helpers#context
	context = test.GetContext(big.NewInt(test.ONE_WEI * 3)) // users have 3 ETH
	// see ./helpers#deployed
	deployed, deployedError = test.Deploy(big.NewInt(test.ONE_WEI*6), context) // 6 tokens in wei
	code := m.Run()
	os.Exit(code)
}
