package listing

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
	"time"
)

func TestList(t *testing.T) {
	_, listErr := deployed.ListingContract.List(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 250000,
	}, "FooMarket, AZ.")

	if listErr != nil {
		t.Fatalf("Error applying for list status: %v", listErr)
	}

	context.Blockchain.Commit()

	// should have created a voting candidate
	listingHash, _ := deployed.ListingContract.GetHash(nil, "FooMarket, AZ.")
	isCan, _ := deployed.VotingContract.IsCandidate(nil, listingHash)

	if !isCan {
		t.Fatalf("Expected isCandidate to be true, got: %v", isCan)
	}
}

func TestIsListed(t *testing.T) {
	// we'll need to gen the listingHash to match
	listingHash, _ := deployed.ListingContract.GetHash(nil, "FooMarket, AZ.")

	isListed, _ := deployed.ListingContract.IsListed(nil, listingHash)
	if isListed == true {
		t.Fatalf("Expected isListed to be false, got: %v", isListed)
	}
}

func TestGetListing(t *testing.T) {
	listingHash, _ := deployed.ListingContract.GetHash(nil, "FooMarket, AZ.")

	owner, supply, rewards, _ := deployed.ListingContract.GetListing(nil, listingHash)

	// should not be owned yet
	if owner == context.AuthMember1.From {
		t.Fatalf("Exepected owner to be %v, got: %v", context.AuthMember1.From, owner)
	}

	if supply.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Exepected supply to be 0, got: %v", supply)
	}

	if rewards.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Exepected rewards to be 0, got: %v", rewards)
	}
}

func TestResolveApplication(t *testing.T) {
	// our listing
	listingHash, _ := deployed.ListingContract.GetHash(nil, "FooMarket, AZ.")
	// the datatrust must have a data_hash for this listing before it will pass
	dataHash, _ := deployed.DatatrustContract.GetHash(nil, "thisissomedata")
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
	// check the market's balance as the mint operation should increment it after successful listing
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)

	// cast a vote for
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

	// any council member can call for resolution
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
	owner, supply, rewards, _ := deployed.ListingContract.GetListing(nil, listingHash)

	if owner != context.AuthMember1.From {
		t.Fatalf("Expected owner to be %v, got: %v", context.AuthMember1.From, owner)
	}

	// supply should still be 0 as owner didn't put any funds in...
	if supply.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Exepected supply to be 0, got: %v", supply)
	}

	// list reward here is 1 token (tokenWei)
	if rewards.Cmp(big.NewInt(ONE_WEI)) != 0 {
		t.Fatalf("Exepected rewards to be 1 (in tokenWei), got: %v", rewards)
	}

	// marketToken should be banking the minted amount
	newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)
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
	// check the market's balance as it should not change in a no-list op
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)

	_, listErr := deployed.ListingContract.List(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, "BarMarket, CA.")

	if listErr != nil {
		t.Fatalf("Error applying for list status: %v", listErr)
	}

	context.Blockchain.Commit()

	// our listing
	listingHash, _ := deployed.ListingContract.GetHash(nil, "BarMarket, CA.")

	// move past the voteBy with no votes being cast
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	// any council member can call for resolution
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

	// should not be listed now, and no reward
	owner, supply, rewards, _ := deployed.ListingContract.GetListing(nil, listingHash)

	if owner == context.AuthMember1.From {
		t.Fatalf("Exepected owner to be zero address, got: %v", owner)
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

	newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)
	if marketBal.Cmp(newMarketBal) != 0 {
		t.Fatalf("Expected %v to be %v", marketBal, newMarketBal)
	}
}

func TestDepositToListing(t *testing.T) {
	// user 1 needs funds and approvals
	_, transErr := deployed.MarketTokenContract.Transfer(&bind.TransactOpts{
		From:     context.AuthFactory.From,
		Signer:   context.AuthFactory.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, context.AuthMember1.From, big.NewInt(ONE_WEI*2))

	if transErr != nil {
		t.Fatalf("Error transferring tokens, %v", transErr)
	}

	context.Blockchain.Commit()

	_, approveErr := deployed.MarketTokenContract.Approve(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, deployed.ListingAddress, big.NewInt(ONE_WEI*2))

	if approveErr != nil {
		t.Fatalf("Error approving, %v", approveErr)
	}

	context.Blockchain.Commit()

	// the user's balance should decrease after they deposit
	userBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember1.From)
	// check the market's balance as it should be banking FooMarket's reward
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)
	// foomarket itself
	listingHash, _ := deployed.ListingContract.GetHash(nil, "FooMarket, AZ.")
	// the current state of foo market
	_, supply, _, _ := deployed.ListingContract.GetListing(nil, listingHash)
	// add some amount
	_, depErr := deployed.ListingContract.DepositToListing(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, listingHash, big.NewInt(ONE_WEI*0.01)) // some random small amount

	if depErr != nil {
		t.Fatalf("Error depositing to listing: %v", depErr)
	}

	context.Blockchain.Commit()

	// marketBal should have increased
	newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)
	if newMarketBal.Cmp(marketBal) != 1 {
		t.Fatalf("Expected new market balance %v to be > %v", newMarketBal, marketBal)
	}

	_, newSupply, _, _ := deployed.ListingContract.GetListing(nil, listingHash)
	if newSupply.Cmp(supply) != 1 {
		t.Fatalf("Expected new supply %v to be > %v", newSupply, supply)
	}

	newUserBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember1.From)
	if newUserBal.Cmp(userBal) != -1 {
		t.Fatalf("Expected new user balance %v to be > %v", userBal, newUserBal)
	}
}

