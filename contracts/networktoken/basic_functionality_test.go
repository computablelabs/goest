package networktoken

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestTotalSupply(t *testing.T) {
	// the supply should have been passed in as initial balance in the deploy spec
	if supply, _ := deployed.Contract.TotalSupply(nil); supply.Cmp(big.NewInt(1000)) != 0 {
		t.Errorf("Expected total supply to equal initial balance, got %v", supply)
	}
}

func TestTransfer(t *testing.T) {
	// should be able to feed HexToAddress a short string and get back a full length address...
	user := common.HexToAddress("0xabc")

	// NOTE: if we want to view the transaction itself, it would be the first return arg
	_, err := deployed.Contract.Transfer(&bind.TransactOpts{
		From:     context.AuthFactory.From,
		Signer:   context.AuthFactory.Signer,
		Value:    nil,
		GasPrice: big.NewInt(2000000000), //2 gwei
		GasLimit: 100000,
	}, user, big.NewInt(100))

	if err != nil {
		t.Fatalf("Error transferring funds to another account: %v", err)
	}

	context.Blockchain.Commit()
}

func TestBalanceOf(t *testing.T) {
	user := common.HexToAddress("0xabc")

	// the user should have been transfered 100000 whatever
	userBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: user}, user)

	if userBal.Cmp(big.NewInt(100)) != 0 {
		t.Errorf("Expected user balance of 100, got %v", userBal)
	}

	// that 100000 should have been subtracted from the original owner, Auth.From in this case
	ownerBal, _ := deployed.Contract.BalanceOf(&bind.CallOpts{From: context.AuthFactory.From}, context.AuthFactory.From)

	if ownerBal.Cmp(big.NewInt(900)) != 0 {
		t.Errorf("Expected owner balance of 900, got %v", ownerBal)
	}
}
