package scenariothree

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

func TestInitialBalance(t *testing.T) {
	// owner has all the eth atm
	etBal, _ := deployed.EtherTokenContract.BalanceOf(nil, context.AuthOwner.From)
	// all market tokens that exist atm
	mtSup, _ := deployed.MarketTokenContract.TotalSupply(nil)
	if etBal.Cmp(oneHundredOneEth()) != 0 {
		t.Errorf("Expected ether token balance of %v, got: %v", oneHundredOneEth(), etBal)
	}
	if mtSup.Cmp(big.NewInt(test.ONE_WEI)) != 0 {
		t.Errorf("Expected market token supply of 1wei, got: %v", mtSup)
	}

	t.Logf("Current Market Token total supply: %v", test.Commafy(mtSup))
}

func TestTransferToReserveThenMakersMake(t *testing.T) {
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
	t.Logf("Number Makers: %d", len(makers))

	// Let's have the maker submit a listing
	for name, maker := range makers {
		t.Logf("Submitting listing for %s", fmt.Sprintf("%s FooMarket, AZ.", name))
		listingHash := test.GenBytes32(fmt.Sprintf("%s FooMarket, AZ.", name))
		t.Logf("listingHash: %v", listingHash)

		_, listErr := deployed.ListingContract.List(test.GetTxOpts(maker, nil,
			big.NewInt(test.ONE_GWEI*2), 250000), listingHash)
		test.IfNotNil(t, listErr, fmt.Sprintf("Error applying for list status: %v", listErr))
		context.Blockchain.Commit()

		// Check that the submitted listing is now a candidate
		isCan, _ := deployed.VotingContract.IsCandidate(nil, listingHash)

		if !isCan {
			t.Errorf("Expected isCandidate to be true, got: %v", isCan)
		}

		// have the datatrust submit a data_hash for this listing.
		dataHash, _ := deployed.DatatrustContract.GetHash(nil,
			fmt.Sprintf("%s: thisissomedata", name))
		_, dataErr := deployed.DatatrustContract.SetDataHash(test.GetTxOpts(context.AuthBackend, nil,
			big.NewInt(test.ONE_GWEI*2), 100000), listingHash, dataHash)
		test.IfNotNil(t, dataErr, "Error setting data hash for listing")
		context.Blockchain.Commit()

		// cast a vote for, voter may need funds...
		transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser2.From,
			big.NewInt(test.ONE_GWEI))
		test.IfNotNil(t, transErr, fmt.Sprintf("Error transferring tokens to member: %v", transErr))
		context.Blockchain.Commit()

		// member will need to have approved the voting contract to spend
		appErr := test.MaybeIncreaseMarketTokenApproval(context, deployed, context.AuthUser2, deployed.VotingAddress,
			big.NewInt(test.ONE_GWEI))
		test.IfNotNil(t, appErr, fmt.Sprintf("Error approving market contract to spend: %v", appErr))
		context.Blockchain.Commit()

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

		// should be listed now, with a reward
		owner, supply, _ := deployed.ListingContract.GetListing(nil, listingHash)

		if owner != maker.From {
			t.Errorf("Expected owner to be %v, got: %v", maker.From, owner)
		}

		// supply should reflect the list reward
		if supply.Cmp(big.NewInt(250000000000000)) != 0 {
			t.Errorf("Expected supply to be list reward, got: %v", supply)
		}
	}
}
