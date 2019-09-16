package dustingtest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
	"time"
)

func TestRequestDelivery(t *testing.T) {
	// user has no credits atm
	t.Log("User3 as demand user")
	purchased, _ := deployed.DatatrustContract.GetBytesPurchased(nil, context.AuthUser3.From)
	t.Logf("User3 current bytes purchased balance: %v", purchased)
	// reserve balance atm. This will go up by the res_payment after a request...
	resBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.ReserveAddress)
	t.Logf("Current reserve balance: %v", resBal)

	// make a deposit in ETH, resulting in a 1:1 ethertoken balance
	_, depErr := deployed.EtherTokenContract.Deposit(test.GetTxOpts(context.AuthUser3,
		big.NewInt(test.ONE_ETH*2), big.NewInt(test.ONE_GWEI*2), 100000))
	test.IfNotNil(t, depErr, fmt.Sprintf("Error depositing funds from member to ether token: %v", depErr))

	context.Blockchain.Commit()

	// user now has an ether token balance
	ethBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser3.From)
	if ethBal.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected member bal to be > 0, got: %v", ethBal)
	}
	t.Logf("User3 ether token balance: %v", test.Commafy(ethBal))

	// purchase payments are banked by the datatrust contract at the ether token (-res_payment)
	dataEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)
	t.Logf("Current ETH banked by the Datatrust contract: %v", dataEthBal)

	// member needs to have approved datatrust to spend ether token
	_, approveErr := deployed.EtherTokenContract.Approve(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), deployed.DatatrustAddress, big.NewInt(test.ONE_ETH*2))
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
	t.Log("User requests delivery for ~10MB")
	_, delErr := deployed.DatatrustContract.RequestDelivery(test.GetTxOpts(
		context.AuthUser3, nil, big.NewInt(test.ONE_GWEI*2), 250000), query, big.NewInt((1024*1024)*10)) // ~10MB
	test.IfNotNil(t, delErr, fmt.Sprintf("Error requesting delivery: %v", delErr))

	context.Blockchain.Commit()

	// user should now have some bytes
	purchasedNow, _ := deployed.DatatrustContract.GetBytesPurchased(nil, context.AuthUser3.From)
	if purchasedNow.Cmp(purchased) != 1 {
		t.Errorf("Expected byte balance %v to be more than %v", purchasedNow, purchased)
	}
	t.Logf("User3 bytes purchased balance after delivery request: %v", test.Commafy(purchasedNow))

	// ether token balances should be updated
	ethBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthUser3.From)
	if ethBalNow.Cmp(ethBal) != -1 {
		t.Errorf("Expected %v to be < %v", ethBalNow, ethBal)
	}
	t.Logf("User3 ether token balance post request: %v", test.Commafy(ethBalNow))

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
	t.Logf("ETH banked by the Datatrust contract post request: %v", test.Commafy(dataEthBalNow))

	// reserve gets its share when delivery is requested
	resBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.ReserveAddress)
	if resBalNow.Cmp(resBal) != 1 {
		t.Errorf("Expected %v to be > %v", resBalNow, resBal)
	}
	// t.Logf("Reserve balance post request: %v", test.Commafy(resBalNow))
	// at this point the sum of what went into the reserve + the amount locked in datatrust will == bytesPurchased * cost_per_byte
	toRes := resBalNow.Sub(resBalNow, resBal)
	summed := toRes.Add(toRes, dataEthBalNow)
	cost, _ := deployed.ParameterizerContract.GetCostPerByte(nil)
	bytesCost := purchasedNow.Mul(purchasedNow, cost)
	if summed.Cmp(bytesCost) != 0 {
		t.Errorf("Expected %v to be %v", summed, bytesCost)
	}

	// delivery object should be present
	ownerNow, reqNow, deliv, _ := deployed.DatatrustContract.GetDelivery(nil, query)
	if ownerNow != context.AuthUser3.From {
		t.Error("Expected delivery object to be owned by user 3")
	}
	if reqNow.Cmp(big.NewInt((1024*1024)*10)) != 0 {
		t.Errorf("Expected bytes requested to be ~10MB, got: %v", reqNow)
	}
	t.Logf("Delivery object bytes_requested: %v", test.Commafy(reqNow))
	t.Logf("Delivery object bytes_delivered: %v", deliv)
}

