package parameterizertest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
	"time"
)

func TestInvalidSpread(t *testing.T) {
	// Spread must be >= 100
	_, err := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 200000), test.SPREAD, big.NewInt(50))
	test.IfNotNil(t, err, fmt.Sprintf("Error creating proposal: %v", err))

	context.Blockchain.Commit()

	// we should see a candidate now
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(4), big.NewInt(50))
	isCan, _ := deployed.VotingContract.IsCandidate(nil, paramHash)

	if !isCan {
		t.Errorf("Expected isCandidate to be true , got: %v", isCan)
	}

	// cast a vote, member may need funding...
	test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser3.From, big.NewInt(test.ONE_GWEI))

	// member will need to have approved the voting contract to spend
	test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser3, deployed.VotingAddress, big.NewInt(test.ONE_GWEI))

	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), paramHash, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))

	context.Blockchain.Commit()

	_, _, _, _, yea, _, _ := deployed.VotingContract.GetCandidate(nil, paramHash)

	if yea.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected number of votes to be > 0, got: %v", yea)
	}

	// move time forward so the poll is closed
	context.Blockchain.AdjustTime(test.MIN_VOTE_BY * time.Second)
	context.Blockchain.Commit()

	// make sure its closed now
	closed, _ := deployed.VotingContract.PollClosed(nil, paramHash)

	if closed != true {
		t.Errorf("Expected pollClosed to be true, got: %v", closed)
	}

	// resolve it
	_, resolveErr := deployed.ParameterizerContract.ResolveReparam(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), paramHash)
	test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving proposal: %v", resolveErr))
	context.Blockchain.Commit()

	// spread should not be changed now
	newSpread, _ := deployed.ParameterizerContract.GetSpread(nil)

	if newSpread.Cmp(big.NewInt(110)) != 0 {
		t.Errorf("Expected the spread param to still be 110, got: %v", newSpread)
	}

}

func TestInvalidStake(t *testing.T) {
	// Stake must be <= 1/3 of MarketToken supply
	_, err := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
		// Attempt to set the stake to 2 MarketToken, greater than 1/3 of
		// existing market cap of 5 MarketToken
		big.NewInt(test.ONE_GWEI*2), 200000), test.STAKE, big.NewInt(test.ONE_ETH*2))
	test.IfNotNil(t, err, fmt.Sprintf("Error creating proposal: %v", err))

	context.Blockchain.Commit()

	// we should see a candidate now
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(1), big.NewInt(test.ONE_ETH*2))
	isCan, _ := deployed.VotingContract.IsCandidate(nil, paramHash)

	if !isCan {
		t.Errorf("Expected isCandidate to be true, got: %v", isCan)
	}

	// cast a vote, member may need funding...
	test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser3.From, big.NewInt(test.ONE_GWEI))

	// member will need to have approved the voting contract to spend
	test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser3, deployed.VotingAddress, big.NewInt(test.ONE_GWEI))

	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), paramHash, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))

	context.Blockchain.Commit()

	_, _, _, _, yea, _, _ := deployed.VotingContract.GetCandidate(nil, paramHash)

	if yea.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected number of votes to be > 0, got: %v", yea)
	}

	// move time forward so the poll is closed
	context.Blockchain.AdjustTime(test.MIN_VOTE_BY * time.Second)
	context.Blockchain.Commit()

	// make sure its closed now
	closed, _ := deployed.VotingContract.PollClosed(nil, paramHash)

	if closed != true {
		t.Errorf("Expected pollClosed to be true, got: %v", closed)
	}

	// resolve it
	_, resolveErr := deployed.ParameterizerContract.ResolveReparam(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), paramHash)
	test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving proposal: %v", resolveErr))
	context.Blockchain.Commit()

	// stake should not be changed now
	newStake, _ := deployed.ParameterizerContract.GetStake(nil)

	if newStake.Cmp(big.NewInt(test.ONE_GWEI)) != 0 {
		t.Errorf("Expected the stake param to still be one gwei, got: %v", newStake)
	}

}

