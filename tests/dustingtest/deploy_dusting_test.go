package dustingtest

import (
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"os"
	"testing"
)

var context *test.Ctx
var deployed *test.Dep
var deployedError error

func TestMain(m *testing.M) {
	context = test.GetContext(big.NewInt(test.ONE_ETH * 9)) // users have 3 ETH account bal
	// etherToken bal, marketToken bal, ctx, params (args)
	deployed, deployedError = test.Deploy(big.NewInt(test.ONE_ETH*9),
		context, &test.Params{
			PriceFloor:  big.NewInt(test.ONE_SZABO),
			Spread:      big.NewInt(110),
			ListReward:  big.NewInt(250000000000000),   // 2.5 x 10**13
			Stake:       big.NewInt(10000000000000000), // 1 X 10**16
			VoteBy:      big.NewInt(test.MIN_VOTE_BY),  // no need to use a "real" voteBy
			Plurality:   big.NewInt(50),
			BackendPct:  big.NewInt(39),
			MakerPct:    big.NewInt(21),
			CostPerByte: big.NewInt(test.ONE_GWEI * 100),
		})

	code := m.Run()
	os.Exit(code)
}
