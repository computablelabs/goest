package scenariofive

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
)

func oneHundredOneEth() *big.Int {
	one := big.NewInt(test.ONE_WEI)
	oneHundred := one.Mul(one, big.NewInt(100))
	return oneHundred.Add(oneHundred, big.NewInt(test.ONE_WEI))
}

func oneHundredEth() *big.Int {
	one := big.NewInt(test.ONE_WEI)
	return one.Mul(one, big.NewInt(100))
}

func TestRequestDelivery(t *testing.T) {
	// user has no credits atm
	purchased, _ := deployed.DatatrustContract.GetBytesPurchased(nil, context.AuthUser3.From)
	// reserve balance atm. This will go up by the res_payment after a request...
	resBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.InvestingAddress)

	// make a deposit in ETH, resulting in a 1:1 ethertoken balance
	_, depErr := deployed.EtherTokenContract.Deposit(test.GetTxOpts(context.AuthUser3,
		big.NewInt(test.ONE_GWEI*10), big.NewInt(test.ONE_GWEI*2), 100000))
	test.IfNotNil(t, depErr, fmt.Sprintf("Error depositing funds from member to ether token: %v", depErr))

	context.Blockchain.Commit()

	// user now has an ether token balance
	ethBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser3.From)
	if ethBal.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected member bal to be > 0, got: %v", ethBal)
	}

	// purchase payments are banked by the datatrust contract at the ether token (-res_payment)
	dataEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)

	// member needs to have approved datatrust to spend ether token
	_, approveErr := deployed.EtherTokenContract.Approve(test.GetTxOpts(context.AuthUser3, nil,
		//big.NewInt(test.ONE_GWEI*2), 100000), deployed.DatatrustAddress, big.NewInt(test.ONE_GWEI*10))
		big.NewInt(test.ONE_GWEI*2), 100000), deployed.DatatrustAddress, big.NewInt(test.ONE_WEI))
	test.IfNotNil(t, approveErr, fmt.Sprintf("Error approving spender: %v", approveErr))

	context.Blockchain.Commit()

	// note the allowance as later purchasing should decrease it
	ethAllow, _ := deployed.EtherTokenContract.Allowance(nil, context.AuthUser3.From, deployed.DatatrustAddress)

	// when a delivery is requested, a hash of the query datatrust will recieve is expected
	query := test.GenBytes32("select * from SPAM where EGGS eq TRUE")

	// assure no delivery exists for this query
	owner, req, _, _ := deployed.DatatrustContract.GetDelivery(nil, query)
	if owner == context.AuthUser3.From {
		t.Error("Expected no delivery object yet")
	}

	if req.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected bytes requested to be 0, got: %v", req)
	}

	// the actual request for a delivery...
	_, delErr := deployed.DatatrustContract.RequestDelivery(test.GetTxOpts(
		//context.AuthUser3, nil, big.NewInt(test.ONE_GWEI*2), 250000), query, big.NewInt(1024*1024)) // ~1MB
		context.AuthUser3, nil, big.NewInt(test.ONE_GWEI*2), 250000), query, big.NewInt(1024)) // ~1MB
	test.IfNotNil(t, delErr, fmt.Sprintf("Error requesting delivery: %v", delErr))

	context.Blockchain.Commit()

	// user should now have some bytes
	purchasedNow, _ := deployed.DatatrustContract.GetBytesPurchased(nil, context.AuthUser3.From)
	if purchasedNow.Cmp(purchased) != 1 {
		t.Errorf("Expected byte balance %v to be more than %v", purchasedNow, purchased)
	}

	// ether token balances should be updated
	ethBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser3.From)
	if ethBalNow.Cmp(ethBal) != -1 {
		t.Errorf("Expected %v to be < %v", ethBalNow, ethBal)
	}

	// allowances also reflect spending...
	ethAllowNow, _ := deployed.EtherTokenContract.Allowance(nil, context.AuthUser3.From, deployed.DatatrustAddress)
	if ethAllowNow.Cmp(ethAllow) != -1 {
		t.Errorf("Expected %v to be < %v", ethAllowNow, ethAllow)
	}

	// funds should be banked by the datatrust now
	dataEthBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)
	if dataEthBalNow.Cmp(dataEthBal) != 1 {
		t.Errorf("Expected %v to be > %v", dataEthBalNow, dataEthBal)
	}

	// reserve gets its share when delivery is requested
	resBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.InvestingAddress)
	if resBalNow.Cmp(resBal) != 1 {
		t.Errorf("Expected %v to be > %v", resBalNow, resBal)
	}

	// at this point the sum of what went into the reserve + the amount locked in datatrust will == bytesPurchased * cost_per_byte
	toRes := resBalNow.Sub(resBalNow, resBal)
	summed := toRes.Add(toRes, dataEthBalNow)
	cost, _ := deployed.ParameterizerContract.GetCostPerByte(nil)
	bytesCost := purchasedNow.Mul(purchasedNow, cost)
	if summed.Cmp(bytesCost) != 0 {
		t.Errorf("Expected %v to be %v", summed, bytesCost)
	}

	// delivery object should be present
	ownerNow, reqNow, _, _ := deployed.DatatrustContract.GetDelivery(nil, query)
	if ownerNow != context.AuthUser3.From {
		t.Error("Expected delivery object to be owned by user 3")
	}
	//if reqNow.Cmp(big.NewInt(1024*1024)) != 0 {
	if reqNow.Cmp(big.NewInt(1024)) != 0 {
		t.Errorf("Expected bytes requested to be 1MB, got: %v", reqNow)
	}
}
