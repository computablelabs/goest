package scenariothree

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core"
	"log"
	"math/big"
	"os"
	"testing"
	"time"
)

func getInvestors(bal *big.Int, numInvestors uint64) [](*bind.TransactOpts) {
	alloc := make(core.GenesisAlloc)
	investors := make([](*bind.TransactOpts), 0)
	var i uint64
	for i = 0; i < numInvestors; i++ {
		authInvestor := test.GetAuthObject()
		alloc[authInvestor.From] = core.GenesisAccount{Balance: bal}
		investors = append(investors, authInvestor)
	}
	return investors
}

// variables decalred here have package scope
var context *test.Ctx
var investors [](*bind.TransactOpts)
var deployed *test.Dep
var deployedError error

// assemble an object which implements the helper's loggish interface
type logr struct{}

func (l *logr) Fatal(a ...interface{}) {
	log.Fatal(a...)
}

func TestMain(m *testing.M) {
	// need this to create bigger ETH balances (literal will overflow)
	var x big.Int
	var numInvestors uint64 = 10
	oneHundredEth := x.Mul(big.NewInt(test.ONE_WEI), big.NewInt(100))
	oneHundredOneEth := x.Add(oneHundredEth, big.NewInt(test.ONE_WEI))

	context = test.GetContext(oneHundredOneEth)              // users have 101 ETH account bal
	investors = getInvestors(oneHundredOneEth, numInvestors) // investors have 101 ETH account balance
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

	// setup the datatrust with a backend
	_, regErr := deployed.DatatrustContract.Register(test.GetTxOpts(context.AuthBackend, nil,
		big.NewInt(test.ONE_GWEI*2), 500000), "https://www.immabackend.biz")
	test.IfNotNil(&logr{}, regErr, fmt.Sprintf("Error registering for backend status: %v", regErr))

	context.Blockchain.Commit()

	//// vote for the backend candidate, member will likely need funds
	//transErr := test.MaybeTransferMarketToken(context.Blockchain, deployed, context.AuthOwner,
	//	context.AuthUser3.From, big.NewInt(test.ONE_GWEI))
	//test.IfNotNil(&logr{}, transErr, "Error transferring tokens")

	// member will need to have approved the voting contract to spend
	appErr := test.MaybeIncreaseMarketTokenApproval(context, deployed, context.AuthUser3,
		deployed.VotingAddress, big.NewInt(test.ONE_GWEI))
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
