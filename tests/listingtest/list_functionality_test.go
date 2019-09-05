package listingtest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
	"time"
)

func TestList(t *testing.T) {
	listingHash := test.GenBytes32("FooMarket, AZ.")

	_, listErr := deployed.ListingContract.List(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 250000), listingHash)
	test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
	context.Blockchain.Commit()

	isCan, _ := deployed.VotingContract.IsCandidate(nil, listingHash)

	if !isCan {
		t.Errorf("Expected isCandidate to be true, got: %v", isCan)
	}
}

func TestIsListed(t *testing.T) {
	// we'll need to gen the listingHash to match
	listingHash := test.GenBytes32("FooMarket, AZ.")

	isListed, _ := deployed.ListingContract.IsListed(nil, listingHash)
	if isListed == true {
		t.Errorf("Expected isListed to be false, got: %v", isListed)
	}
}

func TestGetListing(t *testing.T) {
	listingHash := test.GenBytes32("FooMarket, AZ.")

	owner, supply, _ := deployed.ListingContract.GetListing(nil, listingHash)

	// should not be owned yet
	if owner == context.AuthUser1.From {
		t.Errorf("Exepected owner to be %v, got: %v", context.AuthUser1.From, owner)
	}
	// no list reward yet
	if supply.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Exepected supply to be 0, got: %v", supply)
	}
}

func TestResolveApplication(t *testing.T) {
	// our listing
	listingHash := test.GenBytes32("FooMarket, AZ.")
	// the datatrust must have a data_hash for this listing before it will pass
	dataHash, _ := deployed.DatatrustContract.GetHash(nil, "thisissomedata")
	_, dataErr := deployed.DatatrustContract.SetDataHash(test.GetTxOpts(context.AuthBackend, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), listingHash, dataHash)
	test.IfNotNil(t, dataErr, "Error setting data hash for listing")
	context.Blockchain.Commit()
	// check the market's balance as the mint operation should increment it after successful listing
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)

	// cast a vote for, voter may need funds...
	transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser2.From,
		big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr, fmt.Sprintf("Error transferring tokens to member: %v", transErr))
	context.Blockchain.Commit()
	// member will need to have approved the voting contract to spend
	appErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser2, deployed.VotingAddress,
		big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, appErr, fmt.Sprintf("Error approving market contract to spend: %v", appErr))
	context.Blockchain.Commit()

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

	// should be listed now, with a reward
	owner, supply, _ := deployed.ListingContract.GetListing(nil, listingHash)

	if owner != context.AuthUser1.From {
		t.Errorf("Expected owner to be %v, got: %v", context.AuthUser1.From, owner)
	}

	// supply should reflect the list reward
	if supply.Cmp(big.NewInt(test.ONE_ETH)) != 0 {
		t.Errorf("Exepected supply to be 0, got: %v", supply)
	}

	// marketToken should be banking the minted amount
	newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)
	// we should see new bal being > old one
	if newMarketBal.Cmp(marketBal) != 1 {
		t.Errorf("Expected %v to be > %v", newMarketBal, marketBal)
	}

	// post resolution any candidate should have been removed
	isCandidate, _ := deployed.VotingContract.IsCandidate(nil, listingHash)

	if isCandidate == true {
		t.Error("Expected approved listing's candidate to have been removed")
	}

	// member can unstake
	_, unErr := deployed.VotingContract.Unstake(test.GetTxOpts(context.AuthUser2, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash)
	test.IfNotNil(t, unErr, fmt.Sprintf("Error Unstaking: %v", unErr))
	context.Blockchain.Commit()
}

func TestResolveApplicationThatFails(t *testing.T) {
	// check the market's balance as it should not change in a no-list op
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)

	listingHash := test.GenBytes32("BarMarket, CA.")

	_, listErr := deployed.ListingContract.List(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
	context.Blockchain.Commit()

	// move past the voteBy with no votes being cast
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	// any council member can call for resolution
	_, resErr := deployed.ListingContract.ResolveApplication(test.GetTxOpts(context.AuthUser2, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, resErr, fmt.Sprintf("Error resolving application: %v", resErr))
	context.Blockchain.Commit()

	// should not be listed now, and no reward
	owner, supply, _ := deployed.ListingContract.GetListing(nil, listingHash)

	if owner == context.AuthUser1.From {
		t.Errorf("Exepected owner to be zero address, got: %v", owner)
	}

	if supply.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Exepected supply to be 0, got: %v", supply)
	}

	// post resolution any candidate should have been removed
	isCandidate, _ := deployed.VotingContract.IsCandidate(nil, listingHash)

	if isCandidate == true {
		t.Error("Expected failed listing's candidate to have been removed")
	}

	newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)
	if marketBal.Cmp(newMarketBal) != 0 {
		t.Errorf("Expected %v to be %v", marketBal, newMarketBal)
	}
}

func TestWithdrawFromListing(t *testing.T) {
	// the user's balance should increase after a withdrawal
	userBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser1.From)
	// check the market's balance...
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)
	// foomarket itself
	listingHash := test.GenBytes32("FooMarket, AZ.")
	// the current state of foo market
	_, supply, _ := deployed.ListingContract.GetListing(nil, listingHash)
	// withdraw that same amt
	_, withErr := deployed.ListingContract.WithdrawFromListing(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash, big.NewInt(test.ONE_ETH*0.01))
	test.IfNotNil(t, withErr, fmt.Sprintf("Error withdrawing from listing: %v", withErr))
	context.Blockchain.Commit()

	// marketBal should have decreased
	newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)
	if newMarketBal.Cmp(marketBal) != -1 {
		t.Errorf("Expected %v to be > %v", marketBal, newMarketBal)
	}

	_, newSupply, _ := deployed.ListingContract.GetListing(nil, listingHash)
	if newSupply.Cmp(supply) != -1 {
		t.Errorf("Expected %v to be > %v", supply, newSupply)
	}

	newUserBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser1.From)
	if newUserBal.Cmp(userBal) != 1 {
		t.Errorf("Expected %v to be > %v", newUserBal, userBal)
	}
}

func TestExit(t *testing.T) {
	// going to remove this one
	listingHash := test.GenBytes32("FooMarket, AZ.")
	// is it listed?
	listed, _ := deployed.ListingContract.IsListed(nil, listingHash)
	if !listed {
		t.Errorf("Expected %v to be true", listed)
	}
	// we know that the market's bank will decrease here by the listReward
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)

	// there is a data_hash for this listing
	dataHash, _ := deployed.DatatrustContract.GetDataHash(nil, listingHash)

	_, exitErr := deployed.ListingContract.Exit(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, exitErr, fmt.Sprintf("Error exiting listing from market, %v", exitErr))
	context.Blockchain.Commit()

	listedNow, _ := deployed.ListingContract.IsListed(nil, listingHash)
	if listedNow {
		t.Errorf("Expected %v to be false", listedNow)
	}

	updatedMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)
	if updatedMarketBal.Cmp(marketBal) != -1 {
		t.Errorf("Expected %v to be > %v", marketBal, updatedMarketBal)
	}

	// data_hash should be cleared
	dataHashNow, _ := deployed.DatatrustContract.GetDataHash(nil, listingHash)
	if dataHashNow == dataHash {
		t.Errorf("Expected data hash to be empty, got: %v", dataHashNow)
	}
}
