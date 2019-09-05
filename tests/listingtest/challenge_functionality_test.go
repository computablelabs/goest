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
	listingHash := test.GenBytes32("BarMarket12345")
	_, listErr := deployed.ListingContract.List(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 250000), listingHash)
	test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
	context.Blockchain.Commit()

	// data_hash needs to be set first
	dataHash, _ := deployed.DatatrustContract.GetHash(nil, "thisissomemoardata")
	_, dataErr := deployed.DatatrustContract.SetDataHash(test.GetTxOpts(context.AuthBackend,
		nil, big.NewInt(test.ONE_GWEI*2), 100000), listingHash, dataHash)
	test.IfNotNil(t, dataErr, "Error setting data hash for listing")
	context.Blockchain.Commit()

	// user may need funds
	transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner,
		context.AuthUser3.From, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr, "Error transferring market token")
	// member will need to have approved the voting contract to spend
	appErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed,
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
	transErr2 := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner,
		context.AuthUser2.From, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr2, "Error transferring market token")
	// member will need to have approved the voting contract to spend
	appErr2 := test.MaybeIncreaseMarketTokenAllowance(context, deployed,
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
	listingHash := test.GenBytes32("BarMarket12345")
	// the listing has some marketToken credited in its supply, note it
	_, supply, _ := deployed.ListingContract.GetListing(nil, listingHash)
	if supply.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected supply to be > 0, got: %v", supply)
	}
	// the voting Bal will decrease as the challenger unstakes post resolution
	votingBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.VotingAddress)
	ownerBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser1.From)
	memberBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser2.From)

	// we need to cast a vote for the challenge, or the listing will "win". member3 as voter here
	appErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser3,
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

	// the supply is gone of course
	_, supplyNow, _ := deployed.ListingContract.GetListing(nil, listingHash)
	if supplyNow.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected supply to be 0, got: %v", supplyNow)
	}

	// note that the listing supply was not burned, and instead transferred
	// to the owner on removal
	ownerBalNow, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser1.From)
	if ownerBalNow.Cmp(ownerBal) != 1 {
		t.Errorf("Expected %v to be > %v", ownerBalNow, ownerBal)
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
	listingHash := test.GenBytes32("BazDataXYZ")
	_, listErr := deployed.ListingContract.List(test.GetTxOpts(context.AuthUser1,
		nil, big.NewInt(test.ONE_GWEI*2), 250000), listingHash)
	test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
	context.Blockchain.Commit()

	// data_hash needs to be set first
	dataHash, _ := deployed.DatatrustContract.GetHash(nil, "thisisallthedata")
	_, dataErr := deployed.DatatrustContract.SetDataHash(test.GetTxOpts(context.AuthBackend,
		nil, big.NewInt(test.ONE_GWEI*2), 100000), listingHash, dataHash)
	test.IfNotNil(t, dataErr, "Error setting data hash for listing")
	context.Blockchain.Commit()

	// user 3 as voter
	appErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed,
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
	appErr2 := test.MaybeIncreaseMarketTokenAllowance(context, deployed,
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

	// NOTE leave user 1's stake so that, in a subsequent victory, we can see how the stake is not aggregated
	// _, unErr := deployed.VotingContract.Unstake(test.GetTxOpts(context.AuthUser1, nil,
	// big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	// test.IfNotNil(t, unErr, fmt.Sprintf("Error unstaking: %v", unErr))

	// member 3 unstakes
	_, un3Err := deployed.VotingContract.Unstake(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, un3Err, fmt.Sprintf("Error unstaking: %v", un3Err))
	context.Blockchain.Commit()

	// member 1, having not unstaked, has the same bal
	member1BalNow, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser1.From)
	if member1BalNow.Cmp(member1Bal) != 0 {
		t.Errorf("Expected %v to be %v", member1BalNow, member1Bal)
	}

	// the stake for user 1 is not cleared
	stake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser1.From)
	if stake.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected stake to be > 0, got: %v", stake)
	}

	// user 3 is cleared
	stake3, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser3.From)
	if stake3.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected stake 3 to be 0, got: %v", stake3)
	}

	// is still a listing
	listed, _ := deployed.ListingContract.IsListed(nil, listingHash)
	if listed != true {
		t.Errorf("Expected .listed to be true, got: %v", listed)
	}
}

