package listing

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
	"time"
)

// the challenge wins in this spec...

func TestChallenge(t *testing.T) {
	_, listErr := deployed.ListingContract.List(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 250000,
	}, "BarMarket12345")

	if listErr != nil {
		t.Fatalf("Error applying for list status: %v", listErr)
	}

	context.Blockchain.Commit()

	listingHash, _ := deployed.ListingContract.GetHash(nil, "BarMarket12345")

	// data_hash needs to be set first
	dataHash, _ := deployed.DatatrustContract.GetHash(nil, "thisissomemoardata")
	_, dataErr := deployed.DatatrustContract.SetDataHash(&bind.TransactOpts{
		From:     context.AuthBackend.From,
		Signer:   context.AuthBackend.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 100000,
	}, listingHash, dataHash)

	if dataErr != nil {
		t.Fatal("Error setting data hash for listing")
	}
	context.Blockchain.Commit()

	_, voteErr := deployed.VotingContract.Vote(&bind.TransactOpts{
		From:     context.AuthMember2.From,
		Signer:   context.AuthMember2.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 150000,
	}, listingHash, big.NewInt(1))

	if voteErr != nil {
		t.Fatalf("Error voting for candidate: %v", voteErr)
	}

	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	// call for resolution
	_, resolveErr := deployed.ListingContract.ResolveApplication(&bind.TransactOpts{
		From:     context.AuthMember2.From,
		Signer:   context.AuthMember2.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, listingHash)

	if resolveErr != nil {
		t.Fatalf("Error resolving application: %v", resolveErr)
	}

	context.Blockchain.Commit()

	// should be listed now, with a reward
	owner, _, _, _ := deployed.ListingContract.GetListing(nil, listingHash)

	if owner != context.AuthMember1.From {
		t.Fatalf("Exepected listing to be owned by member 1, got: %v", owner)
	}

	// now we can challenge. member2 will need funds approved && such
	_, transErr := deployed.MarketTokenContract.Transfer(&bind.TransactOpts{
		From:     context.AuthFactory.From,
		Signer:   context.AuthFactory.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, context.AuthMember2.From, big.NewInt(ONE_WEI*2)) // 2 tokens

	if transErr != nil {
		t.Fatalf("Error transferring tokens to member: %v", transErr)
	}

	context.Blockchain.Commit()

	// the transfers happen at the voting contract for challenge
	_, approveErr := deployed.MarketTokenContract.Approve(&bind.TransactOpts{
		From:     context.AuthMember2.From,
		Signer:   context.AuthMember2.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, deployed.VotingAddress, big.NewInt(ONE_WEI*2)) // up to 2 tokenWei

	if approveErr != nil {
		t.Fatalf("Error approving market contract to spend: %v", approveErr)
	}

	context.Blockchain.Commit()

	// note our balances right now as challenging will change them
	memberBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember2.From)
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.VotingAddress)

	_, challengeErr := deployed.ListingContract.Challenge(&bind.TransactOpts{
		From:     context.AuthMember2.From,
		Signer:   context.AuthMember2.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, listingHash)

	if challengeErr != nil {
		t.Fatalf("Error challenging listing: %v", challengeErr)
	}

	context.Blockchain.Commit()

	// memberBal should have decreased due to challengeStake being locked
	newMemberBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember2.From)
	if memberBal.Cmp(newMemberBal) != 1 {
		t.Fatalf("Expected %v to be > %v", memberBal, newMemberBal)
	}
	// market then goes up as it banks that amount...
	newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.VotingAddress)
	if newMarketBal.Cmp(marketBal) != 1 {
		t.Fatalf("Expected %v to be > %v", newMarketBal, marketBal)
	}

	// any staked amounts for a hash can be fetched
	stake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthMember2.From)
	if stake.Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("Expected stake to be > 0, got: %v", stake)
	}
}

