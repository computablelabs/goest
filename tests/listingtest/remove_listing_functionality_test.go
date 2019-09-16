package listingtest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
	"time"
)

func TestAccidentalList(t *testing.T) {
	// In this scenario, we have a listing which is listed, but which the
	// owner probably didn't want to list.
	listingHash := test.GenBytes32("Oopsie, don't list me.")

	_, listErr := deployed.ListingContract.List(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 250000), listingHash)
	test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
	context.Blockchain.Commit()

	isCan, _ := deployed.VotingContract.IsCandidate(nil, listingHash)

	if !isCan {
		t.Errorf("Expected isCandidate to be true, got: %v", isCan)
	}

	// the datatrust must have a data_hash for this listing before it will pass
	dataHash, _ := deployed.DatatrustContract.GetHash(nil, "thisissomedata")
	_, dataErr := deployed.DatatrustContract.SetDataHash(test.GetTxOpts(context.AuthBackend, nil,
		big.NewInt(test.ONE_GWEI*2), 100000), listingHash, dataHash)
	test.IfNotNil(t, dataErr, "Error setting data hash for listing")
	context.Blockchain.Commit()

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

	// Now cast the actual vote
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

}

func TestRemoveList(t *testing.T) {
	// Now we're going to remove the listing which that user probably didn't
	// want to have listed.

	// Let's start by getting the listing hash gain
	listingHash := test.GenBytes32("Oopsie, don't list me.")

	// Let's note the owner's balance
	ownerBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser1.From)

	// Let's note the market balance now
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)

	// Now let's actually remove the listing
	_, exitErr := deployed.ListingContract.Exit(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, exitErr, fmt.Sprintf("Error exiting listing from market, %v", exitErr))
	context.Blockchain.Commit()

	listedNow, _ := deployed.ListingContract.IsListed(nil, listingHash)
	if listedNow {
		t.Errorf("Expected %v to be false", listedNow)
	}

	// Let's note the owner's balance after the listing was removed. This
	// should have gone up.
	ownerBalNow, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser1.From)
	if ownerBalNow.Cmp(ownerBal) != 1 {
		t.Errorf("Exepected new owner balance %v to be > old owner balance %v", ownerBalNow, ownerBal)
	}

	// The market's balance should have gone down accordingly
	updatedMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)
	if updatedMarketBal.Cmp(marketBal) != -1 {
		t.Errorf("Expected %v to be > %v", marketBal, updatedMarketBal)
	}

}