func TestInvalidVoteByTooLong(t *testing.T) {
	// Vote_by must be <= 1209600. We'll try setting it to 1500000
	_, err := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 200000), test.VOTE_BY, big.NewInt(1500000))
	test.IfNotNil(t, err, fmt.Sprintf("Error creating proposal: %v", err))

	context.Blockchain.Commit()

	// we should see a candidate now
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(1500000))
	isCan, _ := deployed.VotingContract.IsCandidate(nil, paramHash)

	if !isCan {
		t.Errorf("Expected isCandidate to be true, got: %v", isCan)
	}

	// cast a vote, member may need funding...
	test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser3.From, big.NewInt(test.ONE_GWEI))

	// member will need to have approved the voting contract to spend
	test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser3, deployed.VotingAddress, big.NewInt(test.ONE_GWEI))

	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), paramHash, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))

	context.Blockchain.Commit()

	_, _, _, _, yea, _, _ := deployed.VotingContract.GetCandidate(nil, paramHash)

	if yea.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected number of votes to be > 0, got: %v", yea)
	}

	// move time forward so the poll is closed
	context.Blockchain.AdjustTime(test.MIN_VOTE_BY * time.Second)
	context.Blockchain.Commit()

	// make sure its closed now
	closed, _ := deployed.VotingContract.PollClosed(nil, paramHash)

	if closed != true {
		t.Errorf("Expected pollClosed to be true, got: %v", closed)
	}

	// resolve it
	_, resolveErr := deployed.ParameterizerContract.ResolveReparam(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), paramHash)
	test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving proposal: %v", resolveErr))
	context.Blockchain.Commit()

	// vote_by should not be changed now
	newVoteBy, _ := deployed.ParameterizerContract.GetVoteBy(nil)

	if newVoteBy.Cmp(big.NewInt(test.MIN_VOTE_BY)) != 0 {
		t.Errorf("Expected the vote_by param to still be one day, got: %v", newVoteBy)
	}

}

func TestInvalidVoteByTooShort(t *testing.T) {
	// Vote_by must be >= test.MIN_VOTE_BY. We'll try setting it to half that
	_, err := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 200000), test.VOTE_BY, big.NewInt(test.MIN_VOTE_BY/2))
	test.IfNotNil(t, err, fmt.Sprintf("Error creating proposal: %v", err))

	context.Blockchain.Commit()

	// we should see a candidate now
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(test.MIN_VOTE_BY/2))
	isCan, _ := deployed.VotingContract.IsCandidate(nil, paramHash)

	if !isCan {
		t.Errorf("Expected isCandidate to be true, got: %v", isCan)
	}

	// cast a vote, member may need funding...
	test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser3.From, big.NewInt(test.ONE_GWEI))

	// member will need to have approved the voting contract to spend
	test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser3, deployed.VotingAddress, big.NewInt(test.ONE_GWEI))

	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), paramHash, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))

	context.Blockchain.Commit()

	_, _, _, _, yea, _, _ := deployed.VotingContract.GetCandidate(nil, paramHash)

	if yea.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected number of votes to be > 0, got: %v", yea)
	}

	// move time forward so the poll is closed
	context.Blockchain.AdjustTime(test.MIN_VOTE_BY * time.Second)
	context.Blockchain.Commit()

	// make sure its closed now
	closed, _ := deployed.VotingContract.PollClosed(nil, paramHash)

	if closed != true {
		t.Errorf("Expected pollClosed to be true, got: %v", closed)
	}

	// resolve it
	_, resolveErr := deployed.ParameterizerContract.ResolveReparam(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), paramHash)
	test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving proposal: %v", resolveErr))
	context.Blockchain.Commit()

	// vote_by should not be changed now
	newVoteBy, _ := deployed.ParameterizerContract.GetVoteBy(nil)

	if newVoteBy.Cmp(big.NewInt(test.MIN_VOTE_BY)) != 0 {
		t.Errorf("Expected the vote_by param to still be one day, got: %v", newVoteBy)
	}

}

func TestInvalidPlurality(t *testing.T) {
	// Plurality must be <= 100
	_, err := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 200000), test.PLURALITY, big.NewInt(150))
	test.IfNotNil(t, err, fmt.Sprintf("Error creating proposal: %v", err))

	context.Blockchain.Commit()

	// we should see a candidate now
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(6), big.NewInt(150))
	isCan, _ := deployed.VotingContract.IsCandidate(nil, paramHash)

	if !isCan {
		t.Errorf("Expected isCandidate to be true, got: %v", isCan)
	}

	// cast a vote, member may need funding...
	test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser3.From, big.NewInt(test.ONE_GWEI))

	// member will need to have approved the voting contract to spend
	test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser3, deployed.VotingAddress, big.NewInt(test.ONE_GWEI))

	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), paramHash, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))

	context.Blockchain.Commit()

	_, _, _, _, yea, _, _ := deployed.VotingContract.GetCandidate(nil, paramHash)

	if yea.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected number of votes to be > 0, got: %v", yea)
	}

	// move time forward so the poll is closed
	context.Blockchain.AdjustTime(test.MIN_VOTE_BY * time.Second)
	context.Blockchain.Commit()

	// make sure its closed now
	closed, _ := deployed.VotingContract.PollClosed(nil, paramHash)

	if closed != true {
		t.Errorf("Expected pollClosed to be true, got: %v", closed)
	}

	// resolve it
	_, resolveErr := deployed.ParameterizerContract.ResolveReparam(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), paramHash)
	test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving proposal: %v", resolveErr))
	context.Blockchain.Commit()

	// plurality should not be changed now
	newPlurality, _ := deployed.ParameterizerContract.GetPlurality(nil)

	if newPlurality.Cmp(big.NewInt(50)) != 0 {
		t.Errorf("Expected the plurality param to still be 50, got: %v", newPlurality)
	}

}

