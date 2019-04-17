package datatrusttest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
	"time"
)

func TestRequestDelivery(t *testing.T) {
	// user has no credits atm
	bytesBal, _ := deployed.DatatrustContract.GetByteCredits(nil, context.AuthUser3.From)
	// reserve balance atm. This will go up by the res_payment after a request...
	resBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.InvestingAddress)

	// make a deposit in ETH, resulting in a 1:1 ethertoken balance
	_, depErr := deployed.EtherTokenContract.Deposit(test.GetTxOpts(context.AuthUser3,
		big.NewInt(test.ONE_GWEI*10), big.NewInt(test.ONE_GWEI*2), 100000))
	test.IfNotNil(t, depErr, fmt.Sprintf("Error depositing funds from member to ether token: %v", depErr))

	context.Blockchain.Commit()

	ethBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser3.From)
	if ethBal.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected member bal to be > 0, got: %v", ethBal)
	}

	// any unused byte credits are banked by the datatrust contract at the ether token (-res_payment)
	dataEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)

	// member needs to have approved datatrust to spend ether token
	_, approveErr := deployed.EtherTokenContract.Approve(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), deployed.DatatrustAddress, big.NewInt(test.ONE_GWEI*10))
	test.IfNotNil(t, approveErr, fmt.Sprintf("Error approving spender: %v", approveErr))

	// transaction cost is (1024*1024) requested * 1000 (cost_per_byte) = 1048576000 wei

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

	_, delErr := deployed.DatatrustContract.RequestDelivery(test.GetTxOpts(
		context.AuthUser3, nil, big.NewInt(test.ONE_GWEI*2), 250000), query, big.NewInt(1024*1024)) // ~1MB
	test.IfNotNil(t, delErr, fmt.Sprintf("Error requesting delivery: %v", delErr))

	context.Blockchain.Commit()

	// user should now have some bytecredit
	bytesBalNow, _ := deployed.DatatrustContract.GetByteCredits(nil, context.AuthUser3.From)
	if bytesBalNow.Cmp(bytesBal) != 1 {
		t.Errorf("Expected byte credit %v to be more than %v", bytesBalNow, bytesBal)
	}

	// ether token balances should be updated
	ethBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser3.From)
	if ethBalNow.Cmp(ethBal) != -1 {
		t.Errorf("Expected %v to be < %v", ethBalNow, ethBal)
	}

	ethAllowNow, _ := deployed.EtherTokenContract.Allowance(nil, context.AuthUser3.From, deployed.DatatrustAddress)
	if ethAllowNow.Cmp(ethAllow) != -1 {
		t.Errorf("Expected %v to be < %v", ethAllowNow, ethAllow)
	}

	dataEthBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)
	if dataEthBalNow.Cmp(dataEthBal) != 1 {
		t.Errorf("Expected %v to be > %v", dataEthBalNow, dataEthBal)
	}

	resBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.InvestingAddress)
	if resBalNow.Cmp(resBal) != 1 {
		t.Errorf("Expected %v to be > %v", resBalNow, resBal)
	}

	// at this point the sum of what went into the reserve + the amount locked in datatrust will == byte_credits
	toRes := resBalNow.Sub(resBalNow, resBal)
	summed := toRes.Add(toRes, dataEthBalNow)
	if summed.Cmp(bytesBalNow) != 0 {
		t.Errorf("Expected %v to be %v", summed, bytesBalNow)
	}

	// delivery object should be present
	ownerNow, reqNow, _, _ := deployed.DatatrustContract.GetDelivery(nil, query)
	if ownerNow != context.AuthUser3.From {
		t.Error("Expected delivery object to be owned by user 3")
	}
	if reqNow.Cmp(big.NewInt(1024*1024)) != 0 {
		t.Errorf("Expected bytes requested to be 1MB, got: %v", reqNow)
	}
}

