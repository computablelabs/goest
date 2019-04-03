package listingtest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
	"time"
)

// the challenge wins in this spec...

func TestChallenge(t *testing.T) {
	_, listErr := deployed.ListingContract.List(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 250000), "BarMarket12345")
	test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
	context.Blockchain.Commit()

	listingHash, _ := deployed.ListingContract.GetHash(nil, "BarMarket12345")

	// data_hash needs to be set first
	dataHash, _ := deployed.DatatrustContract.GetHash(nil, "thisissomemoardata")
	_, dataErr := deployed.DatatrustContract.SetDataHash(test.GetTxOpts(context.AuthBackend,
		nil, big.NewInt(test.ONE_GWEI*2), 100000), listingHash, dataHash)
	test.IfNotNil(t, dataErr, "Error setting data hash for listing")
	context.Blockchain.Commit()

	// user may need funds
	transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthFactory,
		context.AuthUser3.From, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr, "Error transferring market token")
	// member will need to have approved the voting contract to spend
	appErr := test.MaybeIncreaseMarketTokenApproval(context, deployed,
		context.AuthUser3, deployed.VotingAddress, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, appErr, "Error increasing allowance")

	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))
	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
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

	// member can unstake
	_, unErr := deployed.VotingContract.Unstake(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash)
	test.IfNotNil(t, unErr, fmt.Sprintf("Error Unstaking: %v", unErr))
	context.Blockchain.Commit()

	// member 2 as challenger here, may need funds
	transErr2 := test.MaybeTransferMarketToken(context, deployed, context.AuthFactory,
		context.AuthUser2.From, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr2, "Error transferring market token")
	// member will need to have approved the voting contract to spend
	appErr2 := test.MaybeIncreaseMarketTokenApproval(context, deployed,
		context.AuthUser2, deployed.VotingAddress, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, appErr2, "Error increasing allowance")

	// note our balances right now as challenging will change them
	memberBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser2.From)
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.VotingAddress)

	_, challengeErr := deployed.ListingContract.Challenge(test.GetTxOpts(context.AuthUser2,
		nil, big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, challengeErr, fmt.Sprintf("Error challenging listing: %v", challengeErr))
	context.Blockchain.Commit()

	// memberBal should have decreased due to challengeStake being locked
	newMemberBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser2.From)
	if memberBal.Cmp(newMemberBal) != 1 {
		t.Errorf("Expected %v to be > %v", memberBal, newMemberBal)
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

func TestResolveChallenge(t *testing.T) {
	listingHash, _ := deployed.ListingContract.GetHash(nil, "BarMarket12345")
	// the marketBal will decrease as the challenger unstakes post resolution
	votingBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.VotingAddress)
	memberBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser2.From)

	// we need to cast a vote for the challenge, or the listing will "win". member3 as voter here
	appErr := test.MaybeIncreaseMarketTokenApproval(context, deployed, context.AuthUser3,
		deployed.VotingAddress, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, appErr, fmt.Sprintf("Error approving voting contract to spend: %v", appErr))
	context.Blockchain.Commit()

	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))
	context.Blockchain.Commit()

	_, _, _, _, yea, _, _ := deployed.VotingContract.GetCandidate(nil, listingHash)

	if yea.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected at least one vote, got: %v", yea)
	}

	// the vote is recorded now, and the challenge should be ready to be resolved once we pass the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	_, resolveErr := deployed.ListingContract.ResolveChallenge(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving challenge: %v", resolveErr))
	context.Blockchain.Commit()

	// should no longer be a listing
	listed, _ := deployed.ListingContract.IsListed(nil, listingHash)
	if listed == true {
		t.Error("Expected .listed to be false")
	}

	// the candidate should have been pruned
	isCan, _ := deployed.VotingContract.IsCandidate(nil, listingHash)
	if isCan {
		t.Error("Expected candidate to have been removed")
	}

	// challenger may unstake now
	_, unErr := deployed.VotingContract.Unstake(test.GetTxOpts(context.AuthUser2, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, unErr, fmt.Sprintf("Error unstaking: %v", unErr))
	// voter may also unstake
	_, un2Err := deployed.VotingContract.Unstake(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, un2Err, fmt.Sprintf("Error unstaking: %v", un2Err))
	context.Blockchain.Commit()

	votingBalNow, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.VotingAddress)
	memberBalNow, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser2.From)

	// market went down...
	if votingBalNow.Cmp(votingBal) != -1 {
		t.Errorf("Expected %v to be < %v", votingBalNow, votingBal)
	}

	// user bal went up...
	if memberBalNow.Cmp(memberBal) != 1 {
		t.Errorf("Expected %v to be > %v", memberBalNow, memberBal)
	}

	// the stakes are cleared
	stake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser2.From)
	if stake.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected stake to be 0, got: %v", stake)
	}

	stake2, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser3.From)
	if stake2.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected stake to be 0, got: %v", stake2)
	}
}

