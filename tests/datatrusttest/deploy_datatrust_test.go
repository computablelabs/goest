package datatrusttest

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

func TestDeployDatatrust(t *testing.T) {
	if deployedError != nil {
		t.Fatalf("Failed to deploy the Datatrust contract or a dependency: %v", deployedError)
	}

	if len(deployed.DatatrustAddress.Bytes()) == 0 {
		t.Error("Expected a valid Datatrust deployment address to be returned from deploy, got empty byte array instead")
	}
}

func TestDatatrustSetPrivileged(t *testing.T) {
	list, _ := deployed.DatatrustContract.GetPrivileged(nil)

	if list != deployed.ListingAddress {
		t.Fatalf("Expected Listing address of %v but got %v", deployed.ListingAddress, list)
	}
}

func TestDatatrustGetReserve(t *testing.T) {
	inv, _ := deployed.DatatrustContract.GetReserve(nil)

	if inv != deployed.ReserveAddress {
		t.Fatalf("Expected Reserve address of %v but got %v", deployed.ReserveAddress, inv)
	}
}

func TestMain(m *testing.M) {
	context = test.GetContext(big.NewInt(test.ONE_ETH * 3)) // users have 3 ETH
	deployed, deployedError = test.Deploy(big.NewInt(test.ONE_ETH*6), big.NewInt(test.ONE_ETH*6), context, &test.Params{
		ConversionRate: big.NewInt(test.ONE_GWEI),
		Spread:         big.NewInt(110),
		ListReward:     big.NewInt(test.ONE_ETH),
		Stake:          big.NewInt(test.ONE_GWEI),
		VoteBy:         big.NewInt(100),
		Quorum:         big.NewInt(50),
		BackendPct:     big.NewInt(25),
		MakerPct:       big.NewInt(50),
		CostPerByte:    big.NewInt(test.ONE_KWEI * 6),
	})
	code := m.Run()
	os.Exit(code)
}