func TestListingAccessed(t *testing.T) {
	// a backend must be registered
	_, regErr := deployed.DatatrustContract.Register(test.GetTxOpts(context.AuthBackend, nil,
		big.NewInt(test.ONE_GWEI*2), 500000), "https://www.imyerbackend.io")
	test.IfNotNil(t, regErr, fmt.Sprintf("Error registering for backend status: %v", regErr))
	context.Blockchain.Commit()

	hash, _ := deployed.DatatrustContract.GetHash(nil, "https://www.imyerbackend.io")

	// auth member will need at least the stake
	transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthFactory, context.AuthUser1.From,
		big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr, "Error maybe transferring market tokens")

	// member will need to have approved the voting contract to spend at least the stake
	incErr := test.MaybeIncreaseMarketTokenApproval(context, deployed, context.AuthUser1, deployed.VotingAddress,
		big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, incErr, "Error maybe transferring market token approval")

	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), hash, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))
	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	_, resolveErr := deployed.DatatrustContract.ResolveRegistration(test.GetTxOpts(
		context.AuthUser1, nil, big.NewInt(test.ONE_GWEI*2), 1000000), hash)
	test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving application: %v", resolveErr))
	context.Blockchain.Commit()

	// get a listing up
	listingHash := test.GenBytes32("LookAtMyJunk")
	_, listErr := deployed.ListingContract.List(test.GetTxOpts(context.AuthUser2, nil,
		big.NewInt(test.ONE_GWEI*2), 250000), listingHash)
	test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
	context.Blockchain.Commit()

	dataHash := test.GenBytes32("thedata")
	_, dataErr := deployed.DatatrustContract.SetDataHash(test.GetTxOpts(context.AuthBackend, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), listingHash, dataHash)
	test.IfNotNil(t, dataErr, "Error setting data hash for listing")
	context.Blockchain.Commit()

	// cast a vote for, voter may need funds...
	transErr2 := test.MaybeTransferMarketToken(context, deployed, context.AuthFactory,
		context.AuthUser1.From, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr2, fmt.Sprintf("Error transferring tokens to member: %v", transErr2))
	context.Blockchain.Commit()
	// member will need to have approved the voting contract to spend
	appErr := test.MaybeIncreaseMarketTokenApproval(context, deployed, context.AuthUser1,
		deployed.VotingAddress, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, appErr, fmt.Sprintf("Error approving market contract to spend: %v", appErr))
	context.Blockchain.Commit()

	_, voteErr2 := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, big.NewInt(1))
	test.IfNotNil(t, voteErr2, fmt.Sprintf("Error voting for candidate: %v", voteErr2))
	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	// any council member can call for resolution
	_, resolveErr2 := deployed.ListingContract.ResolveApplication(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, resolveErr2, fmt.Sprintf("Error resolving application: %v", resolveErr2))
	context.Blockchain.Commit()

	// with a listing in place we can claim that it was accessed
	query := test.GenBytes32("select * from SPAM where EGGS eq TRUE")
	// current byte credits, should decrease with listing access reporting
	bytesBal, _ := deployed.DatatrustContract.GetByteCredits(nil, context.AuthUser3.From)
	// current access_credits for listiing, should increase with reporting
	accessBal, _ := deployed.DatatrustContract.GetAccessCredits(nil, listingHash)
	// none delivered yet...
	_, _, delivered, _ := deployed.DatatrustContract.GetDelivery(nil, query)

	_, accErr := deployed.DatatrustContract.ListingAccessed(test.GetTxOpts(context.AuthBackend, nil,
		// let's say one listing was used for 1/2 the request
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, query, big.NewInt(1024*512))
	test.IfNotNil(t, accErr, "Error claiming listing accessed")
	context.Blockchain.Commit()

	bytesBalNow, _ := deployed.DatatrustContract.GetByteCredits(nil, context.AuthUser3.From)
	if bytesBalNow.Cmp(bytesBal) != -1 {
		t.Errorf("Expected %v to be > %v", bytesBal, bytesBalNow)
	}
	// current access_credits for listiing, should increase with reporting
	accessBalNow, _ := deployed.DatatrustContract.GetAccessCredits(nil, listingHash)
	if accessBalNow.Cmp(accessBal) != 1 {
		t.Errorf("Expected %v to be > %v", accessBalNow, accessBal)
	}
	// half has been delivered
	_, _, deliveredNow, _ := deployed.DatatrustContract.GetDelivery(nil, query)
	if deliveredNow.Cmp(delivered) != 1 {
		t.Errorf("Expected %v to be > %v", deliveredNow, delivered)
	}
}