// show that any pre-existing stake winnings (user1 enters this spec with an existing stake won) can be overwritten by future stake winnings
// at the same `voting.address.hash`
func TestStakeOverwrite(t *testing.T) {
	// we need to generate the same hashes as before so that the `voting.address.hash` will overwrite
	listingHash := test.GenBytes32("BazDataXYZ")
	// get the actual p11r number
	pStake, _ := deployed.ParameterizerContract.GetStake(nil)

	user1Stake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser1.From)
	if user1Stake.Cmp(pStake) != 0 {
		t.Errorf("Expected stake to be %v, got: %v", pStake, user1Stake)
	}

	// challenge listing the same was as the previous spec
	transErr2 := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner,
		context.AuthUser2.From, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr2, "Error transferring market token")
	// member will need to have approved the voting contract to spend
	appErr2 := test.MaybeIncreaseMarketTokenAllowance(context, deployed,
		context.AuthUser2, deployed.VotingAddress, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, appErr2, "Error increasing allowance")

	_, challengeErr := deployed.ListingContract.Challenge(test.GetTxOpts(context.AuthUser2, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, challengeErr, fmt.Sprintf("Error challenging listing: %v", challengeErr))
	context.Blockchain.Commit()

	// the challenge exists
	isCan, _ := deployed.VotingContract.IsCandidate(nil, listingHash)
	if isCan != true {
		t.Errorf("Expected %v to be true", isCan)
	}

	// note balances as we want to measure after...
	member2Stake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser2.From)

	// we know member2 has staked now
	if member2Stake.Cmp(pStake) != 0 {
		t.Errorf("Expected member 2 stake to be %v, got: %v", pStake, member2Stake)
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

	// member 1 should be credited that stake now, but notice the 2 stakes have aggregated
	user1StakeNow, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser1.From)
	if user1StakeNow.Cmp(pStake) != 1 {
		t.Errorf("Expected member 1 stake to be > %v, got: %v", pStake, user1StakeNow)
	}
}

// show that a user who votes about a list candidate -> list passes -> leave their stake -> challenge listing will overwrite their original vote stake
// that they did not unstake...
func TestVoteStakeOverwrite(t *testing.T) {
	listingHash := test.GenBytes32("QuxMarket00")
	// user3 has nothing staked at this hash
	user3Stake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser3.From)
	if user3Stake.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected user 3 stake to be 0, got: %v", user3Stake)
	}
	t.Log(fmt.Sprintf("user3 stake is %v", user3Stake))
	// get that hash up as a listing
	_, listErr := deployed.ListingContract.List(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 250000), listingHash)
	test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
	context.Blockchain.Commit()

	// data_hash needs to be set first
	dataHash, _ := deployed.DatatrustContract.GetHash(nil, "thisissomemoardata")
	_, dataErr := deployed.DatatrustContract.SetDataHash(test.GetTxOpts(context.AuthBackend,
		nil, big.NewInt(test.ONE_GWEI*2), 100000), listingHash, dataHash)
	test.IfNotNil(t, dataErr, "Error setting data hash for listing")
	context.Blockchain.Commit()

	// user may need funds
	transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner,
		context.AuthUser3.From, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr, "Error transferring market token")
	// member will need to have approved the voting contract to spend
	appErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed,
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

	// the current stake amt
	pStake, _ := deployed.ParameterizerContract.GetStake(nil)
	t.Log(fmt.Sprintf("pstake is now %v", pStake))
	// we'll leave user 3's stake in place...
	user3StakeNow, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser3.From)
	if user3StakeNow.Cmp(pStake) != 0 {
		t.Errorf("Expected %v to be %v", user3StakeNow, pStake)
	}
	t.Log(fmt.Sprintf("user3 stake is now %v", user3StakeNow))

	// user3 decides to challenge the listing
	transErr2 := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner,
		context.AuthUser3.From, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr2, "Error transferring market token")
	// member will need to have approved the voting contract to spend
	appErr2 := test.MaybeIncreaseMarketTokenAllowance(context, deployed,
		context.AuthUser3, deployed.VotingAddress, big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, appErr2, "Error increasing allowance")

	// note the user has market token pre-challenge
	userBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser3.From)
	if userBal.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected %v to be > 0", userBal)
	}

	_, challengeErr := deployed.ListingContract.Challenge(test.GetTxOpts(context.AuthUser3,
		nil, big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
	test.IfNotNil(t, challengeErr, fmt.Sprintf("Error challenging listing: %v", challengeErr))
	context.Blockchain.Commit()

	// note the user has no market token post-challenge
	userBalNow, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthUser3.From)
	if userBalNow.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected user balance to be 0, got: %v", userBalNow)
	}

	// note that the stake is present, but did not aggregate
	user3StakeAgain, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthUser3.From)
	if user3StakeAgain.Cmp(pStake) != 1 {
		t.Errorf("Expected %v to be > %v", user3StakeAgain, pStake)
	}
	t.Log(fmt.Sprintf("user3 stake after challenge is %v", user3StakeAgain))
}