func TestListingAccessed(t *testing.T) {

	// auth backend will need at least the stake
	transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthBackend.From,
		big.NewInt(10000000000000000))
	test.IfNotNil(t, transErr, "Error maybe transferring market tokens")

	// backend will need to have approved the voting contract to spend at least the stake
	incErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthBackend, deployed.VotingAddress,
		big.NewInt(10000000000000000))
	test.IfNotNil(t, incErr, "Error maybe transferring market token approval")

	// a backend must be registered
	_, regErr := deployed.DatatrustContract.Register(test.GetTxOpts(context.AuthBackend, nil,
		big.NewInt(test.ONE_GWEI*2), 500000), "https://www.yermomsbackend.io")
	test.IfNotNil(t, regErr, fmt.Sprintf("Error registering for backend status: %v", regErr))
	context.Blockchain.Commit()

	hash, _ := deployed.DatatrustContract.GetHash(nil, "https://www.yermomsbackend.io")

	// auth member will need at least the stake
	transErrVote := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser1.From,
		big.NewInt(test.ONE_ETH))
	test.IfNotNil(t, transErrVote, "Error maybe transferring market tokens")

	// member will need to have approved the voting contract to spend at least the stake
	incErrVote := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser1, deployed.VotingAddress,
		big.NewInt(test.ONE_ETH))
	test.IfNotNil(t, incErrVote, "Error maybe transferring market token approval")

	// vote to approve the backend
	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), hash, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))
	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	// make it official
	_, resolveErr := deployed.DatatrustContract.ResolveRegistration(test.GetTxOpts(
		context.AuthUser1, nil, big.NewInt(test.ONE_GWEI*2), 1000000), hash)
	test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving application: %v", resolveErr))
	context.Blockchain.Commit()

	// voter can reclaim thier stake
	_, unErr := deployed.VotingContract.Unstake(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), hash)
	test.IfNotNil(t, unErr, fmt.Sprintf("Error Unstaking: %v", unErr))
	context.Blockchain.Commit()

	// get a listing up
	listingHash := test.GenBytes32("LookAtMyJunk")
	_, listErr := deployed.ListingContract.List(test.GetTxOpts(context.AuthUser2, nil,
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
	transErr2 := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner,
		context.AuthUser1.From, big.NewInt(test.ONE_ETH))
	test.IfNotNil(t, transErr2, fmt.Sprintf("Error transferring tokens to member: %v", transErr2))
	context.Blockchain.Commit()
	// member will need to have approved the voting contract to spend
	appErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser1,
		deployed.VotingAddress, big.NewInt(test.ONE_ETH))
	test.IfNotNil(t, appErr, fmt.Sprintf("Error approving market contract to spend: %v", appErr))
	context.Blockchain.Commit()

	// yay vote...
	_, voteErr2 := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, big.NewInt(1))
	test.IfNotNil(t, voteErr2, fmt.Sprintf("Error voting for candidate: %v", voteErr2))
	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	// call for resolution - now it becomes a listing
	_, resolveErr2 := deployed.ListingContract.ResolveApplication(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, resolveErr2, fmt.Sprintf("Error resolving application: %v", resolveErr2))
	context.Blockchain.Commit()

	// unstake again
	_, unErr2 := deployed.VotingContract.Unstake(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash)
	test.IfNotNil(t, unErr2, fmt.Sprintf("Error Unstaking: %v", unErr2))
	context.Blockchain.Commit()

	// with a listing in place we can claim that it was accessed
	query := test.GenBytes32("select * from SPAM where EGGS eq TRUE")
	// current bytes purchased, should decrease with listing access reporting
	bytesBal, _ := deployed.DatatrustContract.GetBytesPurchased(nil, context.AuthUser3.From)
	// current bytes_accessed for listiing, should increase with reporting
	accessBal, _ := deployed.DatatrustContract.GetBytesAccessed(nil, listingHash)
	// none delivered yet...
	_, _, delivered, _ := deployed.DatatrustContract.GetDelivery(nil, query)

	t.Log("Listing one accessed for half the request")
	_, accErr := deployed.DatatrustContract.ListingAccessed(test.GetTxOpts(context.AuthBackend, nil,
		// let's say one listing was used for 1/2 the request
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, query, big.NewInt((1024*1024)*5))
	test.IfNotNil(t, accErr, "Error claiming listing accessed")
	context.Blockchain.Commit()

	// as listings are accessed the user's bytes purchased bal decreases...
	bytesBalNow, _ := deployed.DatatrustContract.GetBytesPurchased(nil, context.AuthUser3.From)
	if bytesBalNow.Cmp(bytesBal) != -1 {
		t.Errorf("Expected %v to be > %v", bytesBal, bytesBalNow)
	}
	t.Logf("User3 bytes purchased balance post listing1 access: %v", test.Commafy(bytesBalNow))
	// current bytes_accessed for listing, should increase with reporting
	accessBalNow, _ := deployed.DatatrustContract.GetBytesAccessed(nil, listingHash)
	if accessBalNow.Cmp(accessBal) != 1 {
		t.Errorf("Expected %v to be > %v", accessBalNow, accessBal)
	}
	t.Logf("Listing1 bytes accessed: %v", test.Commafy(accessBalNow))
	// half has been delivered
	_, _, deliveredNow, _ := deployed.DatatrustContract.GetDelivery(nil, query)
	if deliveredNow.Cmp(delivered) != 1 {
		t.Errorf("Expected %v to be > %v", deliveredNow, delivered)
	}
	t.Logf("Delivery object bytes delivered: %v", test.Commafy(deliveredNow))
}

