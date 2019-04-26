// Package for testing the Computable Ether Token Contract
package ethertokentest

import (
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"os"
	"testing"
)

var context *test.Ctx
var deployed *test.Dep
var deployedError error

func TestDeployEtherToken(t *testing.T) {
	if deployedError != nil {
		t.Fatalf("Failed to deploy the ether token contract: %v", deployedError)
	}

	if len(deployed.EtherTokenAddress.Bytes()) == 0 {
		t.Fatal("Expected a valid deployment address to be returned from deploy, got empty byte array instead")
	}

}

func TestMain(m *testing.M) {
	context = test.GetContext(big.NewInt(test.ONE_WEI * 2)) // 2 ETH in wei
	deployed, deployedError = test.Deploy(big.NewInt(test.ONE_WEI*9), big.NewInt(test.ONE_WEI*9), context, &test.Params{
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
