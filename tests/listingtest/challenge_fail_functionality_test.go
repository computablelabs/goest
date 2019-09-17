package listingtest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
	"time"
)

func TestFailedChallenge(t *testing.T) {
	// member1 submits the new listing candidate
	listingHash := test.GenBytes32("FooMarket12345")
	_, listErr := deployed.ListingContract.List(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 250000), listingHash)
	test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
	context.Blockchain.Commit()

	// data_hash needs to be set first
	dataHash, _ := deployed.DatatrustContract.GetHash(nil, "gooddata")
	_, dataErr := deployed.DatatrustContract.SetDataHash(test.GetTxOpts(context.AuthBackend,
		nil, big.NewInt(test.ONE_GWEI*2), 100000), listingHash, dataHash)
	test.IfNotNil(t, dataErr, "Error setting data hash for listing")
	context.Blockchain.Commit()

	// TODO(rbharath): Is this really needed?
	// member3 as voter may need funds
	transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner,
		context.AuthUser3.From, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr, "Error transferring market token")
	// member will need to have approved the voting contract to spend
	appErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed,
		context.AuthUser3, deployed.VotingAddress, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, appErr, "Error increasing allowance")

	// member3 votes the listing in
	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))
	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(test.MIN_VOTE_BY * time.Second)
	context.Blockchain.Commit()

	// call for resolution
	_, resolveErr := deployed.ListingContract.ResolveApplication(test.GetTxOpts(context.AuthUser2,
		nil, big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving application: %v", resolveErr))
	context.Blockchain.Commit()

	// should be listed now
	owner, _, _ := deployed.ListingContract.GetListing(nil, listingHash)
	if owner != context.AuthUser1.From {
		t.Errorf("Exepected listing to be owned by member 1, got: %v", owner)
	}

	// member3 can unstake
	_, unErr := deployed.VotingContract.Unstake(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash)
	test.IfNotNil(t, unErr, fmt.Sprintf("Error Unstaking: %v", unErr))
	context.Blockchain.Commit()

	// member 2 as challenger here, may need funds
	transErr2 := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner,
		context.AuthUser2.From, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr2, "Error transferring market token")
	// member will need to have approved the voting contract to spend
	appErr2 := test.MaybeIncreaseMarketTokenAllowance(context, deployed,
		context.AuthUser2, deployed.VotingAddress, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, appErr2, "Error increasing allowance")

	// note our balances right now as challenging will change them
	member2Bal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser2.From)
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.VotingAddress)

	// member2 issues challenge
	_, challengeErr := deployed.ListingContract.Challenge(test.GetTxOpts(context.AuthUser2,
		nil, big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, challengeErr, fmt.Sprintf("Error challenging listing: %v", challengeErr))
	context.Blockchain.Commit()

	// member2Bal should have decreased due to challengeStake being locked
	newMember2Bal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser2.From)
	if member2Bal.Cmp(newMember2Bal) != 1 {
		t.Errorf("Expected %v to be > %v", member2Bal, newMember2Bal)
	}

	// market then goes up as it banks that amount...
	newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.VotingAddress)
	if newMarketBal.Cmp(marketBal) != 1 {
		t.Errorf("Expected %v to be > %v", newMarketBal, marketBal)
	}

	// any staked amounts for a hash can be fetched
	stake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser2.From)
	if stake.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected stake to be > 0, got: %v", stake)
	}
}