// if a backend calls for its payment after delivery - it should get paid
func TestDelivered(t *testing.T) {
	// we'll put up another listing to claim was accessed
	listingHash := test.GenBytes32("LookAtMyJunkToo")
	_, listErr := deployed.ListingContract.List(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 250000), listingHash)
	test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
	context.Blockchain.Commit()

	dataHash := test.GenBytes32("moardata")
	_, dataErr := deployed.DatatrustContract.SetDataHash(test.GetTxOpts(context.AuthBackend, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), listingHash, dataHash)
	test.IfNotNil(t, dataErr, "Error setting data hash for listing")
	context.Blockchain.Commit()

	// cast a vote for, voter may need funds...
	transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner,
		context.AuthUser2.From, big.NewInt(test.ONE_ETH))
	test.IfNotNil(t, transErr, fmt.Sprintf("Error transferring tokens to member: %v", transErr))
	context.Blockchain.Commit()
	// member will need to have approved the voting contract to spend
	appErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser2,
		deployed.VotingAddress, big.NewInt(test.ONE_ETH))
	test.IfNotNil(t, appErr, fmt.Sprintf("Error approving market contract to spend: %v", appErr))
	context.Blockchain.Commit()

	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser2, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))
	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	// call for resolution
	_, resolveErr := deployed.ListingContract.ResolveApplication(test.GetTxOpts(context.AuthUser2, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving application: %v", resolveErr))
	context.Blockchain.Commit()

	// unstake
	_, unErr := deployed.VotingContract.Unstake(test.GetTxOpts(context.AuthUser2, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash)
	test.IfNotNil(t, unErr, fmt.Sprintf("Error Unstaking: %v", unErr))
	context.Blockchain.Commit()

	// with that in place we can claim it as the source of the rest of the delivery
	query := test.GenBytes32("select * from SPAM where EGGS eq TRUE")

	t.Log("Listing2 reported as accessed for the 2nd half of the delivery")
	_, accErr := deployed.DatatrustContract.ListingAccessed(test.GetTxOpts(context.AuthBackend, nil,
		// let's say one listing was used for the other 1/2 of the request
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, query, big.NewInt((1024*1024)*5))
	test.IfNotNil(t, accErr, "Error claiming listing accessed")
	context.Blockchain.Commit()

	// we should see a delivery object whose bytes delivered match the bytes requested
	_, req, del, _ := deployed.DatatrustContract.GetDelivery(nil, query)
	if req.Cmp(del) != 0 {
		t.Errorf("Expected %v to be %v", req, del)
	}
	t.Logf("Delivery object bytes delivered: %v", test.Commafy(del))

	// we should see equal access for both listings
	ogListingHash := test.GenBytes32("LookAtMyJunk")
	ogAccessBal, _ := deployed.DatatrustContract.GetBytesAccessed(nil, ogListingHash)
	accessBal, _ := deployed.DatatrustContract.GetBytesAccessed(nil, listingHash)
	if accessBal.Cmp(ogAccessBal) != 0 {
		t.Errorf("Expected %v to be %v", accessBal, ogAccessBal)
	}
	t.Logf("Listing 2 bytes_accessed: %v", test.Commafy(accessBal))

	// the user should have no avail bytes at this time
	purchased, _ := deployed.DatatrustContract.GetBytesPurchased(nil, context.AuthUser3.From)
	if purchased.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("expected %v to be 0", purchased)
	}
	t.Logf("User1 remaining bytes_purchased: %v", purchased)

	// note the ether token balances
	beAddr, _ := deployed.DatatrustContract.GetBackendAddress(nil)
	beBal, _ := deployed.EtherTokenContract.BalanceOf(nil, beAddr)
	t.Logf("Backend Ether Token balance before delivering: %v", beBal)
	dtBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)

	// the backend must show that it put the delivery somewhere (this will be hashed)
	urlHash := test.GenBytes32("someURL.net/roger/shrubber")
	t.Log("Backend calls delivered now to get paid")
	_, delErr := deployed.DatatrustContract.Delivered(test.GetTxOpts(context.AuthBackend, nil,
		big.NewInt(test.ONE_GWEI*2), 250000), query, urlHash)
	test.IfNotNil(t, delErr, "Error calling delivered")
	context.Blockchain.Commit()
	// should see balance changes now
	beBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, beAddr)
	t.Logf("Backend Ether Token balance after delivering: %v", test.Commafy(beBalNow))
	dtBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)
	t.Logf("Datatrust banked balance after backend paid: %v", test.Commafy(beBalNow))

	// newest bal will have increased
	if beBalNow.Cmp(beBal) != 1 {
		t.Errorf("Expected %v to be > %v", beBalNow, beBal)
	}

	// newest bal will have decreased
	if dtBalNow.Cmp(dtBal) != -1 {
		t.Errorf("Expected %v to be < %v", dtBalNow, dtBal)
	}
}

