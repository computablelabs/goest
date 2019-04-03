package datatrusttest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
)

func TestPurchaseByteCredits(t *testing.T) {
	// user has no credits atm
	bytesBal, _ := deployed.DatatrustContract.GetByteCredits(nil, context.AuthUser3.From)

	// make a deposit in ETH, resulting in a 1:1 ethertoken balance
	_, depErr := deployed.EtherTokenContract.Deposit(test.GetTxOpts(context.AuthUser3,
		big.NewInt(test.ONE_WEI), big.NewInt(test.ONE_GWEI*2), 100000))
	test.IfNotNil(t, depErr, fmt.Sprintf("Error depositing funds from member to ether token: %v", depErr))

	context.Blockchain.Commit()

	ethBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser3.From)
	if ethBal.Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("Expected member bal to be > 0, got: %v", ethBal)
	}

	// the datatrust portion of the 'reserve' currently
	dataEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)

	// member needs to have approved datatrust to spend ether token
	_, approveErr := deployed.EtherTokenContract.Approve(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), deployed.DatatrustAddress, big.NewInt(test.ONE_WEI)) // 1 ETH
	test.IfNotNil(t, approveErr, fmt.Sprintf("Error approving spender: %v", approveErr))

	context.Blockchain.Commit()

	// note the allowance as later purchasing should decrease it
	ethAllow, _ := deployed.EtherTokenContract.Allowance(nil, context.AuthUser3.From, deployed.DatatrustAddress)

	_, purchaseErr := deployed.DatatrustContract.PurchaseByteCredits(test.GetTxOpts(
		context.AuthUser3, nil, big.NewInt(test.ONE_GWEI*2), 100000), big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, purchaseErr, fmt.Sprintf("Error buying byte credits: %v", purchaseErr))

	context.Blockchain.Commit()

	// user should now have some bytecredit
	bytesBalNow, _ := deployed.DatatrustContract.GetByteCredits(nil, context.AuthUser3.From)
	if bytesBalNow.Cmp(bytesBal) != 1 {
		t.Errorf("Expected byte credit %v to be more than %v", bytesBalNow, bytesBal)
	}

	// ether token balances should be updated
	ethBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser3.From)
	if ethBalNow.Cmp(ethBal) != -1 {
		t.Fatalf("Expected %v to be < %v", ethBalNow, ethBal)
	}

	ethAllowNow, _ := deployed.EtherTokenContract.Allowance(nil, context.AuthUser3.From, deployed.DatatrustAddress)
	if ethAllowNow.Cmp(ethAllow) != -1 {
		t.Fatalf("Expected %v to be < %v", ethAllowNow, ethAllow)
	}

	dataEthBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)
	if dataEthBalNow.Cmp(dataEthBal) != 1 {
		t.Fatalf("Expected %v to be > %v", dataEthBalNow, dataEthBal)
	}
}