func TestResolveFirstFailedChallenge(t *testing.T) {

	listingHash := test.GenBytes32("FooMarket12345")
	// the listing has some marketToken credited in its supply, note it
	_, supply, _ := deployed.ListingContract.GetListing(nil, listingHash)
	if supply.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected supply to be > 0, got: %v", supply)
	}

	// The challenger stake should not be 0 now
	challengerStake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser2.From)
	if challengerStake.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected challengerStake to be > 0, got: %v", challengerStake)
	}

	// The listing owner stake should be 0 now
	ownerStake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser1.From)
	if ownerStake.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected ownerStake to be 0, got: %v", ownerStake)
	}

	// Let's cast a vote against the challenge, so the listing will "win".
	// member3 here. (Note that not voting would result in the listing winning
	// just as well)
	appErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser3,
		deployed.VotingAddress, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, appErr, fmt.Sprintf("Error approving voting contract to spend: %v", appErr))
	context.Blockchain.Commit()

	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, big.NewInt(0))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))
	context.Blockchain.Commit()

	// the vote is recorded now, and the challenge should be ready to be resolved once we pass the voteBy
	context.Blockchain.AdjustTime(test.MIN_VOTE_BY * time.Second)
	context.Blockchain.Commit()

	// Let's now resolve the challenge
	_, resolveErr := deployed.ListingContract.ResolveChallenge(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving challenge: %v", resolveErr))
	context.Blockchain.Commit()

	// should still be a listing
	listed, _ := deployed.ListingContract.IsListed(nil, listingHash)
	if listed == false {
		t.Error("Expected .listed to be true")
	}

	// The challenger stake should be 0 now
	challengerStakeNow, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser2.From)
	if challengerStakeNow.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected challengerStakeNow to be 0, got: %v", challengerStakeNow)
	}

	// The listing owner stake should be > 0 now
	ownerStakeNow, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser1.From)
	if ownerStakeNow.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected ownerStake to be > 0, got: %v", ownerStakeNow)
	}
}

func TestSecondFailedChallenge(t *testing.T) {
	// Now let's have a second challenge fail.
	listingHash := test.GenBytes32("FooMarket12345")
	// the listing has some marketToken credited in its supply, note it
	_, supply, _ := deployed.ListingContract.GetListing(nil, listingHash)
	if supply.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected supply to be > 0, got: %v", supply)
	}

	// member 2 as repeat challenger here, may need funds
	transErr2 := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner,
		context.AuthUser2.From, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr2, "Error transferring market token")
	// member will need to have approved the voting contract to spend
	appErr2 := test.MaybeIncreaseMarketTokenAllowance(context, deployed,
		context.AuthUser2, deployed.VotingAddress, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, appErr2, "Error increasing allowance")

	// member2 issues second challenge
	_, challengeErr := deployed.ListingContract.Challenge(test.GetTxOpts(context.AuthUser2,
		nil, big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, challengeErr, fmt.Sprintf("Error challenging listing: %v", challengeErr))
	context.Blockchain.Commit()

	// The challenger stake should not be 0 now
	challengerStake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser2.From)
	if challengerStake.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected challengerStake to be > 0, got: %v", challengerStake)
	}

	// The listing owner stake should not be 0 now
	// (since we haven't unstaked proceeds from last failed challenge)
	ownerStake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser1.From)
	if ownerStake.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected ownerStake to be > 0, got: %v", ownerStake)
	}

	// Let's cast a vote against the challenge, so the listing will "win".
	// member3 here. (Note that not voting would result in the listing winning
	// just as well)
	appErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser3,
		deployed.VotingAddress, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, appErr, fmt.Sprintf("Error approving voting contract to spend: %v", appErr))
	context.Blockchain.Commit()

	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, big.NewInt(0))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))
	context.Blockchain.Commit()

	// the vote is recorded now, and the challenge should be ready to be resolved once we pass the voteBy
	context.Blockchain.AdjustTime(test.MIN_VOTE_BY * time.Second)
	context.Blockchain.Commit()

	// Let's now resolve the challenge
	_, resolveErr := deployed.ListingContract.ResolveChallenge(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving challenge: %v", resolveErr))
	context.Blockchain.Commit()

	// should still be a listing
	listed, _ := deployed.ListingContract.IsListed(nil, listingHash)
	if listed == false {
		t.Error("Expected .listed to be true")
	}

	// The challenger stake should be 0 now
	challengerStakeNow, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser2.From)
	if challengerStakeNow.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected challengerStakeNow to be 0, got: %v", challengerStakeNow)
	}

	// The listing owner stake should have increased now
	ownerStakeNow, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser1.From)
	if ownerStakeNow.Cmp(ownerStake) != 1 {
		t.Errorf("Expected ownerStake to increment, got: %v", ownerStakeNow)
	}

}
