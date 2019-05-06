package votingtest

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

func TestDeployVoting(t *testing.T) {
	if deployedError != nil {
		t.Fatalf("Failed to deploy the Voting contract or Market token: %v", deployedError)
	}

	if len(deployed.VotingAddress.Bytes()) == 0 {
		t.Fatal("Expected a valid voting deployment address to be returned from deploy, got empty byte array instead")
	}
}

func TestSetPrivileged(t *testing.T) {
	p11r, data, list, invest, _ := deployed.VotingContract.GetPrivileged(nil)

	if list != deployed.ListingAddress {
		t.Fatalf("Expected Listing address of %v but got %v", deployed.ListingAddress, list)
	}

	if invest != deployed.InvestingAddress {
		t.Fatalf("Expected Investing address of %v but got %v", deployed.InvestingAddress, invest)
	}

	if p11r != deployed.ParameterizerAddress {
		t.Fatalf("Expected p11r address of %v but got %v", deployed.ParameterizerAddress, p11r)
	}

	if data != deployed.DatatrustAddress {
		t.Fatalf("Expected datatrust address of %v but got %v", deployed.DatatrustAddress, data)
	}
}

func TestMain(m *testing.M) {
	context = test.GetContext(big.NewInt(test.ONE_WEI))
	deployed, deployedError = test.Deploy(big.NewInt(test.ONE_WEI*5), big.NewInt(test.ONE_WEI*5), context.AuthOwner, context.Blockchain, &test.Params{
		ConversionRate: big.NewInt(test.ONE_GWEI),
		Spread:         big.NewInt(110),
		ListReward:     big.NewInt(test.ONE_WEI),
		Stake:          big.NewInt(test.ONE_GWEI),
		VoteBy:         big.NewInt(100),
		Quorum:         big.NewInt(50),
		BackendPct:     big.NewInt(25),
		MakerPct:       big.NewInt(50),
		CostPerByte:    big.NewInt(test.ONE_FINNEY * 6),
	})
	code := m.Run()
	os.Exit(code)
}
