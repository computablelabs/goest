package scenariothree

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"log"
	"math/big"
	"os"
	"strings"
	"testing"
)

// variables decalred here have package scope
var context *test.Ctx

// you need a way to refer to the newly created investors
var investors map[string]*bind.TransactOpts
var deployed *test.Dep
var deployedError error

// assemble an object which implements the helper's loggish interface
type logr struct{}

func (l *logr) Fatal(a ...interface{}) {
	log.Fatal(a...)
}

func setupInvestors(bal *big.Int, numInvestors int) {
	investors = make(map[string]*bind.TransactOpts)
	for i := 0; i < numInvestors; i++ {
		var ary []string
		ary = append(ary, "investor", fmt.Sprintf("%v", (i+1)))
		key := strings.Join(ary, "")
		authInvestor := test.GetAuthObject()
		// add them to the original allocation
		context.Alloc[authInvestor.From] = core.GenesisAccount{Balance: bal}
		// add them to the map
		investors[key] = authInvestor
	}
}

func TestMain(m *testing.M) {
	// need this to create bigger ETH balances (literal will overflow)
	var x big.Int
	oneHundredEth := x.Mul(big.NewInt(test.ONE_WEI), big.NewInt(100))
	oneHundredOneEth := x.Add(oneHundredEth, big.NewInt(test.ONE_WEI))

	context = test.GetContext(oneHundredOneEth) // users have 101 ETH account bal
	setupInvestors(oneHundredOneEth, 20)        // investors have 101 ETH account balance

	// override the original simulated backend now that we have appeneded to the allocation
	context.Blockchain = backends.NewSimulatedBackend(context.Alloc, 4700000)

	deployed, deployedError = test.Deploy(oneHundredOneEth, big.NewInt(test.ONE_WEI),
		context, &test.Params{
			ConversionRate: big.NewInt(test.ONE_SZABO),
			Spread:         big.NewInt(110),
			ListReward:     big.NewInt(250000000000000),   // 2.5 x 10**13
			Stake:          big.NewInt(10000000000000000), // 1 X 10**16
			VoteBy:         big.NewInt(100),               // no need to use a "real" voteBy
			Quorum:         big.NewInt(50),
			BackendPct:     big.NewInt(25),
			MakerPct:       big.NewInt(25),
			CostPerByte:    big.NewInt(test.ONE_GWEI * 100),
		})

	code := m.Run()
	os.Exit(code)
}
