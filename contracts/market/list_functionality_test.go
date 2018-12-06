package market

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
	"time"
)

func TestList(t *testing.T) {
	t.Log("Market should allow a new applicant (no supply)")

	// use the getListingHash func to make a faux dataHash
	dataHash, _ := deployed.MarketContract.GetListingHash(nil, "FooData")

	_, listErr := deployed.MarketContract.List(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 250000,
	}, "FooMarket, AZ.", dataHash, big.NewInt(0))

	if listErr != nil {
		t.Fatalf("Error applying for list status: %v", listErr)
	}

	context.Blockchain.Commit()

	// should have created a voting candidate
	candidates, _ := deployed.VotingContract.GetCandidates(nil)

	if len(candidates) < 1 {
		t.Fatalf("Expected candidates length to be 1 or more, got: %v", len(candidates))
	}
}

func TestIsListing(t *testing.T) {
	// we'll need to gen the listingHash to match
	listingHash, _ := deployed.MarketContract.GetListingHash(nil, "FooMarket, AZ.")

	isListing, _ := deployed.MarketContract.IsListing(nil, listingHash)
	if isListing != true {
		t.Fatalf("Expected isListing to be true, got: %v", isListing)
	}
}

func TestIsListed(t *testing.T) {
	// we'll need to gen the listingHash to match
	listingHash, _ := deployed.MarketContract.GetListingHash(nil, "FooMarket, AZ.")

	isListed, _ := deployed.MarketContract.IsListed(nil, listingHash)
	if isListed == true {
		t.Fatalf("Expected isListed to be false, got: %v", isListed)
	}
}

func TestGetListing(t *testing.T) {
	listingHash, _ := deployed.MarketContract.GetListingHash(nil, "FooMarket, AZ.")

	listed, owner, supply, dataHash, rewards, _ := deployed.MarketContract.GetListing(nil, listingHash)

	if listed != false {
		t.Fatalf("Exepected .listed to be false, got: %v", listed)
	}

	if owner != context.AuthMember1.From {
		t.Fatalf("Exepected owner to be %v, got: %v", context.AuthMember1.From, owner)
	}

	if supply.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Exepected supply to be 0, got: %v", supply)
	}

	// this is what we provided when list called...
	hash, _ := deployed.MarketContract.GetListingHash(nil, "FooData")

	if dataHash != hash {
		t.Fatalf("Exepected data hash to match the provided, got: %v", dataHash)
	}

	if rewards.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Exepected rewards to be 0, got: %v", rewards)
	}
}

func TestGetListings(t *testing.T) {
	listings, _ := deployed.MarketContract.GetListings(nil)

	// we should have at least our one...
	if len(listings) < 1 {
		t.Fatalf("Expected a listings length of at least 1, got: %v", len(listings))
	}
}

func TestIsListingOwner(t *testing.T) {
	isOwner, _ := deployed.MarketContract.IsListingOwner(nil, context.AuthMember1.From)

	if isOwner != true {
		t.Fatalf("Exepected isOwner to be true, got: %v", isOwner)
	}
}

func TestResolveApplication(t *testing.T) {
	// our listing
	listingHash, _ := deployed.MarketContract.GetListingHash(nil, "FooMarket, AZ.")

	// that candidate should be an application
	isApp, _ := deployed.VotingContract.CandidateIs(nil, listingHash, APPLICATION)

	if isApp != true {
		t.Fatalf("Expected candidate is application to be true, got: %v", isApp)
	}

	// make member2 a council member, the owner (factory) can do this...
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

	// verify is a member
	isCouncil, _ := deployed.VotingContract.InCouncil(nil, context.AuthMember2.From)
	if isCouncil != true {
		t.Fatal("Expected member2 to be in council")
	}

	// cast a vote for
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

	// assure the vote was recorded
	voted, _ := deployed.VotingContract.DidVote(nil, listingHash, context.AuthMember2.From)

	if voted != true {
		t.Fatalf("Expected voted to be true, got: %v", voted)
	}

	// move past the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	// poll should be closed now
	closed, _ := deployed.VotingContract.PollClosed(nil, listingHash)

	if closed != true {
		t.Fatalf("Expected pollClosed to be true, got: %v", closed)
	}

	// we should see a truthy didPass call...
	passed, _ := deployed.VotingContract.DidPass(nil, listingHash, big.NewInt(50))

	if passed != true {
		t.Fatalf("Expected didPass to be true, got: %v", passed)
	}

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
	listed, _, supply, _, rewards, _ := deployed.MarketContract.GetListing(nil, listingHash)

	if listed != true {
		t.Fatalf("Exepected .listed to be true, got: %v", listed)
	}

	// supply should still be 0 as owner didn't put any funds in...
	if supply.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Exepected supply to be 0, got: %v", supply)
	}

	// list reward here is 1 token (tokenWei)
	if rewards.Cmp(big.NewInt(1000000000000000000)) != 0 {
		t.Fatalf("Exepected rewards to be 1 (in tokenWei), got: %v", rewards)
	}

	// TODO check markettoken balance before and after to account for mint operation

	// post resolution any candidate should have been removed
	isCandidate, _ := deployed.VotingContract.IsCandidate(nil, listingHash)

	if isCandidate == true {
		t.Fatal("Expected approved listing's candidate to have been removed")
	}
}
