package parameterizertest

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

func TestDeployParameterizer(t *testing.T) {
	if deployedError != nil {
		t.Fatalf("Failed to deploy the Parameterizer contract or a dependency: %v", deployedError)
	}

	if len(deployed.ParameterizerAddress.Bytes()) == 0 {
		t.Fatal("Expected a valid parameterizer deployment address to be returned from deploy, got empty byte array instead")
	}
}

func TestMain(m *testing.M) {
	context = test.GetContext(big.NewInt(test.ONE_WEI * 5))
	deployed, deployedError = test.Deploy(big.NewInt(test.ONE_WEI*5), context)
	code := m.Run()
	os.Exit(code)
}
