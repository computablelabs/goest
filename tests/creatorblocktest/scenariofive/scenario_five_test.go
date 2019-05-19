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

func TestFullSimulation(t *testing.T) {
	// The creator of the market starts by transfering 100 ETH to reserve
	_, transError := deployed.EtherTokenContract.Transfer(test.GetTxOpts(
		context.AuthOwner, nil, big.NewInt(test.ONE_GWEI*2), 100000),
		deployed.InvestingAddress, oneHundredEth())
	test.IfNotNil(t, transError, "Error transferring token")
	context.Blockchain.Commit()

	ownerEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthOwner.From)
	resEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.InvestingAddress)

	// has 1 ETH left
	if ownerEthBal.Cmp(big.NewInt(test.ONE_WEI)) != 0 {
		t.Errorf("Expected ether token balance of 1 ETH, got: %v", ownerEthBal)
	}

	// got the 100
	if resEthBal.Cmp(oneHundredEth()) != 0 {
		t.Errorf("Expected reserve of 100 Eth, got: %v", resEthBal)
	}

	t.Logf("Current Reserve balance: %v", test.Commafy(resEthBal))
	// Original market token supply
	mtSup, _ := deployed.MarketTokenContract.TotalSupply(nil)
	t.Logf("Original Market token total supply post creation: %v", test.Commafy(mtSup))

	t.Logf("Number Makers: %d", len(makers))
	t.Logf("Number Investors: %d", len(investors))

	// Let's have the investors invest
	for name, investor := range investors {
		t.Logf("%s is investing", name)

		// investor deposits eth in the ethToken
		etBal, _ := deployed.EtherTokenContract.BalanceOf(nil, investor.From)
		if etBal.Cmp(big.NewInt(0)) != 0 {
			t.Errorf("Expected ether token balance of 0, got: %v", etBal)
		}
		_, depErr := deployed.EtherTokenContract.Deposit(test.GetTxOpts(investor,
			oneHundredEth(), big.NewInt(test.ONE_GWEI*2), 100000))
		test.IfNotNil(t, depErr, "Error depositing ETH")
		context.Blockchain.Commit()

		// Check that the investor's deposit into ethToken worked
		etBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, investor.From)
		if etBalNow.Cmp(oneHundredEth()) != 0 {
			t.Errorf("Expected ether token balance of 100 eth, got: %v", etBalNow)
		}

		// Check that investor currently has no market token
		mtBal, _ := deployed.MarketTokenContract.BalanceOf(nil, investor.From)
		if mtBal.Cmp(big.NewInt(0)) != 0 {
			t.Errorf("Expected market token balance of 0, got: %v", mtBal)
		}
		t.Logf("%s current market token balance: %v", name, mtBal)

		// Get current invest price
		invPrice, _ := deployed.InvestingContract.GetInvestmentPrice(nil)
		t.Logf("Current invest price: %v", test.Commafy(invPrice))

		// investor wants to invest 100 ETH. Must approve inv contract first...
		_, approveErr := deployed.EtherTokenContract.Approve(test.GetTxOpts(investor, nil,
			big.NewInt(test.ONE_GWEI*2), 100000), deployed.InvestingAddress, oneHundredEth()) // up to 100 ETH
		test.IfNotNil(t, approveErr, fmt.Sprintf("Error approving market contract to spend: %v", approveErr))
		context.Blockchain.Commit()

		// investing has that allowance now
		allowed, _ := deployed.EtherTokenContract.Allowance(nil, investor.From, deployed.InvestingAddress)
		if allowed.Cmp(oneHundredEth()) != 0 {
			t.Errorf("Expected allowance of 100 ETH, got: %v", allowed)
		}

		// the actual investment (now that we know we can)
		_, invErr := deployed.InvestingContract.Invest(test.GetTxOpts(investor, nil,
			big.NewInt(test.ONE_GWEI*2), 150000), oneHundredEth())
		test.IfNotNil(t, invErr, "Error investing")
		context.Blockchain.Commit()

		// check current market token balance
		mtBalNow, _ := deployed.MarketTokenContract.BalanceOf(nil, investor.From)
		if mtBalNow.Cmp(mtBal) != 1 {
			t.Errorf("Expected %v to be > %v", mtBalNow, mtBal)
		}
		t.Logf("%s market token balance post 100 ETH investment: %v", name, test.Commafy(mtBalNow))

		// market token total supply should be updated
		mtSup, _ := deployed.MarketTokenContract.TotalSupply(nil)
		t.Logf("Market token total supply post %s investment of 100 ETH: %v", name, test.Commafy(mtSup))

		// Get new reserve balance
		resEthBal, _ = deployed.EtherTokenContract.BalanceOf(nil, deployed.InvestingAddress)
		// reserve should be updated
		t.Logf("Reserve balance post %s investment of 100 ETH: %v", name, test.Commafy(resEthBal))

		// invest price should change
		invPriceNow, _ := deployed.InvestingContract.GetInvestmentPrice(nil)
		t.Logf("Investment Price post %s investment of 100 ETH: %v", name, test.Commafy(invPriceNow))

	}

	// Let's have the maker submit a listing
	for name, maker := range makers {
		t.Logf("Submitting listing for %s", name)
		listingHash := test.GenBytes32(name)
		listings[name] = listingHash

		_, listErr := deployed.ListingContract.List(test.GetTxOpts(maker, nil,
			big.NewInt(test.ONE_GWEI*2), 250000), listingHash)
		test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
		context.Blockchain.Commit()

		// Check that the submitted listing is now a candidate
		isCandidate, _ := deployed.VotingContract.IsCandidate(nil, listingHash)

		if isCandidate == false {
			t.Error("Submitted listing is not listed as a candidate")
		}

		// have the datatrust submit a data_hash for this listing.
		dataHash, _ := deployed.DatatrustContract.GetHash(nil, name)
		_, dataErr := deployed.DatatrustContract.SetDataHash(test.GetTxOpts(context.AuthBackend, nil,
			big.NewInt(test.ONE_GWEI*2), 100000), listingHash, dataHash)
		test.IfNotNil(t, dataErr, "Error setting data hash for listing")
		context.Blockchain.Commit()

		// check the market's balance as the mint operation should increment it after successful listing
		marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)

		// cast a vote for, voter may need funds...
		// TODO(rbharath): Can this line perhaps be removed? Since we start accounts off with generous balances?
		transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser2.From,
			big.NewInt(10000000000000000)) // 1 X 10**16
		test.IfNotNil(t, transErr, fmt.Sprintf("Error transferring tokens to member: %v", transErr))
		context.Blockchain.Commit()

		// member will need to have approved the voting contract to spend
		appErr := test.MaybeIncreaseMarketTokenApproval(context, deployed, context.AuthUser2, deployed.VotingAddress,
			big.NewInt(10000000000000000)) // 1 X 10**16
		test.IfNotNil(t, appErr, fmt.Sprintf("Error approving market contract to spend: %v", appErr))
		context.Blockchain.Commit()

		// TODO: Should we increase the amount the contract votes with?
		_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser2, nil,
			big.NewInt(test.ONE_GWEI*2), 150000), listingHash, big.NewInt(1))
		test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))
		context.Blockchain.Commit()

		// move past the voteBy
		context.Blockchain.AdjustTime(100 * time.Second)
		context.Blockchain.Commit()

		// any stakeholder can call for resolution
		_, resolveErr := deployed.ListingContract.ResolveApplication(test.GetTxOpts(context.AuthUser2, nil,
			big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
		test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving application: %v", resolveErr))
		context.Blockchain.Commit()

		// Check that the submitted listing is no longer a candidate
		isCandidate, _ = deployed.VotingContract.IsCandidate(nil, listingHash)
		if isCandidate == true {
			t.Error("Expected approved listing's candidate to have been removed")
		}

		// should be listed now, with a reward
		owner, supply, _ := deployed.ListingContract.GetListing(nil, listingHash)

		if owner != maker.From {
			t.Errorf("Expected owner to be %v, got: %v", maker.From, owner)
		}
		t.Logf("Listing supply: %v", test.Commafy(supply))

		// makerMarketBal should be 0 since the balance has not been withdrawn yet
		makerMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, maker.From)
		t.Logf("%s MarketToken balance: %v", name, test.Commafy(makerMarketBal))

		// supply should reflect the list reward
		if supply.Cmp(big.NewInt(250000000000000)) != 0 {
			t.Errorf("Expected supply to be list reward, got: %v", supply)
		}

		// marketToken should be banking the minted amount
		newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)

		// we should see new bal being > old one
		if newMarketBal.Cmp(marketBal) != 1 {
			t.Errorf("Expected %v to be > %v", newMarketBal, marketBal)
		}

		// Market Token supply post listing
		mtSup, _ := deployed.MarketTokenContract.TotalSupply(nil)
		t.Logf("Original Market token total supply post %s listing: %v", name, test.Commafy(mtSup))

	}

	// Let's have the buyers buy
	for name, buyer := range buyers {
		t.Logf("%s is buying", name)

		// user has no credits atm
		purchased, _ := deployed.DatatrustContract.GetBytesPurchased(nil, buyer.From)

		if purchased.Cmp(big.NewInt(0)) != 0 {
			t.Errorf("Expected number purchased bytes to be 0, got: %v", purchased)
		}

		// reserve balance atm. This will go up by the res_payment after a request...
		resBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.InvestingAddress)
		t.Logf("Reserve balance before purchase: %v", test.Commafy(resBal))

		// make a deposit in ETH, resulting in a 1:1 ethertoken balance
		_, depErr := deployed.EtherTokenContract.Deposit(test.GetTxOpts(buyer,
			oneHundredEth(), big.NewInt(test.ONE_GWEI*2), 100000))
		test.IfNotNil(t, depErr, fmt.Sprintf("Error depositing funds from member to ether token: %v", depErr))

		context.Blockchain.Commit()

		// buyer now has an ether token balance
		ethBal, _ := deployed.EtherTokenContract.BalanceOf(nil, buyer.From)
		if ethBal.Cmp(big.NewInt(0)) != 1 {
			t.Errorf("Expected member bal to be > 0, got: %v", ethBal)
		}

		// purchase payments are banked by the datatrust contract at the ether token (-res_payment)
		dataEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.DatatrustAddress)

		// buyer needs to have approved datatrust to spend ether token
		_, approveErr := deployed.EtherTokenContract.Approve(test.GetTxOpts(buyer, nil,
			big.NewInt(test.ONE_GWEI*2), 100000), deployed.DatatrustAddress, oneHundredEth())
		test.IfNotNil(t, approveErr, fmt.Sprintf("Error approving spender: %v", approveErr))

		context.Blockchain.Commit()

		// note the allowance as later purchasing should decrease it
		ethAllow, _ := deployed.EtherTokenContract.Allowance(nil, buyer.From, deployed.DatatrustAddress)

		// when a delivery is requested, a hash of the query datatrust will recieve is expected
		// TODO(rbharath): This test highlights a bit of funkiness that two people can't make exactly the same query. Might need to prepend queries with names as done here.
		query := test.GenBytes32(fmt.Sprintf("%s: select * from SPAM where EGGS eq TRUE", name))

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

		// allowances should be lower now to reflect spending...
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
		t.Logf("Current Reserve balance: %v", test.Commafy(resBalNow))

		// at this point the sum of what went into the reserve + the amount locked in datatrust will == bytesPurchased * cost_per_byte
		toRes := resBalNow.Sub(resBalNow, resBal)
		toData := dataEthBal.Sub(dataEthBalNow, dataEthBal)
		summed := toRes.Add(toRes, toData)
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
			t.Errorf("Expected bytes requested to be 1MB, got: %v", reqNow)
		}

		// Note that the number of listings has to divide 1024
		numListings := len(listings)
		creditPerListing := big.NewInt(int64(1024 * 1024 / numListings))
		// Let's loop over the listings to assign credits
		for _, listingHash := range listings {
			_, accErr := deployed.DatatrustContract.ListingAccessed(test.GetTxOpts(context.AuthBackend, nil,
				// let's say one listing was used for 1/2 the request
				big.NewInt(test.ONE_GWEI*2), 150000), listingHash, query, creditPerListing)
			test.IfNotNil(t, accErr, "Error claiming listing accessed")
			context.Blockchain.Commit()
		}

	}
}
