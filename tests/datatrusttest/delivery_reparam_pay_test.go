package datatrusttest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
	"time"
)

func TestRequestDeliveryReparamPay(t *testing.T) {

	// Record parameters at start before reparam
	costPerByte, _ := deployed.ParameterizerContract.GetCostPerByte(nil)
	backendPayment, _ := deployed.ParameterizerContract.GetBackendPayment(nil)
	requested := big.NewInt(1024 * 1024)

	// user has no credits atm
	purchased, _ := deployed.DatatrustContract.GetBytesPurchased(nil, context.AuthUser2.From)

	// reserve balance atm. This will go up by the res_payment after a request...
	resBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.ReserveAddress)

	// make a deposit in ETH, resulting in a 1:1 ethertoken balance
	_, depErr := deployed.EtherTokenContract.Deposit(test.GetTxOpts(context.AuthUser2,
		big.NewInt(test.ONE_GWEI*10), big.NewInt(test.ONE_GWEI*2), 100000))
	test.IfNotNil(t, depErr, fmt.Sprintf("Error depositing funds from member to ether token: %v", depErr))
	context.Blockchain.Commit()

	// user now has an ether token balance
	ethBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser2.From)
	if ethBal.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected member bal to be > 0, got: %v", ethBal)
	}

	// member needs to have approved datatrust to spend ether token
	_, approveErr := deployed.EtherTokenContract.Approve(test.GetTxOpts(context.AuthUser2, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), deployed.DatatrustAddress, big.NewInt(test.ONE_GWEI*10))
	test.IfNotNil(t, approveErr, fmt.Sprintf("Error approving spender: %v", approveErr))
	context.Blockchain.Commit()

	// when a delivery is requested, a hash of the query datatrust will recieve is expected
	query := test.GenBytes32("select * from SPAM where EGGS eq TRUE")

	// assure no delivery exists for this query
	owner, req, _, _ := deployed.DatatrustContract.GetDelivery(nil, query)
	if owner == context.AuthUser2.From {
		t.Error("Expected no delivery object yet")
	}

	if req.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected bytes requested to be 0, got: %v", req)
	}

	// the actual request for a delivery...
	_, delErr := deployed.DatatrustContract.RequestDelivery(test.GetTxOpts(
		context.AuthUser2, nil, big.NewInt(test.ONE_GWEI*2), 250000), query, requested) // ~1MB
	test.IfNotNil(t, delErr, fmt.Sprintf("Error requesting delivery: %v", delErr))
	context.Blockchain.Commit()

	// user should now have some bytes
	purchasedNow, _ := deployed.DatatrustContract.GetBytesPurchased(nil, context.AuthUser2.From)
	if purchasedNow.Cmp(purchased) != 1 {
		t.Errorf("Expected byte balance %v to be more than %v", purchasedNow, purchased)
	}

	// reserve gets its share when delivery is requested
	resBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.ReserveAddress)
	if resBalNow.Cmp(resBal) != 1 {
		t.Errorf("Expected %v to be > %v", resBalNow, resBal)
	}

	// delivery object should be present
	ownerNow, reqNow, _, _ := deployed.DatatrustContract.GetDelivery(nil, query)
	if ownerNow != context.AuthUser2.From {
		t.Error("Expected delivery object to be owned by user 2")
	}
	if reqNow.Cmp(big.NewInt(1024*1024)) != 0 {
		t.Errorf("Expected bytes requested to be 1MB, got: %v", reqNow)
	}

	// auth memeber must stake reparam...
	transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser1.From,
		big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr, "Error maybe transferring market tokens")

	// member will need to have approved the voting contract to spend at least the stake
	incErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser1, deployed.VotingAddress,
		big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, incErr, "Error maybe transferring market token approval")

	// Invoke reparameterize
	_, err := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 500000), test.COST_PER_BYTE, big.NewInt(test.ONE_ETH))
	test.IfNotNil(t, err, fmt.Sprintf("Error creating proposal: %v", err))
	context.Blockchain.Commit()

	// we should see a candidate now
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(11), big.NewInt(test.ONE_ETH))
	isCan, _ := deployed.VotingContract.IsCandidate(nil, paramHash)

	if !isCan {
		t.Errorf("Expected isCandidate to be true, got: %v", isCan)
	}

	// Vote for the candidate
	// member may need funding for the vote
	test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser3.From, big.NewInt(test.ONE_GWEI))

	// member will need to have approved the voting contract to spend
	test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser3, deployed.VotingAddress, big.NewInt(test.ONE_GWEI))

	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), paramHash, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))
	context.Blockchain.Commit()

	// Check vote was made
	_, _, _, _, yea, _, _ := deployed.VotingContract.GetCandidate(nil, paramHash)

	if yea.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected number of votes to be > 0, got: %v", yea)
	}

	// move time forward so the poll is closed
	context.Blockchain.AdjustTime(test.MIN_VOTE_BY * time.Second)
	context.Blockchain.Commit()

	// make sure its closed now
	closed, _ := deployed.VotingContract.PollClosed(nil, paramHash)

	if closed != true {
		t.Errorf("Expected pollClosed to be true, got: %v", closed)
	}

	// resolve the reparam
	_, resolveErr := deployed.ParameterizerContract.ResolveReparam(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), paramHash)
	test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving proposal: %v", resolveErr))

	context.Blockchain.Commit()

	// cost_per_byte should be changed now
	newCostPerByte, _ := deployed.ParameterizerContract.GetCostPerByte(nil)
	if newCostPerByte.Cmp(big.NewInt(test.ONE_ETH)) != 0 {
		t.Errorf("Expected the new cost_per_byte param to be ONE_ETH, got: %v", newCostPerByte)
	}

	// a backend must be registered (will have to stake now)
	transErr2 := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthBackend.From,
		big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr2, "Error maybe transferring market tokens")

	// member will need to have approved the voting contract to spend at least the stake
	incErr2 := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthBackend, deployed.VotingAddress,
		big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, incErr2, "Error maybe transferring market token approval")

	// Let's register the backend
	_, regErr := deployed.DatatrustContract.Register(test.GetTxOpts(context.AuthBackend, nil,
		big.NewInt(test.ONE_GWEI*2), 500000), "https://www.imyerbackend.io")
	test.IfNotNil(t, regErr, fmt.Sprintf("Error registering for backend status: %v", regErr))
	context.Blockchain.Commit()

	hash, _ := deployed.DatatrustContract.GetHash(nil, "https://www.imyerbackend.io")

	// auth member will need at least the stake
	transErr3 := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser1.From,
		big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr3, "Error maybe transferring market tokens")

	// member will need to have approved the voting contract to spend at least the stake
	incErr3 := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser1, deployed.VotingAddress,
		big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, incErr3, "Error maybe transferring market token approval")

	// vote to approve the backend
	_, voteErr2 := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), hash, big.NewInt(1))
	test.IfNotNil(t, voteErr2, fmt.Sprintf("Error voting for candidate: %v", voteErr2))
	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(test.MIN_VOTE_BY * time.Second)
	context.Blockchain.Commit()

	// make it official
	_, resolveErr2 := deployed.DatatrustContract.ResolveRegistration(test.GetTxOpts(
		context.AuthUser1, nil, big.NewInt(test.ONE_GWEI*2), 1000000), hash)
	test.IfNotNil(t, resolveErr2, fmt.Sprintf("Error resolving application: %v", resolveErr2))
	context.Blockchain.Commit()

	// get a listing up
	listingHash := test.GenBytes32("LookAtMyJunk")
	_, listErr := deployed.ListingContract.List(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 250000), listingHash)
	test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
	context.Blockchain.Commit()

	// a listing will not be accepted without a data-hash
	dataHash := test.GenBytes32("thedata")
	_, dataErr := deployed.DatatrustContract.SetDataHash(test.GetTxOpts(context.AuthBackend, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), listingHash, dataHash)
	test.IfNotNil(t, dataErr, "Error setting data hash for listing")
	context.Blockchain.Commit()

	// cast a vote for the listing, voter may need funds...
	transErr4 := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner,
		context.AuthUser1.From, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr4, fmt.Sprintf("Error transferring tokens to member: %v", transErr4))
	context.Blockchain.Commit()
	// member will need to have approved the voting contract to spend
	incErr4 := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser1,
		deployed.VotingAddress, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, incErr4, fmt.Sprintf("Error approving market contract to spend: %v", incErr4))
	context.Blockchain.Commit()

	// yay vote...
	_, voteErr3 := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, big.NewInt(1))
	test.IfNotNil(t, voteErr3, fmt.Sprintf("Error voting for candidate: %v", voteErr3))
	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(test.MIN_VOTE_BY * time.Second)
	context.Blockchain.Commit()

	// any stakeholder can call for resolution - now it becomes a listing
	_, resolveErr3 := deployed.ListingContract.ResolveApplication(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, resolveErr3, fmt.Sprintf("Error resolving application: %v", resolveErr3))
	context.Blockchain.Commit()

	// current bytes purchased, should decrease with listing access reporting
	bytesBal, _ := deployed.DatatrustContract.GetBytesPurchased(nil, context.AuthUser2.From)

	_, accErr := deployed.DatatrustContract.ListingAccessed(test.GetTxOpts(context.AuthBackend, nil,
		// let's say one listing was used for the whole request
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, query, big.NewInt(1024*1024))
	test.IfNotNil(t, accErr, "Error claiming listing accessed")
	context.Blockchain.Commit()

	// as listings are accessed the user's bytes purchased bal decreases...
	bytesBalNow, _ := deployed.DatatrustContract.GetBytesPurchased(nil, context.AuthUser2.From)
	if bytesBalNow.Cmp(bytesBal) != -1 {
		t.Errorf("Expected %v to be > %v", bytesBal, bytesBalNow)
	}

	// we should see a delivery object whose bytes delivered match the bytes requested
	_, req, del, _ := deployed.DatatrustContract.GetDelivery(nil, query)
	if req.Cmp(del) != 0 {
		t.Errorf("Expected %v to be %v", req, del)
	}

	// the user should have no avail bytes at this time
	purchasedFinal, _ := deployed.DatatrustContract.GetBytesPurchased(nil, context.AuthUser2.From)
	if purchasedFinal.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("expected %v to be 0", purchasedFinal)
	}

	// note the ether token balances
	beAddr, _ := deployed.DatatrustContract.GetBackendAddress(nil)
	beBal, _ := deployed.EtherTokenContract.BalanceOf(nil, beAddr)
	dtBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)

	// Let's try calling delivered now
	_, delErr2 := deployed.DatatrustContract.Delivered(test.GetTxOpts(context.AuthBackend, nil,
		big.NewInt(test.ONE_GWEI*2), 250000), query, query)
	test.IfNotNil(t, delErr2, "Error calling delivered")
	context.Blockchain.Commit()

	// should not see balance changes now (delivery failed since reparam set costs too high)
	beBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, beAddr)
	dtBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)

	// newest bal will have increased
	if beBalNow.Cmp(beBal) != 1 {
		t.Errorf("Expected %v to be %v", beBalNow, beBal)
	}

	// Compute theoretic backend payment
	// (costPerByte * requested * backendPayment) / 100
	bePayment := big.NewInt(1)
	bePayment = bePayment.Mul(bePayment, costPerByte)
	bePayment = bePayment.Mul(bePayment, requested)
	bePayment = bePayment.Mul(bePayment, backendPayment)
	bePayment = bePayment.Div(bePayment, big.NewInt(100))
	// Here is the actual backend payment. They should match
	beActualPayment := beBalNow.Sub(beBalNow, beBal)
	if beActualPayment.Cmp(bePayment) != 0 {
		t.Errorf("Expected %v to be %v", beActualPayment, bePayment)
	}

	// newest bal will have decreased
	if dtBalNow.Cmp(dtBal) != -1 {
		t.Errorf("Expected %v to be %v", dtBalNow, dtBal)
	}

}
