package scenariothree

import (
	//"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
	//"time"
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

func TestInitialBalance(t *testing.T) {
	// owner has all the eth atm
	etBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthOwner.From)
	// all market tokens that exist atm
	mtSup, _ := deployed.MarketTokenContract.TotalSupply(nil)
	if etBal.Cmp(oneHundredOneEth()) != 0 {
		t.Errorf("Expected ether token balance of %v, got: %v", oneHundredOneEth(), etBal)
	}
	if mtSup.Cmp(big.NewInt(test.ONE_WEI)) != 0 {
		t.Errorf("Expected market token supply of 1wei, got: %v", mtSup)
	}

	t.Logf("Current Market Token total supply: %v", test.Commafy(mtSup))
}

// This differs from test in scenario-two since it uses an extended context.
func TestTransferToReserveThenInvest(t *testing.T) {
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
	t.Logf("Number Investors: %d", len(investors))

	for _, investor := range investors {

		// investor deposits eth in the ethToken
		etBal, _ := deployed.EtherTokenContract.BalanceOf(nil, investor.From)
		if etBal.Cmp(big.NewInt(0)) != 0 {
			t.Errorf("Expected ether token balance of 0, got: %v", etBal)
		}
		_, depErr := deployed.EtherTokenContract.Deposit(test.GetTxOpts(investor,
			oneHundredEth(), big.NewInt(test.ONE_GWEI*2), 100000))
		test.IfNotNil(t, depErr, "Error depositing ETH")
		context.Blockchain.Commit()

		//	// snapshot current balances
		//	etBalNow, _ := deployed.EtherTokenContract.BalanceOf(nil, investor.From)
		//	if etBalNow.Cmp(oneHundredEth()) != 0 {
		//		t.Errorf("Expected ether token balance of 100 eth, got: %v", etBalNow)
		//	}
		//	mtBal, _ := deployed.MarketTokenContract.BalanceOf(nil, investor.From)
		//	if mtBal.Cmp(big.NewInt(0)) != 0 {
		//		t.Errorf("Expected market token balance of 0, got: %v", mtBal)
		//	}
		//	t.Logf("Investor %d current market token balance: %v", (ind + 1), mtBal)

		//	// invest price currently
		//	invPrice, _ := deployed.InvestingContract.GetInvestmentPrice(nil)
		//	t.Logf("Current invest price: %v", test.Commafy(invPrice))

		//	// investor wants to invest 100 ETH. Must approve inv contract first...
		//	_, approveErr := deployed.EtherTokenContract.Approve(test.GetTxOpts(investor, nil,
		//		big.NewInt(test.ONE_GWEI*2), 100000), deployed.InvestingAddress, oneHundredEth()) // up to 100 ETH
		//	test.IfNotNil(t, approveErr, fmt.Sprintf("Error approving market contract to spend: %v", approveErr))
		//	context.Blockchain.Commit()

		//	// investing has that allowance now
		//	allowed, _ := deployed.EtherTokenContract.Allowance(nil, investor.From, deployed.InvestingAddress)
		//	if allowed.Cmp(oneHundredEth()) != 0 {
		//		t.Errorf("Expected allowance of 100 ETH, got: %v", allowed)
		//	}

		//	// the actual investment (now that we know we can)
		//	_, invErr := deployed.InvestingContract.Invest(test.GetTxOpts(investor, nil,
		//		big.NewInt(test.ONE_GWEI*2), 150000), oneHundredEth())
		//	test.IfNotNil(t, invErr, "Error investing")
		//	context.Blockchain.Commit()

		//	// check current market token balance
		//	mtBalNow, _ := deployed.MarketTokenContract.BalanceOf(nil, investor.From)
		//	if mtBalNow.Cmp(mtBal) != 1 {
		//		t.Errorf("Expected %v to be > %v", mtBalNow, mtBal)
		//	}
		//	t.Logf("Investor %d market token balance post 100 ETH investment: %v", (ind + 1), test.Commafy(mtBalNow))

		//	// market token total supply should be updated
		//	mtSup, _ := deployed.MarketTokenContract.TotalSupply(nil)
		//	t.Logf("Market token total supply post investor %d investment of 100 ETH: %v", (ind + 1), test.Commafy(mtSup))

		//	// Get new reserve balance
		//	resEthBal, _ = deployed.EtherTokenContract.BalanceOf(nil, deployed.InvestingAddress)
		//	// reserve should be updated
		//	t.Logf("Reserve balance post investor %d investment of 100 ETH: %v", (ind + 1), test.Commafy(resEthBal))

		//	// invest price should change
		//	invPriceNow, _ := deployed.InvestingContract.GetInvestmentPrice(nil)
		//	t.Logf("Investment Price post investor %d investment of 100 ETH: %v", (ind + 1), test.Commafy(invPriceNow))
		//}

		//// Let's have the maker submit a listing
		//maker := context.AuthUser6
		//listingHash := test.GenBytes32("FooMarket, AZ.")

		//_, listErr := deployed.ListingContract.List(test.GetTxOpts(maker, nil,
		//	big.NewInt(test.ONE_GWEI*2), 250000), listingHash)
		//test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
		//context.Blockchain.Commit()

		//// Check that the submitted listing is now a candidate
		//isCan, _ := deployed.VotingContract.IsCandidate(nil, listingHash)

		//if !isCan {
		//	t.Errorf("Expected isCandidate to be true, got: %v", isCan)
		//}

		//// have the datatrust submit a data_hash for this listing.
		//dataHash, _ := deployed.DatatrustContract.GetHash(nil, "thisissomedata")
		//_, dataErr := deployed.DatatrustContract.SetDataHash(test.GetTxOpts(context.AuthBackend, nil,
		//	big.NewInt(test.ONE_GWEI*2), 100000), listingHash, dataHash)
		//test.IfNotNil(t, dataErr, "Error setting data hash for listing")
		//context.Blockchain.Commit()

		////// check the market's balance as the mint operation should increment it after successful listing
		////marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)

		//// cast a vote for, voter may need funds...
		//transErr := test.MaybeTransferMarketToken(context.Blockchain, deployed, context.AuthOwner, context.AuthUser2.From,
		//	big.NewInt(test.ONE_GWEI))
		//test.IfNotNil(t, transErr, fmt.Sprintf("Error transferring tokens to member: %v", transErr))
		//context.Blockchain.Commit()

		//// member will need to have approved the voting contract to spend
		//appErr := test.MaybeIncreaseMarketTokenApproval(context.Blockchain, deployed, context.AuthUser2, deployed.VotingAddress,
		//	big.NewInt(test.ONE_GWEI))
		//test.IfNotNil(t, appErr, fmt.Sprintf("Error approving market contract to spend: %v", appErr))
		//context.Blockchain.Commit()

		//_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser2, nil,
		//	big.NewInt(test.ONE_GWEI*2), 150000), listingHash, big.NewInt(1))
		//test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))
		//context.Blockchain.Commit()

		//// move past the voteBy
		//context.Blockchain.AdjustTime(100 * time.Second)
		//context.Blockchain.Commit()

		//// any stakeholder can call for resolution
		//_, resolveErr := deployed.ListingContract.ResolveApplication(test.GetTxOpts(context.AuthUser2, nil,
		//	big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
		//test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving application: %v", resolveErr))
		//context.Blockchain.Commit()

		//// should be listed now, with a reward
		//owner, _, _ := deployed.ListingContract.GetListing(nil, listingHash)

		//if owner != maker.From {
		//	t.Errorf("Expected owner to be %v, got: %v", maker.From, owner)
		//}

		//// supply should reflect the list reward
		//if supply.Cmp(big.NewInt(250000000000000)) != 0 {
		//	t.Errorf("Exepected supply to be list reward, got: %v", supply)
	}
}
