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
		GasLimit: 100000,
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

	_, approveErr := deployed.MarketTokenContract.Approve(&bind.TransactOpts{
		From:     context.AuthMember2.From,
		Signer:   context.AuthMember2.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, deployed.ListingAddress, big.NewInt(ONE_WEI*2)) // up to 2 tokenWei

	if approveErr != nil {
		t.Fatalf("Error approving market contract to spend: %v", approveErr)
	}

	context.Blockchain.Commit()

	// note our balances right now as challenging will change them
	memberBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember2.From)
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)

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
	newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)
	if newMarketBal.Cmp(marketBal) != 1 {
		t.Fatalf("Expected %v to be > %v", newMarketBal, marketBal)
	}
}

func TestResolveChallenge(t *testing.T) {
	listingHash, _ := deployed.ListingContract.GetHash(nil, "BarMarket12345")
	// the marketBal will decrease as the challenger gets their stake + listings's stake
	// marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.MarketAddress)
	// memberBal, _ deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember2)

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

	_, _, _, yea, _, _ := deployed.VotingContract.GetCandidate(nil, listingHash)

	if yea.Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("Expected at least one vote, got: %v", yea)
	}

	// the vote is recorded now, and the challenge should be ready to be resolved once we pass the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	// either member could call for this as both are council members here, NOTE this will change in future iteration
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
}