func TestResolveChallenge(t *testing.T) {
	listingHash, _ := deployed.ListingContract.GetHash(nil, "BarMarket12345")
	// the marketBal will decrease as the challenger unstakes post resolution
	votingBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.VotingAddress)
	memberBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember2.From)

	// we need to cast a vote for the challenge, or the listing will "win"
	_, voteErr := deployed.VotingContract.Vote(&bind.TransactOpts{
		From:     context.AuthMember2.From,
		Signer:   context.AuthMember2.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, listingHash, big.NewInt(1))

	if voteErr != nil {
		t.Fatalf("Error voting for candidate: %v", voteErr)
	}

	context.Blockchain.Commit()

	_, _, _, _, yea, _, _ := deployed.VotingContract.GetCandidate(nil, listingHash)

	if yea.Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("Expected at least one vote, got: %v", yea)
	}

	// the vote is recorded now, and the challenge should be ready to be resolved once we pass the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	_, resolveErr := deployed.ListingContract.ResolveChallenge(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, listingHash)

	if resolveErr != nil {
		t.Fatalf("Error resolving challenge: %v", resolveErr)
	}

	context.Blockchain.Commit()

	// should no longer be a listing
	listed, _ := deployed.ListingContract.IsListed(nil, listingHash)

	if listed == true {
		t.Fatalf("Expected .listed to be false, got")
	}

	// the candidate should have been pruned
	isCan, _ := deployed.VotingContract.IsCandidate(nil, listingHash)
	if isCan {
		t.Fatal("Expected candidate to have been removed")
	}

	// challenger may unstake now
	_, unErr := deployed.VotingContract.Unstake(&bind.TransactOpts{
		From:     context.AuthMember2.From,
		Signer:   context.AuthMember2.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, listingHash)

	if unErr != nil {
		t.Fatalf("Error unstaking: %v", unErr)
	}

	context.Blockchain.Commit()

	votingBalNow, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.VotingAddress)
	memberBalNow, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember2.From)

	// market went down...
	if votingBalNow.Cmp(votingBal) != -1 {
		t.Fatalf("Expected %v to be < %v", votingBalNow, votingBal)
	}

	// user bal went up...
	if memberBalNow.Cmp(memberBal) != 1 {
		t.Fatalf("Expected %v to be > %v", memberBalNow, memberBal)
	}

	// the stake is cleared
	stake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthMember2.From)
	if stake.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Expected stake to be 0, got: %v", stake)
	}
}

func TestLosingChallenge(t *testing.T) {
	// get a listing up
	_, listErr := deployed.ListingContract.List(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 250000,
	}, "BazDataXYZ")

	if listErr != nil {
		t.Fatalf("Error applying for list status: %v", listErr)
	}

	context.Blockchain.Commit()

	listingHash, _ := deployed.ListingContract.GetHash(nil, "BazDataXYZ")

	// data_hash needs to be set first
	dataHash, _ := deployed.DatatrustContract.GetHash(nil, "thisisallthedata")
	_, dataErr := deployed.DatatrustContract.SetDataHash(&bind.TransactOpts{
		From:     context.AuthBackend.From,
		Signer:   context.AuthBackend.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 100000,
	}, listingHash, dataHash)

	if dataErr != nil {
		t.Fatal("Error setting data hash for listing")
	}
	context.Blockchain.Commit()

	_, voteErr := deployed.VotingContract.Vote(&bind.TransactOpts{
		From:     context.AuthMember2.From,
		Signer:   context.AuthMember2.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 150000,
	}, listingHash, big.NewInt(1))

	if voteErr != nil {
		t.Fatalf("Error voting for candidate: %v", voteErr)
	}

	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	// call for resolution
	_, resolveErr := deployed.ListingContract.ResolveApplication(&bind.TransactOpts{
		From:     context.AuthMember2.From,
		Signer:   context.AuthMember2.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, listingHash)

	if resolveErr != nil {
		t.Fatalf("Error resolving application: %v", resolveErr)
	}

	context.Blockchain.Commit()

	// challenge it -- we won't vote for the chall so that it loses...
	_, challengeErr := deployed.ListingContract.Challenge(&bind.TransactOpts{
		From:     context.AuthMember2.From,
		Signer:   context.AuthMember2.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, listingHash)

	if challengeErr != nil {
		t.Fatalf("Error challenging listing: %v", challengeErr)
	}

	context.Blockchain.Commit()

	// note balances as we want to measure after...
	member1Bal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember1.From)
	member1Stake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthMember1.From)
	member2Stake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthMember2.From)

	// we know member2 has staked now
	if member2Stake.Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("Expected member 2 stake to be > 0, got: %v", member2Stake)
	}

	// we know member1 has not staked now
	if member1Stake.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Expected member 2 stake to be > 0, got: %v", member2Stake)
	}

	// move past  the vote by
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	_, resolveChallErr := deployed.ListingContract.ResolveChallenge(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, listingHash)

	if resolveChallErr != nil {
		t.Fatalf("Error resolving challenge: %v", resolveChallErr)
	}

	context.Blockchain.Commit()

	// member 2 should have lost that stake
	member2StakeNow, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthMember2.From)
	if member2StakeNow.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Expected member 2 stake to be 0, got: %v", member2StakeNow)
	}

	// member 1 should be credited that stake now
	member1StakeNow, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthMember1.From)
	if member1StakeNow.Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("Expected member 1 stake to be > 0, got: %v", member1StakeNow)
	}

	// member 1 can actually unstake that credit
	_, unErr := deployed.VotingContract.Unstake(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, listingHash)

	if unErr != nil {
		t.Fatalf("Error unstaking: %v", unErr)
	}

	context.Blockchain.Commit()

	// member 1 has a higher bal now
	member1BalNow, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember1.From)
	if member1BalNow.Cmp(member1Bal) != 1 {
		t.Fatalf("Expected %v to be > %v", member1BalNow, member1Bal)
	}

	// the stake is cleared
	stake, _ := deployed.VotingContract.GetStake(nil, listingHash, context.AuthMember1.From)
	if stake.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Expected stake to be 0, got: %v", stake)
	}
}
