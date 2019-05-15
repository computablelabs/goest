package scenariofive

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
	"time"
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
	buyer := context.AuthUser3
	// user has no credits atm
	purchased, _ := deployed.DatatrustContract.GetBytesPurchased(nil, buyer.From)
	// reserve balance atm. This will go up by the res_payment after a request...
	resBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.InvestingAddress)

	// make a deposit in ETH, resulting in a 1:1 ethertoken balance
	_, depErr := deployed.EtherTokenContract.Deposit(test.GetTxOpts(buyer,
		big.NewInt(test.ONE_WEI), big.NewInt(test.ONE_GWEI*2), 100000))
	test.IfNotNil(t, depErr, fmt.Sprintf("Error depositing funds from member to ether token: %v", depErr))

	context.Blockchain.Commit()

	// user now has an ether token balance
	ethBal, _ := deployed.EtherTokenContract.BalanceOf(nil, buyer.From)
	if ethBal.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected member bal to be > 0, got: %v", ethBal)
	}

	// purchase payments are banked by the datatrust contract at the ether token (-res_payment)
	dataEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)

	// member needs to have approved datatrust to spend ether token
	_, approveErr := deployed.EtherTokenContract.Approve(test.GetTxOpts(buyer, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), deployed.DatatrustAddress, big.NewInt(test.ONE_WEI))
	test.IfNotNil(t, approveErr, fmt.Sprintf("Error approving spender: %v", approveErr))

	context.Blockchain.Commit()

	// note the allowance as later purchasing should decrease it
	ethAllow, _ := deployed.EtherTokenContract.Allowance(nil, buyer.From, deployed.DatatrustAddress)

	// when a delivery is requested, a hash of the query datatrust will recieve is expected
	query := test.GenBytes32("select * from SPAM where EGGS eq TRUE")

	// assure no delivery exists for this query
	owner, req, _, _ := deployed.DatatrustContract.GetDelivery(nil, query)
	if owner == buyer.From {
		t.Error("Expected no delivery object yet")
	}

	if req.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected bytes requested to be 0, got: %v", req)
	}

	// the actual request for a delivery...
	_, delErr := deployed.DatatrustContract.RequestDelivery(test.GetTxOpts(
		buyer, nil, big.NewInt(test.ONE_GWEI*2), 250000), query, big.NewInt(1024*1024)) // ~1MB
	test.IfNotNil(t, delErr, fmt.Sprintf("Error requesting delivery: %v", delErr))

	context.Blockchain.Commit()

	// user should now have some bytes
	purchasedNow, _ := deployed.DatatrustContract.GetBytesPurchased(nil, buyer.From)
	if purchasedNow.Cmp(purchased) != 1 {
		t.Errorf("Expected byte balance %v to be more than %v", purchasedNow, purchased)
	}

	// ether token balances should be updated
	ethBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, buyer.From)
	if ethBalNow.Cmp(ethBal) != -1 {
		t.Errorf("Expected %v to be < %v", ethBalNow, ethBal)
	}

	// allowances also reflect spending...
	ethAllowNow, _ := deployed.EtherTokenContract.Allowance(nil, buyer.From, deployed.DatatrustAddress)
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
	if ownerNow != buyer.From {
		t.Error("Expected delivery object to be owned by user 3")
	}
	if reqNow.Cmp(big.NewInt(1024*1024)) != 0 {
		//if reqNow.Cmp(big.NewInt(1024)) != 0 {
		t.Errorf("Expected bytes requested to be 1MB, got: %v", reqNow)
	}
}