func TestDeliveredThatShouldFail(t *testing.T) {
	// note the ether token balances
	beAddr, _ := deployed.DatatrustContract.GetBackendAddress(nil)
	beBal, _ := deployed.EtherTokenContract.BalanceOf(nil, beAddr)
	dtBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)
	query := test.GenBytes32("select * from SPAM where EGGS eq TRUE")
	_, req, del, _ := deployed.DatatrustContract.GetDelivery(nil, query)
	// assure its not delivered
	if req.Cmp(del) != 1 {
		t.Errorf("Expected %v to be > %v", req, del)
	}
	// we can attempt to call delivered for a delivery object that has not been 'filled'
	urlHash := test.GenBytes32("someURL.net/lumberjack")
	_, delErr := deployed.DatatrustContract.Delivered(test.GetTxOpts(context.AuthBackend, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), query, urlHash)
	test.IfNotNil(t, delErr, "Error calling delivered")
	context.Blockchain.Commit()
	// no balance changes should have occurred
	beBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, beAddr)
	dtBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)

	if beBal.Cmp(beBalNow) != 0 {
		t.Errorf("Expected %v to be %v", beBal, beBalNow)
	}

	if dtBal.Cmp(dtBalNow) != 0 {
		t.Errorf("Expected %v to be %v", dtBal, dtBalNow)
	}
}

func TestDelivered(t *testing.T) {
	// we'll put up another listing to claim was accessed
	listingHash := test.GenBytes32("LookAtMyJunkToo")
	_, listErr := deployed.ListingContract.List(test.GetTxOpts(context.AuthUser2, nil,
		big.NewInt(test.ONE_GWEI*2), 250000), listingHash)
	test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
	context.Blockchain.Commit()

	dataHash := test.GenBytes32("moardata")
	_, dataErr := deployed.DatatrustContract.SetDataHash(test.GetTxOpts(context.AuthBackend, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), listingHash, dataHash)
	test.IfNotNil(t, dataErr, "Error setting data hash for listing")
	context.Blockchain.Commit()

	// cast a vote for, voter may need funds...
	transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthFactory,
		context.AuthUser1.From, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr, fmt.Sprintf("Error transferring tokens to member: %v", transErr))
	context.Blockchain.Commit()
	// member will need to have approved the voting contract to spend
	appErr := test.MaybeIncreaseMarketTokenApproval(context, deployed, context.AuthUser1,
		deployed.VotingAddress, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, appErr, fmt.Sprintf("Error approving market contract to spend: %v", appErr))
	context.Blockchain.Commit()

	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))
	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	// any council member can call for resolution
	_, resolveErr := deployed.ListingContract.ResolveApplication(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving application: %v", resolveErr))
	context.Blockchain.Commit()

	// with that in place we can claim it as the source of the rest of the delivery
	query := test.GenBytes32("select * from SPAM where EGGS eq TRUE")
	_, accErr := deployed.DatatrustContract.ListingAccessed(test.GetTxOpts(context.AuthBackend, nil,
		// let's say one listing was used for the other 1/2 of the request
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, query, big.NewInt(1024*512))
	test.IfNotNil(t, accErr, "Error claiming listing accessed")
	context.Blockchain.Commit()

	// we should see a delivery object whose bytes delivered match the bytes requested
	_, req, del, _ := deployed.DatatrustContract.GetDelivery(nil, query)
	if req.Cmp(del) != 0 {
		t.Errorf("Expected %v to be %v", req, del)
	}

	// we should see equal access credits for both listings
	ogListingHash := test.GenBytes32("LookAtMyJunk")
	ogAccessBal, _ := deployed.DatatrustContract.GetAccessCredits(nil, ogListingHash)
	accessBal, _ := deployed.DatatrustContract.GetAccessCredits(nil, listingHash)
	if accessBal.Cmp(ogAccessBal) != 0 {
		t.Errorf("Expected %v to be %v", accessBal, ogAccessBal)
	}

	// note the ether token balances
	beAddr, _ := deployed.DatatrustContract.GetBackendAddress(nil)
	beBal, _ := deployed.EtherTokenContract.BalanceOf(nil, beAddr)
	dtBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)

	urlHash := test.GenBytes32("someURL.net/roger/shrubber")
	_, delErr := deployed.DatatrustContract.Delivered(test.GetTxOpts(context.AuthBackend, nil,
		big.NewInt(test.ONE_GWEI*2), 250000), query, urlHash)
	test.IfNotNil(t, delErr, "Error calling delivered")
	context.Blockchain.Commit()
	// should see balance changes now
	beBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, beAddr)
	dtBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)

	// newest bal will have increased
	if beBalNow.Cmp(beBal) != 1 {
		t.Errorf("Expected %v to be > %v", beBalNow, beBal)
	}

	// newest bal will have decreased
	if dtBalNow.Cmp(dtBal) != -1 {
		t.Errorf("Expected %v to be < %v", dtBalNow, dtBal)
	}
}
