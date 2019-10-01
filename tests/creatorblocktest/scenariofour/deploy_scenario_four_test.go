package scenariofour

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

// you need a way to refer to the newly created makers
var makers map[string]*bind.TransactOpts
var deployed *test.Dep
var deployedError error

// assemble an object which implements the helper's loggish interface
type logr struct{}

func (l *logr) Fatal(a ...interface{}) {
	log.Fatal(a...)
}

func setupMakers(bal *big.Int, numMakers int) {
	makers = make(map[string]*bind.TransactOpts)
	for i := 0; i < numMakers; i++ {
		var ary []string
		ary = append(ary, "makers", fmt.Sprintf("%v", (i+1)))
		key := strings.Join(ary, "")
		authMaker := test.GetAuthObject()
		// add them to the original allocation
		context.Alloc[authMaker.From] = core.GenesisAccount{Balance: bal}
		// add them to the map
		makers[key] = authMaker
	}
}

func TestMain(m *testing.M) {
	// need this to create bigger ETH balances (literal will overflow)
	var x big.Int
	oneHundredEth := x.Mul(big.NewInt(test.ONE_ETH), big.NewInt(100))
	oneHundredTwoEth := x.Add(oneHundredEth, big.NewInt(test.ONE_ETH*2))

	context = test.GetContext(oneHundredTwoEth) // users have 102 ETH account bal
	setupMakers(big.NewInt(test.ONE_ETH), 10)   // Makers have 1 ETH account balance

	// override the original simulated backend now that we have appeneded to the allocation
	context.Blockchain = backends.NewSimulatedBackend(context.Alloc, 4700000)

	deployed, deployedError = test.Deploy(big.NewInt(test.ONE_ETH), context, &test.Params{
		PriceFloor:  big.NewInt(test.ONE_MWEI),
		Spread:      big.NewInt(110),
		ListReward:  big.NewInt(250000000000000),      // 2.5 x 10**13
		Stake:       big.NewInt(test.ONE_FINNEY * 10), // 1 X 10**16
		VoteBy:      big.NewInt(test.MIN_VOTE_BY),
		Plurality:   big.NewInt(50),
		BackendPct:  big.NewInt(25),
		MakerPct:    big.NewInt(25),
		CostPerByte: big.NewInt(test.ONE_GWEI * 100),
	})

	// setup the datatrust with a backend, the auth will need CMT to stake
	transErr1 := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthBackend.From,
		big.NewInt(test.ONE_FINNEY*10))
	test.IfNotNil(&logr{}, transErr1, "Error maybe transferring market tokens")

	// member will need to have approved the voting contract to spend at least the stake
	incErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthBackend, deployed.VotingAddress,
		big.NewInt(test.ONE_FINNEY*10))
	test.IfNotNil(&logr{}, incErr, "Error maybe transferring market token approval")

	_, regErr := deployed.DatatrustContract.Register(test.GetTxOpts(context.AuthBackend, nil,
		big.NewInt(test.ONE_GWEI*2), 500000), "https://www.immabackend.biz")
	test.IfNotNil(&logr{}, regErr, fmt.Sprintf("Error registering for backend status: %v", regErr))

	context.Blockchain.Commit()

	// vote for the backend candidate, member will likely need funds
	transErr2 := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner,
		context.AuthUser3.From, big.NewInt(10000000000000000)) // 1 X 10**16
	test.IfNotNil(&logr{}, transErr2, "Error transferring tokens")

	// member will need to have approved the voting contract to spend
	appErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser3,
		deployed.VotingAddress, big.NewInt(10000000000000000)) // 1 X 10**16
	test.IfNotNil(&logr{}, appErr, "Error increasing allowance")

	hash, _ := deployed.DatatrustContract.GetHash(nil, "https://www.immabackend.biz")
	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), hash, big.NewInt(1))
	test.IfNotNil(&logr{}, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))

	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(test.MIN_VOTE_BY * time.Second)
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