func TestInvalidMakerPayment(t *testing.T) {
	// maker_payment + backend_payment must be <= 100. backend_payment is 25.
	// Let's try setting maker_payment to 90
	_, err := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 200000), test.MAKER_PAYMENT, big.NewInt(90))
	test.IfNotNil(t, err, fmt.Sprintf("Error creating proposal: %v", err))

	context.Blockchain.Commit()

	// we should see a candidate now
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(9), big.NewInt(90))
	isCan, _ := deployed.VotingContract.IsCandidate(nil, paramHash)

	if !isCan {
		t.Errorf("Expected isCandidate to be true, got: %v", isCan)
	}

	// cast a vote, member may need funding...
	test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser3.From, big.NewInt(test.ONE_GWEI))

	// member will need to have approved the voting contract to spend
	test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser3, deployed.VotingAddress, big.NewInt(test.ONE_GWEI))

	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), paramHash, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))

	context.Blockchain.Commit()

	_, _, _, _, yea, _, _ := deployed.VotingContract.GetCandidate(nil, paramHash)

	if yea.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected number of votes to be > 0, got: %v", yea)
	}

	// move time forward so the poll is closed
	context.Blockchain.AdjustTime(test.MIN_VOTE_BY * time.Second)
	context.Blockchain.Commit()

	// make sure its closed now
	closed, _ := deployed.VotingContract.PollClosed(nil, paramHash)

	if closed != true {
		t.Errorf("Expected pollClosed to be true, got: %v", closed)
	}

	// resolve it
	_, resolveErr := deployed.ParameterizerContract.ResolveReparam(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), paramHash)
	test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving proposal: %v", resolveErr))
	context.Blockchain.Commit()

	// maker_payment should not be changed now
	newMakerPayment, _ := deployed.ParameterizerContract.GetMakerPayment(nil)

	if newMakerPayment.Cmp(big.NewInt(50)) != 0 {
		t.Errorf("Expected the maker payment param to still be 50, got: %v", newMakerPayment)
	}
}

func TestInvalidBackendPayment(t *testing.T) {
	// maker_payment + backend_payment must be <= 100. maker is 50.
	// Let's try setting backend to 60
	_, err := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 200000), test.BACKEND_PAYMENT, big.NewInt(60))
	test.IfNotNil(t, err, fmt.Sprintf("Error creating proposal: %v", err))

	context.Blockchain.Commit()

	// we should see a candidate now
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(8), big.NewInt(60))
	isCan, _ := deployed.VotingContract.IsCandidate(nil, paramHash)

	if !isCan {
		t.Errorf("Expected isCandidate to be true, got: %v", isCan)
	}

	// cast a vote, member may need funding...
	test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser3.From, big.NewInt(test.ONE_GWEI))

	// member will need to have approved the voting contract to spend
	test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser3, deployed.VotingAddress, big.NewInt(test.ONE_GWEI))

	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), paramHash, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))

	context.Blockchain.Commit()

	_, _, _, _, yea, _, _ := deployed.VotingContract.GetCandidate(nil, paramHash)

	if yea.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected number of votes to be > 0, got: %v", yea)
	}

	// move time forward so the poll is closed
	context.Blockchain.AdjustTime(test.MIN_VOTE_BY * time.Second)
	context.Blockchain.Commit()

	// make sure its closed now
	closed, _ := deployed.VotingContract.PollClosed(nil, paramHash)

	if closed != true {
		t.Errorf("Expected pollClosed to be true, got: %v", closed)
	}

	// resolve it
	_, resolveErr := deployed.ParameterizerContract.ResolveReparam(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), paramHash)
	test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving proposal: %v", resolveErr))
	context.Blockchain.Commit()

	// backend_payment should not be changed now
	newBackendPayment, _ := deployed.ParameterizerContract.GetBackendPayment(nil)

	if newBackendPayment.Cmp(big.NewInt(25)) != 0 {
		t.Errorf("Expected the backend payment param to still be 25, got: %v", newBackendPayment)
	}
}
