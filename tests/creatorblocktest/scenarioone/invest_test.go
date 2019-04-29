package scenarioone

import (
	// 	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
)

func TestInitialBalance(t *testing.T) {
	etBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthOwner.From)
	mtBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthOwner.From)
	if etBal.Cmp(big.NewInt(test.ONE_WEI)) != 0 {
		t.Errorf("Expected ether token balance of 1, got: %v", etBal)
	}

	if mtBal.Cmp(big.NewInt(test.ONE_WEI)) != 0 {
		t.Errorf("Expected market token balance of 1, got: %v", etBal)
	}
}

func TestTransferToReserve(t *testing.T) {
	// _, transError := deployed.EtherTokenContract.Transfer(test.GetTxOpts(
	// context.AuthOwner, nil, big.NewInt(test.ONE_GWEI*2), 100000),
	// deployed.InvestingAddress, big.NewInt(test.ONE_WEI))
}