func TestListingAccessed(t *testing.T) {
	maker := context.AuthUser2
	buyer := context.AuthUser3
	// get a listing up
	listingHash := test.GenBytes32("LookAtMyJunk")
	_, listErr := deployed.ListingContract.List(test.GetTxOpts(maker, nil,
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
		context.AuthUser1.From, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr2, fmt.Sprintf("Error transferring tokens to member: %v", transErr2))
	context.Blockchain.Commit()
	// member will need to have approved the voting contract to spend
	appErr := test.MaybeIncreaseMarketTokenApproval(context, deployed, context.AuthUser1,
		deployed.VotingAddress, big.NewInt(test.ONE_GWEI))
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

	// any council member can call for resolution - now it becomes a listing
	_, resolveErr2 := deployed.ListingContract.ResolveApplication(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, resolveErr2, fmt.Sprintf("Error resolving application: %v", resolveErr2))
	context.Blockchain.Commit()

	// with a listing in place we can claim that it was accessed
	query := test.GenBytes32("select * from SPAM where EGGS eq TRUE")
	// current bytes purchased, should decrease with listing access reporting
	//bytesBal, _ := deployed.DatatrustContract.GetBytesPurchased(nil, buyer.From)
	_, _ = deployed.DatatrustContract.GetBytesPurchased(nil, buyer.From)
	// current bytes_accessed for listiing, should increase with reporting
	//accessBal, _ := deployed.DatatrustContract.GetBytesAccessed(nil, listingHash)
	_, _ = deployed.DatatrustContract.GetBytesAccessed(nil, listingHash)
	// none delivered yet...
	//_, _, delivered, _ := deployed.DatatrustContract.GetDelivery(nil, query)
	_, _, _, _ = deployed.DatatrustContract.GetDelivery(nil, query)

	_, accErr := deployed.DatatrustContract.ListingAccessed(test.GetTxOpts(context.AuthBackend, nil,
		// let's say one listing was used for 1/2 the request
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, query, big.NewInt(1024*512))
	test.IfNotNil(t, accErr, "Error claiming listing accessed")
	context.Blockchain.Commit()

	//// as listings are accessed the user's bytes purchased bal decreases...
	//bytesBalNow, _ := deployed.DatatrustContract.GetBytesPurchased(nil, buyer.From)
	//if bytesBalNow.Cmp(bytesBal) != -1 {
	//	t.Errorf("Expected %v to be > %v", bytesBal, bytesBalNow)
	//}
	//// current bytes_accessed for listing, should increase with reporting
	//accessBalNow, _ := deployed.DatatrustContract.GetBytesAccessed(nil, listingHash)
	//if accessBalNow.Cmp(accessBal) != 1 {
	//	t.Errorf("Expected %v to be > %v", accessBalNow, accessBal)
	//}
	//// half has been delivered
	//_, _, deliveredNow, _ := deployed.DatatrustContract.GetDelivery(nil, query)
	//if deliveredNow.Cmp(delivered) != 1 {
	//	t.Errorf("Expected %v to be > %v", deliveredNow, delivered)
	//}
}

//
//// if a backend calls for its payment after delivery - it should get paid
//func TestDelivered(t *testing.T) {
//	// we'll put up another listing to claim was accessed
//	listingHash := test.GenBytes32("LookAtMyJunkToo")
//	_, listErr := deployed.ListingContract.List(test.GetTxOpts(context.AuthUser1, nil,
//		big.NewInt(test.ONE_GWEI*2), 250000), listingHash)
//	test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
//	context.Blockchain.Commit()
//
//	dataHash := test.GenBytes32("moardata")
//	_, dataErr := deployed.DatatrustContract.SetDataHash(test.GetTxOpts(context.AuthBackend, nil,
//		big.NewInt(test.ONE_GWEI*2), 100000), listingHash, dataHash)
//	test.IfNotNil(t, dataErr, "Error setting data hash for listing")
//	context.Blockchain.Commit()
//
//	// cast a vote for, voter may need funds...
//	transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner,
//		context.AuthUser2.From, big.NewInt(test.ONE_GWEI))
//	test.IfNotNil(t, transErr, fmt.Sprintf("Error transferring tokens to member: %v", transErr))
//	context.Blockchain.Commit()
//	// member will need to have approved the voting contract to spend
//	appErr := test.MaybeIncreaseMarketTokenApproval(context, deployed, context.AuthUser2,
//		deployed.VotingAddress, big.NewInt(test.ONE_GWEI))
//	test.IfNotNil(t, appErr, fmt.Sprintf("Error approving market contract to spend: %v", appErr))
//	context.Blockchain.Commit()
//
//	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser2, nil,
//		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, big.NewInt(1))
//	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))
//	context.Blockchain.Commit()
//
//	// move past the voteBy
//	context.Blockchain.AdjustTime(100 * time.Second)
//	context.Blockchain.Commit()
//
//	// any council member can call for resolution
//	_, resolveErr := deployed.ListingContract.ResolveApplication(test.GetTxOpts(context.AuthUser2, nil,
//		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
//	test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving application: %v", resolveErr))
//	context.Blockchain.Commit()
//
//	// with that in place we can claim it as the source of the rest of the delivery
//	query := test.GenBytes32("select * from SPAM where EGGS eq TRUE")
//	_, accErr := deployed.DatatrustContract.ListingAccessed(test.GetTxOpts(context.AuthBackend, nil,
//		// let's say one listing was used for the other 1/2 of the request
//		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, query, big.NewInt(1024*512))
//	test.IfNotNil(t, accErr, "Error claiming listing accessed")
//	context.Blockchain.Commit()
//
//	// we should see a delivery object whose bytes delivered match the bytes requested
//	_, req, del, _ := deployed.DatatrustContract.GetDelivery(nil, query)
//	if req.Cmp(del) != 0 {
//		t.Errorf("Expected %v to be %v", req, del)
//	}
//
//	// we should see equal access for both listings
//	ogListingHash := test.GenBytes32("LookAtMyJunk")
//	ogAccessBal, _ := deployed.DatatrustContract.GetBytesAccessed(nil, ogListingHash)
//	accessBal, _ := deployed.DatatrustContract.GetBytesAccessed(nil, listingHash)
//	if accessBal.Cmp(ogAccessBal) != 0 {
//		t.Errorf("Expected %v to be %v", accessBal, ogAccessBal)
//	}
//
//	// the user should have no avail bytes at this time
//	purchased, _ := deployed.DatatrustContract.GetBytesPurchased(nil, context.AuthUser3.From)
//	if purchased.Cmp(big.NewInt(0)) != 0 {
//		t.Errorf("expected %v to be 0", purchased)
//	}
//
//	// note the ether token balances
//	beAddr, _ := deployed.DatatrustContract.GetBackendAddress(nil)
//	beBal, _ := deployed.EtherTokenContract.BalanceOf(nil, beAddr)
//	dtBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)
//
//	// the backend must show that it put the delivery somewhere (this will be hashed)
//	urlHash := test.GenBytes32("someURL.net/roger/shrubber")
//	_, delErr := deployed.DatatrustContract.Delivered(test.GetTxOpts(context.AuthBackend, nil,
//		big.NewInt(test.ONE_GWEI*2), 250000), query, urlHash)
//	test.IfNotNil(t, delErr, "Error calling delivered")
//	context.Blockchain.Commit()
//	// should see balance changes now
//	beBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, beAddr)
//	dtBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)
//
//	// newest bal will have increased
//	if beBalNow.Cmp(beBal) != 1 {
//		t.Errorf("Expected %v to be > %v", beBalNow, beBal)
//	}
//
//	// newest bal will have decreased
//	if dtBalNow.Cmp(dtBal) != -1 {
//		t.Errorf("Expected %v to be < %v", dtBalNow, dtBal)
//	}
//}
//
//// testing bytes access claiming here as it is, technically, part of the delivery flow.
//func TestClaimBytesAccessed(t *testing.T) {
//	// get the current accumulated byte access balance for the listings used
//	listingHash1 := test.GenBytes32("LookAtMyJunk")
//	listingHash2 := test.GenBytes32("LookAtMyJunkToo")
//	// note the supply of those listings
//	_, supply1, _ := deployed.ListingContract.GetListing(nil, listingHash1)
//	_, supply2, _ := deployed.ListingContract.GetListing(nil, listingHash2)
//	// note the current datatrust banked eth token amount, at this point it should only be the maker split(s) from the outstanding bytes accessed
//	dataBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)
//	// claim the listing for 1
//	_, clErr := deployed.ListingContract.ClaimBytesAccessed(test.GetTxOpts(context.AuthUser2, nil,
//		big.NewInt(test.ONE_GWEI*2), 250000), listingHash1)
//	test.IfNotNil(t, clErr, "Error claiming access")
//	context.Blockchain.Commit()
//
//	// supply should have increased
//	_, supply1Now, _ := deployed.ListingContract.GetListing(nil, listingHash1)
//	if supply1Now.Cmp(supply1) != 1 {
//		t.Errorf("Expected %v to be > %v", supply1Now, supply1)
//	}
//
//	// access bal should be cleared
//	accessBal1, _ := deployed.DatatrustContract.GetBytesAccessed(nil, listingHash1)
//	if accessBal1.Cmp(big.NewInt(0)) != 0 {
//		t.Errorf("Expected %v to be 0", accessBal1)
//	}
//
//	// datatrust bank should be lower
//	dataBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)
//	if dataBalNow.Cmp(dataBal) != -1 {
//		t.Errorf("Expected %v to be < %v", dataBalNow, dataBal)
//	}
//
//	// NOTE if cost_per_byte is too low there is a scenario here where the 2nd listing's
//	// maker_fee is too low to invest. Which is _not_ erroneous, but worth noting...
//
//	// claim the other listing access
//	_, clErr2 := deployed.ListingContract.ClaimBytesAccessed(test.GetTxOpts(context.AuthUser1, nil,
//		big.NewInt(test.ONE_GWEI*2), 250000), listingHash2)
//	test.IfNotNil(t, clErr2, "Error claiming access")
//	context.Blockchain.Commit()
//
//	// supply should have increased
//	_, supply2Now, _ := deployed.ListingContract.GetListing(nil, listingHash2)
//	if supply2Now.Cmp(supply2) != 1 {
//		t.Errorf("Expected %v to be > %v", supply2Now, supply2)
//	}
//
//	// access bal should be cleared
//	accessBal2, _ := deployed.DatatrustContract.GetBytesAccessed(nil, listingHash2)
//	if accessBal2.Cmp(big.NewInt(0)) != 0 {
//		t.Errorf("Expected %v to be 0", accessBal2)
//	}
//
//	// datatrust bank should be empty now. NOTE this will not always be the case #dusting TODO
//	dataBalAgain, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)
//	if dataBalAgain.Cmp(big.NewInt(0)) != 0 {
//		t.Errorf("Expected %v to be 0", dataBalAgain)
//	}
//}