func TestLosingChallenge(t *testing.T) {
	// get a listing up
	_, listErr := deployed.ListingContract.List(test.GetTxOpts(context.AuthUser1,
		nil, big.NewInt(test.ONE_GWEI*2), 250000), "BazDataXYZ")
	test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
	context.Blockchain.Commit()

	listingHash, _ := deployed.ListingContract.GetHash(nil, "BazDataXYZ")

	// data_hash needs to be set first
	dataHash, _ := deployed.DatatrustContract.GetHash(nil, "thisisallthedata")
	_, dataErr := deployed.DatatrustContract.SetDataHash(test.GetTxOpts(context.AuthBackend,
		nil, big.NewInt(test.ONE_GWEI*2), 100000), listingHash, dataHash)
	test.IfNotNil(t, dataErr, "Error setting data hash for listing")
	context.Blockchain.Commit()

	// user 3 as voter
	appErr := test.MaybeIncreaseMarketTokenApproval(context, deployed,
		context.AuthUser3, deployed.VotingAddress, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, appErr, "Error increasing allowance")

	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))
	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	// call for resolution
	_, resolveErr := deployed.ListingContract.ResolveApplication(test.GetTxOpts(
		context.AuthUser2, nil, big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving application: %v", resolveErr))
	context.Blockchain.Commit()

	// challenge it -- we won't vote for the chall so that it loses...
	appErr2 := test.MaybeIncreaseMarketTokenApproval(context, deployed,
		context.AuthUser2, deployed.VotingAddress, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, appErr2, "Error increasing allowance")

	_, challengeErr := deployed.ListingContract.Challenge(test.GetTxOpts(context.AuthUser2, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, challengeErr, fmt.Sprintf("Error challenging listing: %v", challengeErr))
	context.Blockchain.Commit()

	// note balances as we want to measure after...
	member1Bal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser1.From)
	member1Stake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser1.From)
	member2Stake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser2.From)
	member3Stake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser3.From)

	// we know member2 has staked now
	if member2Stake.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected member 2 stake to be > 0, got: %v", member2Stake)
	}

	// and 3...
	if member3Stake.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected member 3 stake to be > 0, got: %v", member3Stake)
	}

	// we know member1 has not staked now
	if member1Stake.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected member 2 stake to be > 0, got: %v", member2Stake)
	}

	// move past  the vote by
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	_, resolveChallErr := deployed.ListingContract.ResolveChallenge(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, resolveChallErr, fmt.Sprintf("Error resolving challenge: %v", resolveChallErr))
	context.Blockchain.Commit()

	// member 2 should have lost that stake
	member2StakeNow, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser2.From)
	if member2StakeNow.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected member 2 stake to be 0, got: %v", member2StakeNow)
	}

	// member 1 should be credited that stake now
	member1StakeNow, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser1.From)
	if member1StakeNow.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected member 1 stake to be > 0, got: %v", member1StakeNow)
	}

	// member 3 stake is unchanged
	member3StakeNow, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser3.From)
	if member3StakeNow.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected member 3 stake to be > 0, got: %v", member3StakeNow)
	}

	// member 1 can actually unstake that credit
	_, unErr := deployed.VotingContract.Unstake(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, unErr, fmt.Sprintf("Error unstaking: %v", unErr))

	// member 3 too
	_, un3Err := deployed.VotingContract.Unstake(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, un3Err, fmt.Sprintf("Error unstaking: %v", un3Err))
	context.Blockchain.Commit()

	// member 1 has a higher bal now
	member1BalNow, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser1.From)
	if member1BalNow.Cmp(member1Bal) != 1 {
		t.Errorf("Expected %v to be > %v", member1BalNow, member1Bal)
	}

	// the stakes are cleared
	stake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser1.From)
	if stake.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Expected stake to be 0, got: %v", stake)
	}

	stake3, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser3.From)
	if stake3.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Expected stake 3 to be 0, got: %v", stake3)
	}
}
