package reservetest

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

func TestDeployReserve(t *testing.T) {
	if deployedError != nil {
		t.Fatalf("Failed to deploy the Reserve contract or a dependency: %v", deployedError)
	}
}

func TestMain(m *testing.M) {
	// see ./helpers#context
	context = test.GetContext(big.NewInt(test.ONE_ETH * 3)) // users have 3 ETH
	// see ./helpers#deployed
	deployed, deployedError = test.Deploy(big.NewInt(test.ONE_ETH*6), context, &test.Params{
		PriceFloor:  big.NewInt(test.ONE_GWEI),
		Spread:      big.NewInt(110),
		ListReward:  big.NewInt(test.ONE_ETH),
		Stake:       big.NewInt(test.ONE_GWEI),
		VoteBy:      big.NewInt(100),
		Plurality:   big.NewInt(50),
		BackendPct:  big.NewInt(25),
		MakerPct:    big.NewInt(50),
		CostPerByte: big.NewInt(test.ONE_KWEI * 6),
	})
	code := m.Run()
	os.Exit(code)
}
