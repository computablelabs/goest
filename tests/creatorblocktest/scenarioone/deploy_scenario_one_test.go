package scenarioone

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

func TestMain(m *testing.M) {
	context = test.GetContext(big.NewInt(test.ONE_WEI * 2))
	// etherToken bal, marketTokenBal, ctx, params
	deployed, deployedError = test.Deploy(big.NewInt(test.ONE_WEI), big.NewInt(test.ONE_WEI), context, &test.Params{
		ConversionRate: big.NewInt(test.ONE_SZABO),
		Spread:         big.NewInt(100),
		ListReward:     big.NewInt(test.ONE_WEI),
		Stake:          big.NewInt(test.ONE_GWEI),
		VoteBy:         big.NewInt(100),
		Quorum:         big.NewInt(50),
		BackendPct:     big.NewInt(50),
		MakerPct:       big.NewInt(25),
		CostPerByte:    big.NewInt(test.ONE_GWEI * 100000),
	})
	code := m.Run()
	os.Exit(code)
}
