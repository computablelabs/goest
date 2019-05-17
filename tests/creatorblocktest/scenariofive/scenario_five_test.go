package scenariofive

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
	"time"
)

func oneHundredOneEth() *big.Int {
	one := big.NewInt(test.ONE_WEI)
	oneHundred := one.Mul(one, big.NewInt(100))
	return oneHundred.Add(oneHundred, big.NewInt(test.ONE_WEI))
}

func oneHundredEth() *big.Int {
	one := big.NewInt(test.ONE_WEI)
	return one.Mul(one, big.NewInt(100))
}

func TestFullSimulation(t *testing.T) {
	// The creator of the market starts by transfering 100 ETH to reserve
	_, transError := deployed.EtherTokenContract.Transfer(test.GetTxOpts(
		context.AuthOwner, nil, big.NewInt(test.ONE_GWEI*2), 100000),
		deployed.InvestingAddress, oneHundredEth())
	test.IfNotNil(t, transError, "Error transferring token")
	context.Blockchain.Commit()

	ownerEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthOwner.From)
	resEthBal, _ := deployed.EtherTokenContract.BalanceOf(nil, deployed.InvestingAddress)

	// has 1 ETH left
	if ownerEthBal.Cmp(big.NewInt(test.ONE_WEI)) != 0 {
		t.Errorf("Expected ether token balance of 1 ETH, got: %v", ownerEthBal)
	}

	// got the 100
	if resEthBal.Cmp(oneHundredEth()) != 0 {
		t.Errorf("Expected reserve of 100 Eth, got: %v", resEthBal)
	}

	t.Logf("Current Reserve balance: %v", test.Commafy(resEthBal))
	// Original market token supply
	mtSup, _ := deployed.MarketTokenContract.TotalSupply(nil)
	t.Logf("Original Market token total supply post creation: %v", test.Commafy(mtSup))

	t.Logf("Number Makers: %d", len(makers))
	t.Logf("Number Investors: %d", len(investors))

	// Let's have the maker submit a listing
	for name, maker := range makers {
		t.Logf("Submitting listing for %s", name)
		listingHash := test.GenBytes32(name)

		_, listErr := deployed.ListingContract.List(test.GetTxOpts(maker, nil,
			big.NewInt(test.ONE_GWEI*2), 250000), listingHash)
		test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
		context.Blockchain.Commit()

		// Check that the submitted listing is now a candidate
		isCandidate, _ := deployed.VotingContract.IsCandidate(nil, listingHash)

		if isCandidate == false {
			t.Error("Submitted listing is not listed as a candidate")
		}

		// have the datatrust submit a data_hash for this listing.
		dataHash, _ := deployed.DatatrustContract.GetHash(nil, name)
		_, dataErr := deployed.DatatrustContract.SetDataHash(test.GetTxOpts(context.AuthBackend, nil,
			big.NewInt(test.ONE_GWEI*2), 100000), listingHash, dataHash)
		test.IfNotNil(t, dataErr, "Error setting data hash for listing")
		context.Blockchain.Commit()

		// check the market's balance as the mint operation should increment it after successful listing
		marketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)

		// cast a vote for, voter may need funds...
		// TODO(rbharath): Can this line perhaps be removed? Since we start accounts off with generous balances?
		transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser2.From,
			big.NewInt(10000000000000000)) // 1 X 10**16
		test.IfNotNil(t, transErr, fmt.Sprintf("Error transferring tokens to member: %v", transErr))
		context.Blockchain.Commit()

		// member will need to have approved the voting contract to spend
		appErr := test.MaybeIncreaseMarketTokenApproval(context, deployed, context.AuthUser2, deployed.VotingAddress,
			big.NewInt(10000000000000000)) // 1 X 10**16
		test.IfNotNil(t, appErr, fmt.Sprintf("Error approving market contract to spend: %v", appErr))
		context.Blockchain.Commit()

		// TODO: Should we increase the amount the contract votes with?
		_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser2, nil,
			big.NewInt(test.ONE_GWEI*2), 150000), listingHash, big.NewInt(1))
		test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))
		context.Blockchain.Commit()

		// move past the voteBy
		context.Blockchain.AdjustTime(100 * time.Second)
		context.Blockchain.Commit()

		// any stakeholder can call for resolution
		_, resolveErr := deployed.ListingContract.ResolveApplication(test.GetTxOpts(context.AuthUser2, nil,
			big.NewInt(test.ONE_GWEI*2), 1000000), listingHash)
		test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving application: %v", resolveErr))
		context.Blockchain.Commit()

		// Check that the submitted listing is no longer a candidate
		isCandidate, _ = deployed.VotingContract.IsCandidate(nil, listingHash)
		if isCandidate == true {
			t.Error("Expected approved listing's candidate to have been removed")
		}

		// should be listed now, with a reward
		owner, supply, _ := deployed.ListingContract.GetListing(nil, listingHash)

		if owner != maker.From {
			t.Errorf("Expected owner to be %v, got: %v", maker.From, owner)
		}
		t.Logf("Listing supply: %v", test.Commafy(supply))

		// makerMarketBal should be 0 since the balance has not been withdrawn yet
		makerMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, maker.From)
		t.Logf("%s MarketToken balance: %v", name, test.Commafy(makerMarketBal))

		// supply should reflect the list reward
		if supply.Cmp(big.NewInt(250000000000000)) != 0 {
			t.Errorf("Expected supply to be list reward, got: %v", supply)
		}

		// marketToken should be banking the minted amount
		newMarketBal, _ := deployed.MarketTokenContract.BalanceOf(nil, deployed.ListingAddress)

		// we should see new bal being > old one
		if newMarketBal.Cmp(marketBal) != 1 {
			t.Errorf("Expected %v to be > %v", newMarketBal, marketBal)
		}

		// Market Token supply post listing
		mtSup, _ := deployed.MarketTokenContract.TotalSupply(nil)
		t.Logf("Original Market token total supply post %s listing: %v", name, test.Commafy(mtSup))

	}
}
