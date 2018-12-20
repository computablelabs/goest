package market

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
	"time"
)

// the challenge wins in this spec...

func TestChallenge(t *testing.T) {
	// create a applicant
	dataHash, _ := deployed.MarketContract.GetListingHash(nil, "BazData")

	_, listErr := deployed.MarketContract.List(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 250000,
	}, "BarMarket12345", dataHash, big.NewInt(0))

	if listErr != nil {
		t.Fatalf("Error applying for list status: %v", listErr)
	}

	context.Blockchain.Commit()

	// vote for it
	listingHash, _ := deployed.MarketContract.GetListingHash(nil, "BarMarket12345")
	// must be a council member
	isMember, _ := deployed.VotingContract.InCouncil(nil, context.AuthMember2.From)
	if isMember != true {
		_, councilErr := deployed.VotingContract.AddToCouncil(&bind.TransactOpts{
			From:     context.AuthFactory.From,
			Signer:   context.AuthFactory.Signer,
			GasPrice: big.NewInt(2000000000), // 2 Gwei
			GasLimit: 100000,
		}, context.AuthMember2.From)

		if councilErr != nil {
			t.Fatal("Error adding member to council")
		}

		context.Blockchain.Commit()
	}

	_, voteErr := deployed.VotingContract.Vote(&bind.TransactOpts{
		From:     context.AuthMember2.From,
		Signer:   context.AuthMember2.Signer,
		GasPrice: big.NewInt(2000000000),
		GasLimit: 100000,
	}, listingHash)

	if voteErr != nil {
		t.Fatalf("Error voting for candidate: %v", voteErr)
	}

	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	// any council member can call for resolution
	_, resolveErr := deployed.MarketContract.ResolveApplication(&bind.TransactOpts{
		From:     context.AuthMember2.From,
		Signer:   context.AuthMember2.Signer,
		GasPrice: big.NewInt(2000000000),
		GasLimit: 1000000,
	}, listingHash)

	if resolveErr != nil {
		t.Fatalf("Error resolving application: %v", resolveErr)
	}

	context.Blockchain.Commit()

	// should be listed now, with a reward
	listed, _, _, _, _, _ := deployed.MarketContract.GetListing(nil, listingHash)

	if listed != true {
		t.Fatalf("Exepected .listed to be true, got: %v", listed)
	}

	// now we can challenge. member2 will need funds approved && such
	_, transErr := deployed.MarketTokenContract.Transfer(&bind.TransactOpts{
		From:     context.AuthFactory.From,
		Signer:   context.AuthFactory.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 1000000,
	}, context.AuthMember2.From, big.NewInt(2000000000000000000)) // 2 token in wei

	if transErr != nil {
		t.Fatalf("Error transferring tokens to member: %v", transErr)
	}

	context.Blockchain.Commit()

	_, approveErr := deployed.MarketTokenContract.Approve(&bind.TransactOpts{
		From:     context.AuthMember2.From,
		Signer:   context.AuthMember2.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 1000000,
	}, deployed.MarketAddress, big.NewInt(2000000000000000000)) // up to 2 tokenWei

	if approveErr != nil {
		t.Fatalf("Error approving market contract to spend: %v", approveErr)
	}

	context.Blockchain.Commit()

	// note our balances right now as challenging will change them
	memberBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember2.From)
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.MarketAddress)

	_, challengeErr := deployed.MarketContract.Challenge(&bind.TransactOpts{
		From:     context.AuthMember2.From,
		Signer:   context.AuthMember2.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
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
	newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.MarketAddress)
	if newMarketBal.Cmp(marketBal) != 1 {
		t.Fatalf("Expected %v to be > %v", newMarketBal, marketBal)
	}
}

func TestGetChallenge(t *testing.T) {
	listingHash, _ := deployed.MarketContract.GetListingHash(nil, "BarMarket12345")

	challenger, fromSupply, fromRewards, _ := deployed.MarketContract.GetChallenge(nil, listingHash)

	if challenger != context.AuthMember2.From {
		t.Fatalf("Expected challenger to be %v, got: %v", context.AuthMember2.From, challenger)
	}

	if fromSupply.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Expected amount to be 0, got: %v", fromSupply)
	}

	if fromRewards.Cmp(big.NewInt(1000000000000000000)) != 0 {
		t.Fatalf("Expected amount to be 1 tokenWei, got: %v", fromRewards)
	}

	// should be a candidate for this challenge as well
	kind, _, _, _ := deployed.VotingContract.GetCandidate(nil, listingHash)

	if kind != 2 {
		t.Fatalf("Expected kind to be challenge (2), got: %v", kind)
	}
}

func TestResolveChallenge(t *testing.T) {
	listingHash, _ := deployed.MarketContract.GetListingHash(nil, "BarMarket12345")
	// the marketBal will decrease as the challenger gets their stake + listings's stake
	// marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.MarketAddress)
	// memberBal, _ deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember2)

	// we need to cast a vote for the challenge, or the listing will "win"
	_, voteErr := deployed.VotingContract.Vote(&bind.TransactOpts{
		From:     context.AuthMember2.From,
		Signer:   context.AuthMember2.Signer,
		GasPrice: big.NewInt(2000000000),
		GasLimit: 1000000,
	}, listingHash)

	if voteErr != nil {
		t.Fatalf("Error voting for candidate: %v", voteErr)
	}

	context.Blockchain.Commit()

	// member2's vote should be recorded now
	voted, _ := deployed.VotingContract.DidVote(nil, listingHash, context.AuthMember2.From)

	if voted != true {
		t.Fatalf("Expected voted to be true, got: %v", voted)
	}

	_, _, votes, _ := deployed.VotingContract.GetCandidate(nil, listingHash)

	if votes.Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("Expected at least one vote, got: %v", votes)
	}

	// the vote is recorded now, and the challenge should be ready to be resolved once we pass the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	// either member could call for this as both are council members here, NOTE this will change in future iteration
	_, resolveErr := deployed.MarketContract.ResolveChallenge(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 1000000,
	}, listingHash)

	if resolveErr != nil {
		t.Fatalf("Error resolving challenge: %v", resolveErr)
	}

	context.Blockchain.Commit()

	// should no longer be a listing
	listed, _ := deployed.MarketContract.IsListing(nil, listingHash)

	if listed == true {
		t.Fatalf("Expected .listed to be false, got")
	}
}
