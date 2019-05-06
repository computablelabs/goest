package scenariothree

import (
	"github.com/computablelabs/goest/tests/test"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"math/big"
	"os"
	"testing"
)

// ExtendedCtx, holds Auth objects capable of signing transactions.
// We need more users than other tests so added in extra
// Also holds the Geth simulated backend.
type ExtendedCtx struct {
	AuthOwner   *bind.TransactOpts
	AuthBackend *bind.TransactOpts
	AuthUser1   *bind.TransactOpts
	AuthUser2   *bind.TransactOpts
	AuthUser3   *bind.TransactOpts
	AuthUser4   *bind.TransactOpts
	AuthUser5   *bind.TransactOpts
	AuthUser6   *bind.TransactOpts
	Blockchain  *backends.SimulatedBackend
}

// GetExtendedContext returns a hydrated ExtendedCtx struct, ready for use.
// Given a bal argument, it assigns this as the wallet balance for
// each authorization object in the Ctx
func GetExtendedContext(bal *big.Int) *ExtendedCtx {
	authOwn := test.GetAuthObject()
	authBac := test.GetAuthObject()
	authU1 := test.GetAuthObject()
	authU2 := test.GetAuthObject()
	authU3 := test.GetAuthObject()
	authU4 := test.GetAuthObject()
	authU5 := test.GetAuthObject()
	authU6 := test.GetAuthObject()
	alloc := make(core.GenesisAlloc)
	alloc[authOwn.From] = core.GenesisAccount{Balance: bal}
	alloc[authBac.From] = core.GenesisAccount{Balance: bal}
	alloc[authU1.From] = core.GenesisAccount{Balance: bal}
	alloc[authU2.From] = core.GenesisAccount{Balance: bal}
	alloc[authU3.From] = core.GenesisAccount{Balance: bal}
	alloc[authU4.From] = core.GenesisAccount{Balance: bal}
	alloc[authU5.From] = core.GenesisAccount{Balance: bal}
	alloc[authU6.From] = core.GenesisAccount{Balance: bal}
	// 2nd arg is a gas limit, a uint64. we'll use 4.7 million
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &ExtendedCtx{
		AuthOwner:   authOwn,
		AuthBackend: authBac,
		AuthUser1:   authU1,
		AuthUser2:   authU2,
		AuthUser3:   authU3,
		AuthUser4:   authU4,
		AuthUser5:   authU5,
		AuthUser6:   authU6,
		Blockchain:  bc,
	}
}

// variables decalred here have package scope
var extContext *ExtendedCtx
var deployed *test.Dep
var deployedError error

func TestMain(m *testing.M) {
	// need this to create bigger ETH balances (literal will overflow)
	var x big.Int
	oneHundredEth := x.Mul(big.NewInt(test.ONE_WEI), big.NewInt(100))
	oneHundredOneEth := x.Add(oneHundredEth, big.NewInt(test.ONE_WEI))

	extContext = GetExtendedContext(oneHundredOneEth) // users have 101 ETH account bal
	deployed, deployedError = test.Deploy(oneHundredOneEth, big.NewInt(test.ONE_WEI),
		extContext.AuthOwner, extContext.Blockchain, &test.Params{
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
