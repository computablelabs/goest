package ethertokentest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestTotalSupply(t *testing.T) {
	// the supply should have been passed in as initial balance in the deploy spec
	if supply, _ := deployed.EtherTokenContract.TotalSupply(nil); supply.Cmp(big.NewInt(test.ONE_WEI*9)) != 0 {
		t.Errorf("Expected total supply to equal initial balance, got %v", supply)
	}
}

func TestTransfer(t *testing.T) {
	// should be able to feed HexToAddress a short string and get back a full length address...
	user := common.HexToAddress("0xabc")

	// NOTE: if we want to view the transaction itself, it would be the first return arg
	_, err := deployed.EtherTokenContract.Transfer(test.GetTxOpts(context.AuthFactory, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), user, big.NewInt(test.ONE_WEI))
	test.IfNotNil(t, err, fmt.Sprintf("Error transferring funds to another account: %v", err))

	context.Blockchain.Commit()
}

func TestBalanceOf(t *testing.T) {
	user := common.HexToAddress("0xabc")

	// the user should have been transfered 100000 whatever
	userBal, _ := deployed.EtherTokenContract.BalanceOf(nil, user)

	if userBal.Cmp(big.NewInt(test.ONE_WEI)) != 0 {
		t.Errorf("Expected user balance of 1 in wei, got %v", userBal)
	}

	// that 100000 should have been subtracted from the original owner, Auth.From in this case
	ownerBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthFactory.From)

	if ownerBal.Cmp(big.NewInt(test.ONE_WEI*8)) != 0 {
		t.Errorf("Expected owner balance of 8 tokens in wei, got %v", ownerBal)
	}
}
