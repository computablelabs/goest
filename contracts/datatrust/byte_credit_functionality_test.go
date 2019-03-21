package datatrust

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
)

func TestPurchaseByteCredits(t *testing.T) {
	// user has no credits atm
	bytesBal, _ := deployed.DatatrustContract.GetByteCredits(nil, context.AuthMember3.From)

	// make a deposit in ETH, resulting in a 1:1 ethertoken balance
	_, depErr := deployed.EtherTokenContract.Deposit(&bind.TransactOpts{
		From:     context.AuthMember3.From,
		Signer:   context.AuthMember3.Signer,
		Value:    big.NewInt(ONE_WEI), // depositing one eth in wei
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 100000,
	})

	if depErr != nil {
		t.Fatalf("Error depositing funds from member to ether token: %v", depErr)
	}

	context.Blockchain.Commit()

	ethBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthMember3.From)
	if ethBal.Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("Expected member bal to be > 0, got: %v", ethBal)
	}

	// the datatrust portion of the 'reserve' currently
	dataEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)

	// member needs to have approved datatrust to spend ether token
	_, approveErr := deployed.EtherTokenContract.Approve(&bind.TransactOpts{
		From:     context.AuthMember3.From,
		Signer:   context.AuthMember3.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 100000,
	}, deployed.DatatrustAddress, big.NewInt(ONE_WEI)) // 1 ETH

	if approveErr != nil {
		t.Fatalf("Error approving spender: %v", approveErr)
	}

	context.Blockchain.Commit()

	// note the allowance as later purchasing should decrease it
	ethAllow, _ := deployed.EtherTokenContract.Allowance(nil, context.AuthMember3.From, deployed.DatatrustAddress)

	_, purchaseErr := deployed.DatatrustContract.PurchaseByteCredits(&bind.TransactOpts{
		From:     context.AuthMember3.From,
		Signer:   context.AuthMember3.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 100000,
	}, big.NewInt(ONE_GWEI))

	if purchaseErr != nil {
		t.Fatalf("Error buying byte credits: %v", purchaseErr)
	}

	context.Blockchain.Commit()

	// user should now have some bytecredit
	bytesBalNow, _ := deployed.DatatrustContract.GetByteCredits(nil, context.AuthMember3.From)
	if bytesBalNow.Cmp(bytesBal) != 1 {
		t.Fatalf("Expected byte credit %v to be more than %v", bytesBalNow, bytesBal)
	}

	// ether token balances should be updated
	ethBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthMember3.From)
	if ethBalNow.Cmp(ethBal) != -1 {
		t.Fatalf("Expected %v to be < %v", ethBalNow, ethBal)
	}

	ethAllowNow, _ := deployed.EtherTokenContract.Allowance(nil, context.AuthMember3.From, deployed.DatatrustAddress)
	if ethAllowNow.Cmp(ethAllow) != -1 {
		t.Fatalf("Expected %v to be < %v", ethAllowNow, ethAllow)
	}

	dataEthBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)
	if dataEthBalNow.Cmp(dataEthBal) != 1 {
		t.Fatalf("Expected %v to be > %v", dataEthBalNow, dataEthBal)
	}
}

func TestPurchaseBytes(t *testing.T) {
	// var x big.Int
	// and client (d)app can easily estimate how many can be bought (with fee)
	cost, _ := deployed.ParameterizerContract.GetCostPerByte(nil)
	backendPct, _ := deployed.ParameterizerContract.GetBackendPayment(nil) // a % of 100
	userCredits, _ := deployed.DatatrustContract.GetByteCredits(nil, context.AuthMember3.From)
	// convert them to int64s
	cost64 := cost.Int64()
	backendPct64 := backendPct.Int64()
	userCredits64 := userCredits.Int64()

	// a (d)app ui could easily help a user figure out their purchaseable amount - 833333 in this case
	subTotal := 833333 * cost64
	feeForN := subTotal / (100 / backendPct64)
	estimatedCredits := userCredits64 - (subTotal + feeForN) // should be the remainder of credits

	_, purchaseErr := deployed.DatatrustContract.PurchaseBytes(&bind.TransactOpts{
		From:     context.AuthBackend.From,
		Signer:   context.AuthBackend.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 100000,
	}, context.AuthMember3.From, big.NewInt(833333))

	if purchaseErr != nil {
		t.Fatal("Error purchasing bytes")
	}

	context.Blockchain.Commit()

	// would should see credits adjusted
	userCreditsNow, _ := deployed.DatatrustContract.GetByteCredits(nil, context.AuthMember3.From)
	if userCreditsNow.Cmp(big.NewInt(estimatedCredits)) != 0 {
		t.Fatalf("Expected %v to be %v", userCreditsNow, estimatedCredits)
	}
}
