package scenariofive

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
	"time"
)

// variables decalred here have package scope
var context *test.Ctx

// you need a way to refer to the newly created investors
var investors map[string]*bind.TransactOpts
var makers map[string]*bind.TransactOpts
var buyers map[string]*bind.TransactOpts
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

func setupMakers(bal *big.Int, numMakers int) {
	makers = make(map[string]*bind.TransactOpts)
	for i := 0; i < numMakers; i++ {
		var ary []string
		ary = append(ary, "maker", fmt.Sprintf("%v", (i+1)))
		key := strings.Join(ary, "")
		authMaker := test.GetAuthObject()
		// add them to the original allocation
		context.Alloc[authMaker.From] = core.GenesisAccount{Balance: bal}
		// add them to the map
		makers[key] = authMaker
	}
}

func setupBuyers(bal *big.Int, numBuyers int) {
	buyers = make(map[string]*bind.TransactOpts)
	for i := 0; i < numBuyers; i++ {
		var ary []string
		ary = append(ary, "buyer", fmt.Sprintf("%v", (i+1)))
		key := strings.Join(ary, "")
		authMaker := test.GetAuthObject()
		// add them to the original allocation
		context.Alloc[authMaker.From] = core.GenesisAccount{Balance: bal}
		// add them to the map
		buyers[key] = authMaker
	}
}

func TestMain(m *testing.M) {
	var x big.Int
	oneHundredEth := x.Mul(big.NewInt(test.ONE_WEI), big.NewInt(100))
	oneHundredOneEth := x.Add(oneHundredEth, big.NewInt(test.ONE_WEI))

	//context = test.GetContext(big.NewInt(test.ONE_WEI * 3)) // users have 3 ETH
	context = test.GetContext(oneHundredOneEth) // users have 101 ETH account bal
	setupInvestors(oneHundredOneEth, 10)        // investors have 101 ETH account balance
	setupMakers(big.NewInt(test.ONE_WEI), 10)   // makers have 1 ETH account balance
	setupBuyers(oneHundredOneEth, 1)            // buyers have 101 ETH account balance

	// override the original simulated backend now that we have appeneded to the allocation
	context.Blockchain = backends.NewSimulatedBackend(context.Alloc, 4700000)

	deployed, deployedError = test.Deploy(oneHundredOneEth, big.NewInt(test.ONE_WEI), context, &test.Params{
		ConversionRate: big.NewInt(test.ONE_SZABO),
		Spread:         big.NewInt(110),
		ListReward:     big.NewInt(250000000000000),   // 2.5 x 10**13
		Stake:          big.NewInt(10000000000000000), // 1 X 10**16
		VoteBy:         big.NewInt(100),
		Quorum:         big.NewInt(50),
		BackendPct:     big.NewInt(25),
		MakerPct:       big.NewInt(50),
		CostPerByte:    big.NewInt(test.ONE_GWEI * 100),
	})

	// setup the datatrust with a backend
	_, regErr := deployed.DatatrustContract.Register(test.GetTxOpts(context.AuthBackend, nil,
		big.NewInt(test.ONE_GWEI*2), 500000), "https://www.immabackend.biz")
	test.IfNotNil(&logr{}, regErr, fmt.Sprintf("Error registering for backend status: %v", regErr))

	context.Blockchain.Commit()

	// vote for the backend candidate, member will likely need funds
	transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner,
		context.AuthUser3.From, big.NewInt(10000000000000000)) // 1 X 10**16
	test.IfNotNil(&logr{}, transErr, "Error transferring tokens")

	// member will need to have approved the voting contract to spend
	appErr := test.MaybeIncreaseMarketTokenApproval(context, deployed, context.AuthUser3,
		deployed.VotingAddress, big.NewInt(10000000000000000)) // 1 X 10**16
	test.IfNotNil(&logr{}, appErr, "Error increasing allowance")

	hash, _ := deployed.DatatrustContract.GetHash(nil, "https://www.immabackend.biz")
	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), hash, big.NewInt(1))
	test.IfNotNil(&logr{}, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))

	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	// call for resolution
	_, resolveErr := deployed.DatatrustContract.ResolveRegistration(test.GetTxOpts(context.AuthUser2, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), hash)
	test.IfNotNil(&logr{}, resolveErr, fmt.Sprintf("Error resolving application: %v", resolveErr))

	context.Blockchain.Commit()

	// member can unstake now
	_, unErr := deployed.VotingContract.Unstake(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), hash)
	test.IfNotNil(&logr{}, unErr, fmt.Sprintf("Error Unstaking: %v", unErr))

	code := m.Run()
	os.Exit(code)
}
