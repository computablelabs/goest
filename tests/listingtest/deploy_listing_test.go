package listingtest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"log"
	"math/big"
	"os"
	"testing"
	"time"
)

// variables decalred here have package scope
var context *test.Ctx
var deployed *test.Dep
var deployedError error

// assemble an object which implements the helper's loggish interface
type logr struct{}

func (l *logr) Fatal(a ...interface{}) {
	log.Fatal(a...)
}

func TestDeployListing(t *testing.T) {
	if deployedError != nil {
		t.Fatalf("Failed to deploy the Listing contract or a dependency: %v", deployedError)
	}

	// what does the hex string look like?
	// t.Log("********************************************************")
	// t.Log(deployed.MarketAddress.Hex())

	// ntr, _ := context.Blockchain.TransactionReceipt(nil, deployed.EtherTokenTransaction.Hash())
	// mtr, _ := context.Blockchain.TransactionReceipt(nil, deployed.MarketTokenTransaction.Hash())
	// pr, _ := context.Blockchain.TransactionReceipt(nil, deployed.ParameterizerTransaction.Hash())
	// vr, _ := context.Blockchain.TransactionReceipt(nil, deployed.VotingTransaction.Hash())
	// mr, _ := context.Blockchain.TransactionReceipt(nil, deployed.MarketTransaction.Hash())

	// t.Log("******************** Costs to deploy all-the-things *************************")
	// t.Logf("Gas used to deploy Ether Token %v", ntr.GasUsed)
	// t.Logf("Gas used to deploy Market Token %v", mtr.GasUsed)
	// t.Logf("Gas used to deploy Parameterizer %v", pr.GasUsed)
	// t.Logf("Gas used to deploy Voting %v", vr.GasUsed)
	// t.Logf("Gas used to deploy Market %v", mr.GasUsed)
}

func TestMain(m *testing.M) {
	context = test.GetContext(big.NewInt(test.ONE_ETH * 3)) // users have 3 ETH
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

	// setup the datatrust with a backend
	_, regErr := deployed.DatatrustContract.Register(test.GetTxOpts(context.AuthBackend, nil,
		big.NewInt(test.ONE_GWEI*2), 500000), "https://www.immabackend.biz")
	test.IfNotNil(&logr{}, regErr, fmt.Sprintf("Error registering for backend status: %v", regErr))

	context.Blockchain.Commit()

	// vote for the backend candidate, member will likely need funds
	transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner,
		context.AuthUser3.From, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(&logr{}, transErr, "Error transferring tokens")

	// member will need to have approved the voting contract to spend
	appErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser3,
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