// testing bytes access claiming here as it is, technically, part of the delivery flow.
func TestClaimBytesAccessed(t *testing.T) {
	// get the current accumulated byte access balance for the listings used
	listingHash1 := test.GenBytes32("LookAtMyJunk")
	listingHash2 := test.GenBytes32("LookAtMyJunkToo")
	// note the supply of those listings
	_, supply1, _ := deployed.ListingContract.GetListing(nil, listingHash1)
	t.Logf("Current supply of Listing1: %v", test.Commafy(supply1))
	_, supply2, _ := deployed.ListingContract.GetListing(nil, listingHash2)
	t.Logf("Current supply of Listing2: %v", test.Commafy(supply2))
	// note the current datatrust banked eth token amount, at this point it should only be the maker split(s) from the outstanding bytes accessed
	dataBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)
	// claim the listing for 1
	t.Log("Listing1 Maker claims bytes accessed")
	_, clErr := deployed.ListingContract.ClaimBytesAccessed(test.GetTxOpts(context.AuthUser2, nil,
		big.NewInt(test.ONE_GWEI*2), 250000), listingHash1)
	test.IfNotNil(t, clErr, "Error claiming access")
	context.Blockchain.Commit()

	// supply should have increased
	_, supply1Now, _ := deployed.ListingContract.GetListing(nil, listingHash1)
	if supply1Now.Cmp(supply1) != 1 {
		t.Errorf("Expected %v to be > %v", supply1Now, supply1)
	}
	t.Logf("Supply of Listing1 *IN MARKET TOKEN* after claiming: %v", test.Commafy(supply1Now))

	// access bal should be cleared
	accessBal1, _ := deployed.DatatrustContract.GetBytesAccessed(nil, listingHash1)
	if accessBal1.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected %v to be 0", accessBal1)
	}
	t.Logf("Listing1 bytes accessed balance after claiming: %v", accessBal1)

	// datatrust bank should be lower
	dataBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)
	if dataBalNow.Cmp(dataBal) != -1 {
		t.Errorf("Expected %v to be < %v", dataBalNow, dataBal)
	}
	t.Logf("Datatrust banked balance after maker 1 paid: %v", test.Commafy(dataBalNow))

	// NOTE if cost_per_byte is too low there is a scenario here where the 2nd listing's
	// maker_fee is too low to invest. Which is _not_ erroneous, but worth noting...

	// claim the other listing access
	t.Log("Listing2 Maker claims bytes accessed")
	_, clErr2 := deployed.ListingContract.ClaimBytesAccessed(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 250000), listingHash2)
	test.IfNotNil(t, clErr2, "Error claiming access")
	context.Blockchain.Commit()

	// supply should have increased
	_, supply2Now, _ := deployed.ListingContract.GetListing(nil, listingHash2)
	if supply2Now.Cmp(supply2) != 1 {
		t.Errorf("Expected %v to be > %v", supply2Now, supply2)
	}
	t.Logf("Supply of Listing2 *IN MARKET TOKEN* after claiming: %v", test.Commafy(supply2Now))

	// access bal should be cleared
	accessBal2, _ := deployed.DatatrustContract.GetBytesAccessed(nil, listingHash2)
	if accessBal2.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected %v to be 0", accessBal2)
	}
	t.Logf("Listing2 bytes accessed balance after claiming: %v", accessBal2)

	// datatrust bank should be empty now. NOTE this will not always be the case #dusting
	dataBalAgain, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)
	if dataBalAgain.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected %v to be 0", dataBalAgain)
	}
	t.Logf("Datatrust banked balance after maker 2 paid: %v", test.Commafy(dataBalAgain))
}
