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
	// need this to create bigger ETH balances (literal will overflow)
	var x big.Int
	elevenEth := x.Mul(big.NewInt(test.ONE_ETH), big.NewInt(11))
	oneHundredEth := x.Mul(big.NewInt(test.ONE_ETH), big.NewInt(100))
	oneHundredOneEth := x.Add(oneHundredEth, big.NewInt(test.ONE_ETH))

	context = test.GetContext(elevenEth) // users have 11 ETH account bal
	// etherToken bal, marketToken bal, ctx, params (args)
	deployed, deployedError = test.Deploy(oneHundredOneEth, big.NewInt(test.ONE_ETH),
		context, &test.Params{
			PriceFloor:  big.NewInt(test.ONE_MWEI),
			Spread:      big.NewInt(110),
			ListReward:  big.NewInt(250000000000000),   // 2.5 x 10**13
			Stake:       big.NewInt(10000000000000000), // 1 X 10**16
			VoteBy:      big.NewInt(100),               // no need to use a "real" voteBy
			Plurality:   big.NewInt(50),
			BackendPct:  big.NewInt(25),
			MakerPct:    big.NewInt(25),
			CostPerByte: big.NewInt(test.ONE_GWEI * 100),
		})

	code := m.Run()
	os.Exit(code)
}
