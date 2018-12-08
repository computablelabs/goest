package market

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
	"time"
)

func TestList(t *testing.T) {
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
	// check the market's balance as the mint operation should increment it after successful listing
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.MarketAddress)
	// our listing
	listingHash, _ := deployed.MarketContract.GetListingHash(nil, "FooMarket, AZ.")

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

	// marketToken should be banking the minted amount
	newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.MarketAddress)
	// we should see new bal being > old one
	if newMarketBal.Cmp(marketBal) != 1 {
		t.Fatalf("Expected %v to be > %v", newMarketBal, marketBal)
	}

	// post resolution any candidate should have been removed
	isCandidate, _ := deployed.VotingContract.IsCandidate(nil, listingHash)

	if isCandidate == true {
		t.Fatal("Expected approved listing's candidate to have been removed")
	}
}

func TestResolveApplicationThatFails(t *testing.T) {
	dataHash, _ := deployed.MarketContract.GetListingHash(nil, "BarData")
	// member1 will need funds transferred to them
	_, transErr := deployed.MarketTokenContract.Transfer(&bind.TransactOpts{
		From:     context.AuthFactory.From,
		Signer:   context.AuthFactory.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 1000000,
	}, context.AuthMember1.From, big.NewInt(2000000000000000000)) // 2 token in wei

	if transErr != nil {
		t.Fatalf("Error transferring tokens to member: %v", transErr)
	}

	context.Blockchain.Commit()

	// we'll test the supply functionality in this one, as such we'll need to approve...
	_, approveErr := deployed.MarketTokenContract.Approve(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 1000000,
	}, deployed.MarketAddress, big.NewInt(2000000000000000000)) // up to 2 tokenWei

	if approveErr != nil {
		t.Fatalf("Error approving market contract to spend: %v", approveErr)
	}

	context.Blockchain.Commit()

	// assure balances and permissions
	memberBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember1.From)
	if memberBal.Cmp(big.NewInt(2000000000000000000)) != 0 {
		t.Fatalf("Expected member to have a 2 token balance, got: %v", memberBal)
	}

	// check the market's balance as the list operation, with supply, should increase it
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.MarketAddress)

	_, listErr := deployed.MarketContract.List(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 1000000,
	}, "BarMarket, CA.", dataHash, big.NewInt(1000000000000000000)) // 1 token in wei

	if listErr != nil {
		t.Fatalf("Error applying for list status: %v", listErr)
	}

	context.Blockchain.Commit()

	newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.MarketAddress)

	// list should have transferred the listAmount over
	if newMarketBal.Cmp(marketBal) != 1 {
		t.Fatalf("Expected new market bal, %v, to be > old market bal, %v", newMarketBal, marketBal)
	}

	// our listing
	listingHash, _ := deployed.MarketContract.GetListingHash(nil, "BarMarket, CA.")

	// we should see a listing with a supply in place
	_, _, supply, _, _, _ := deployed.MarketContract.GetListing(nil, listingHash)

	if supply.Cmp(big.NewInt(1000000000000000000)) != 0 {
		t.Fatalf("Expected listing supply to be 1 token in wei, got: %v", supply)
	}

	// move past the voteBy with no votes being cast
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

	// the de-listing should transfer any amount back to owner
	updatedMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.MarketAddress)

	// list should have transferred the listAmount over, going back to the original amt
	if updatedMarketBal.Cmp(marketBal) != 0 {
		t.Fatalf("Expected reverted market bal, %v, to be == old market bal, %v", updatedMarketBal, marketBal)
	}

	// should not be listed now, and no reward
	listed, _, supply, _, rewards, _ := deployed.MarketContract.GetListing(nil, listingHash)

	if listed == true {
		t.Fatalf("Exepected .listed to be false, got: %v", listed)
	}

	if supply.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Exepected supply to be 0, got: %v", supply)
	}

	if rewards.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Exepected rewards to be 0, got: %v", rewards)
	}

	// post resolution any candidate should have been removed
	isCandidate, _ := deployed.VotingContract.IsCandidate(nil, listingHash)

	if isCandidate == true {
		t.Fatal("Expected failed listing's candidate to have been removed")
	}
}

func TestDepositToListing(t *testing.T) {
	// the user's balance should decrease after they deposit
	userBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember1.From)
	// check the market's balance as it should be banking FooMarket's reward
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.MarketAddress)
	// foomarket itself
	listingHash, _ := deployed.MarketContract.GetListingHash(nil, "FooMarket, AZ.")
	// the current state of foo market
	_, _, supply, _, _, _ := deployed.MarketContract.GetListing(nil, listingHash)
	// add some amount
	_, depErr := deployed.MarketContract.DepositToListing(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(2000000000),
		GasLimit: 1000000,
	}, listingHash, big.NewInt(1000000000))

	if depErr != nil {
		t.Fatalf("Error depositing to listing: %v", depErr)
	}

	context.Blockchain.Commit()

	// marketBal should have increased
	newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.MarketAddress)
	if newMarketBal.Cmp(marketBal) != 1 {
		t.Fatalf("Expected %v to be > %v", newMarketBal, marketBal)
	}

	_, _, newSupply, _, _, _ := deployed.MarketContract.GetListing(nil, listingHash)
	if newSupply.Cmp(supply) != 1 {
		t.Fatalf("Expected %v to be > %v", newSupply, supply)
	}

	newUserBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember1.From)
	if newUserBal.Cmp(userBal) != -1 {
		t.Fatalf("Expected %v to be > %v", userBal, newUserBal)
	}
}

func TestWithdrawFromListing(t *testing.T) {
	// the user's balance should increase after a withdrawal
	userBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember1.From)
	// check the market's balance...
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.MarketAddress)
	// foomarket itself
	listingHash, _ := deployed.MarketContract.GetListingHash(nil, "FooMarket, AZ.")
	// the current state of foo market
	_, _, supply, _, _, _ := deployed.MarketContract.GetListing(nil, listingHash)
	// withdraw that same amt
	_, withErr := deployed.MarketContract.WithdrawFromListing(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(2000000000),
		GasLimit: 1000000,
	}, listingHash, big.NewInt(1000000000))

	if withErr != nil {
		t.Fatalf("Error withdrawing from listing: %v", withErr)
	}

	context.Blockchain.Commit()

	// marketBal should have decreased
	newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.MarketAddress)
	if newMarketBal.Cmp(marketBal) != -1 {
		t.Fatalf("Expected %v to be > %v", marketBal, newMarketBal)
	}

	_, _, newSupply, _, _, _ := deployed.MarketContract.GetListing(nil, listingHash)
	if newSupply.Cmp(supply) != -1 {
		t.Fatalf("Expected %v to be > %v", supply, newSupply)
	}

	newUserBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember1.From)
	if newUserBal.Cmp(userBal) != 1 {
		t.Fatalf("Expected %v to be > %v", newUserBal, userBal)
	}
}
