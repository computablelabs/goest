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

	owner, supply, _ := deployed.ListingContract.GetListing(nil, listingHash)

	// should not be owned yet
	if owner == context.AuthMember1.From {
		t.Fatalf("Exepected owner to be %v, got: %v", context.AuthMember1.From, owner)
	}

	if supply.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Exepected supply to be 0, got: %v", supply)
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

	// cast a vote for, voter may need funds...
	memBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember2.From)
	if memBal.Cmp(big.NewInt(ONE_GWEI)) == -1 {
		// transfer one
		_, transErr := deployed.MarketTokenContract.Transfer(&bind.TransactOpts{
			From:     context.AuthFactory.From,
			Signer:   context.AuthFactory.Signer,
			GasPrice: big.NewInt(ONE_GWEI * 2),
			GasLimit: 1000000,
		}, context.AuthMember2.From, big.NewInt(ONE_GWEI))

		if transErr != nil {
			t.Fatalf("Error transferring tokens to member: %v", transErr)
		}

		context.Blockchain.Commit()
	}

	// member will need to have approved the voting contract to spend
	allowed, _ := deployed.MarketTokenContract.Allowance(nil, context.AuthMember2.From, deployed.VotingAddress)
	if !(allowed.Cmp(big.NewInt(ONE_GWEI)) >= 0) {
		_, approveErr := deployed.MarketTokenContract.Approve(&bind.TransactOpts{
			From:     context.AuthMember2.From,
			Signer:   context.AuthMember2.Signer,
			GasPrice: big.NewInt(ONE_GWEI * 2),
			GasLimit: 1000000,
		}, deployed.VotingAddress, big.NewInt(ONE_GWEI))

		if approveErr != nil {
			t.Fatalf("Error approving market contract to spend: %v", approveErr)
		}

		context.Blockchain.Commit()
	}

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

	// should be listed now, with a reward having gone to listing supply
	owner, supply, _ := deployed.ListingContract.GetListing(nil, listingHash)

	if owner != context.AuthMember1.From {
		t.Fatalf("Expected owner to be %v, got: %v", context.AuthMember1.From, owner)
	}

	if supply.Cmp(big.NewInt(ONE_WEI)) != 0 {
		t.Fatalf("Exepected supply to be one token in wei, got: %v", supply)
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

	// member can unstake
	_, unErr := deployed.VotingContract.Unstake(&bind.TransactOpts{
		From:     context.AuthMember2.From,
		Signer:   context.AuthMember2.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 150000,
	}, listingHash)

	if unErr != nil {
		t.Fatalf("Error Unstaking: %v", unErr)
	}

	context.Blockchain.Commit()
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
	owner, supply, _ := deployed.ListingContract.GetListing(nil, listingHash)

	if owner == context.AuthMember1.From {
		t.Fatalf("Exepected owner to be zero address, got: %v", owner)
	}

	if supply.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Exepected supply to be 0, got: %v", supply)
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

func TestWithdrawFromListing(t *testing.T) {
	// the user's balance should increase after a withdrawal
	userBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember1.From)
	// check the market's balance...
	marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)
	// foomarket itself
	listingHash, _ := deployed.ListingContract.GetHash(nil, "FooMarket, AZ.")
	// the current state of foo market
	_, supply, _ := deployed.ListingContract.GetListing(nil, listingHash)
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

	_, newSupply, _ := deployed.ListingContract.GetListing(nil, listingHash)
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