func TestWithdrawFromListing(t *testing.T) {
	// the user's balance should increase after a withdrawal
	userBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember1.From)
	// check the market's balance...
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)
	// foomarket itself
	listingHash, _ := deployed.ListingContract.GetHash(nil, "FooMarket, AZ.")
	// the current state of foo market
	_, _, supply, _ := deployed.ListingContract.GetListing(nil, listingHash)
	// withdraw that same amt
	_, withErr := deployed.ListingContract.WithdrawFromListing(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, listingHash, big.NewInt(ONE_WEI*0.01))

	if withErr != nil {
		t.Fatalf("Error withdrawing from listing: %v", withErr)
	}

	context.Blockchain.Commit()

	// marketBal should have decreased
	newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)
	if newMarketBal.Cmp(marketBal) != -1 {
		t.Fatalf("Expected %v to be > %v", marketBal, newMarketBal)
	}

	_, newSupply, _, _ := deployed.ListingContract.GetListing(nil, listingHash)
	if newSupply.Cmp(supply) != -1 {
		t.Fatalf("Expected %v to be > %v", supply, newSupply)
	}

	newUserBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember1.From)
	if newUserBal.Cmp(userBal) != 1 {
		t.Fatalf("Expected %v to be > %v", newUserBal, userBal)
	}
}

func TestExit(t *testing.T) {
	// going to remove this one
	listingHash, _ := deployed.ListingContract.GetHash(nil, "FooMarket, AZ.")
	// is it listed?
	listed, _ := deployed.ListingContract.IsListed(nil, listingHash)
	if !listed {
		t.Fatalf("Expected %v to be true", listed)
	}
	// we know that the market's bank will decrease here by the listReward
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)

	// there is a data_hash for this listing
	dataHash, _ := deployed.DatatrustContract.GetDataHash(nil, listingHash)

	_, exitErr := deployed.ListingContract.Exit(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, listingHash)

	if exitErr != nil {
		t.Fatalf("Error exiting listing from market, %v", exitErr)
	}

	context.Blockchain.Commit()

	listedNow, _ := deployed.ListingContract.IsListed(nil, listingHash)
	if listedNow {
		t.Fatalf("Expected %v to be false", listedNow)
	}

	updatedMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)
	if updatedMarketBal.Cmp(marketBal) != -1 {
		t.Fatalf("Expected %v to be > %v", marketBal, updatedMarketBal)
	}

	// data_hash should be cleared
	dataHashNow, _ := deployed.DatatrustContract.GetDataHash(nil, listingHash)
	if dataHashNow == dataHash {
		t.Fatalf("Expected data hash to be empty, got: %v", dataHashNow)
	}
}

func TestConvertListing(t *testing.T) {
	_, listErr := deployed.ListingContract.List(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 250000,
	}, "SpamMarket, AZ.")

	if listErr != nil {
		t.Fatalf("Error applying for list status: %v", listErr)
	}

	context.Blockchain.Commit()

	listingHash, _ := deployed.ListingContract.GetHash(nil, "SpamMarket, AZ.")

	// data_trust must have the data_hash or the listing will fail
	dataHash, _ := deployed.DatatrustContract.GetHash(nil, "all-the-data!")
	_, dataErr := deployed.DatatrustContract.SetDataHash(&bind.TransactOpts{
		From:     context.AuthBackend.From,
		Signer:   context.AuthBackend.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 100000,
	}, listingHash, dataHash)

	if dataErr != nil {
		t.Fatal("Error setting data hash for listing")
	}

	// cast a vote for (we know member2 is a council member at this point)
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

	// any council member can call for resolution
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
	owner, _, rewards, _ := deployed.ListingContract.GetListing(nil, listingHash)

	if owner != context.AuthMember1.From {
		t.Fatalf("Exepected %v to be %v", context.AuthMember1.From, owner)
	}

	// list reward here is 1 token (tokenWei)
	if rewards.Cmp(big.NewInt(ONE_WEI)) != 0 {
		t.Fatalf("Exepected rewards to be 1 (in tokenWei), got: %v", rewards)
	}

	// get the owner's market token bal as it should increase on conversion
	userBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember1.From)
	// same for the market, but should go down
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)

	_, convertErr := deployed.ListingContract.ConvertListing(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, listingHash)

	if convertErr != nil {
		t.Fatalf("Error converting listing: %v", convertErr)
	}

	context.Blockchain.Commit()

	reward, _ := deployed.ParameterizerContract.GetListReward(nil)
	// user's market bal should have gone up by the reward now
	newUserBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember1.From)
	expectedUserBal := userBal.Add(userBal, reward)
	// same for the market, but should go down
	newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)
	expectedMarketBal := marketBal.Sub(marketBal, reward)

	if newUserBal.Cmp(expectedUserBal) != 0 {
		t.Fatalf("expected user market token balance of %v, got: %v", expectedUserBal, newUserBal)
	}

	if newMarketBal.Cmp(expectedMarketBal) != 0 {
		t.Fatalf("expected market token balance of %v, got: %v", expectedMarketBal, newMarketBal)
	}

	// data hash should still be intact
	dataHashNow, _ := deployed.DatatrustContract.GetDataHash(nil, listingHash)
	if dataHashNow != dataHash {
		t.Fatalf("Expected data hash not to be empty, got: %v", dataHashNow)
	}
}
