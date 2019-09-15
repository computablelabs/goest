package votingtest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
	"time"
)

func TestVote(t *testing.T) {
	// auth member will need at least the stake
	transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser1.From,
		big.NewInt(2*test.ONE_GWEI))
	test.IfNotNil(t, transErr, "Error maybe transferring market tokens")

	// member will need to have approved the voting contract to spend at least the stake
	incErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser1, deployed.VotingAddress,
		big.NewInt(2*test.ONE_GWEI))
	test.IfNotNil(t, incErr, "Error maybe transferring market token approval")

	// Get the hash
	reparamHash, _ := deployed.ParameterizerContract.GetHash(nil, test.VOTE_BY, big.NewInt(100000))

	// Check that there is not a candidate now
	isCan, _ := deployed.VotingContract.IsCandidate(nil, reparamHash)
	if isCan {
		t.Errorf("Expected isCandidate to be false before reparam submission, got: %v", isCan)
	}

	// we need a candidate, a privileged contract must add it...
	_, candErr := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), test.VOTE_BY, big.NewInt(100000))
	test.IfNotNil(t, candErr, fmt.Sprintf("Error adding candidate: %v", candErr))
	context.Blockchain.Commit()

	// Check that there is a candidate now
	isCanNow, _ := deployed.VotingContract.IsCandidate(nil, reparamHash)
	if !isCanNow {
		t.Errorf("Expected isCandidate to be true, got: %v", isCan)
	}
	// should be no votes atm
	_, _, _, _, preYea, _, _ := deployed.VotingContract.GetCandidate(nil, reparamHash)

	if preYea.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected number of votes to be 0, got: %v", preYea)
	}

	// cast a yea vote
	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), reparamHash, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))

	context.Blockchain.Commit()

	// has been recorded in the candidate's votes
	_, _, _, _, yea, _, _ := deployed.VotingContract.GetCandidate(nil, reparamHash)

	if yea.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected number of votes to be > 0, got: %v", yea)
	}

	// should see the stake bal at voting.stakes.address.hash
	stake, _ := deployed.VotingContract.GetStake(nil, reparamHash, context.AuthUser1.From)

	if stake.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected stake to be > 0, got: %v", stake)
	}
}

func TestPollClosed(t *testing.T) {
	// should still be open
	reparamHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(100000))
	closed, _ := deployed.VotingContract.PollClosed(nil, reparamHash)

	if closed != false {
		t.Errorf("Expected pollClosed to be false, got: %v", closed)
	}

	// now move time forward so that voteBy has elapsed
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	updated, _ := deployed.VotingContract.PollClosed(nil, reparamHash)

	if updated != true {
		t.Errorf("Expected pollClosed to be true, got: %v", updated)
	}
}

func TestDidPass(t *testing.T) {
	reparamHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(100000))
	// the one total vote will do it
	passed, _ := deployed.VotingContract.DidPass(nil, reparamHash, big.NewInt(50))

	if passed != true {
		t.Errorf("Expected didPass to be true, got: %v", passed)
	}
}

// test that the voting stake remains after reparam is resolved...
func TestResolveReparam(t *testing.T) {
	bytes, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(100000))

	// does not matter who calls for the resolution
	_, err := deployed.ParameterizerContract.ResolveReparam(test.GetTxOpts(context.AuthUser2, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), bytes)
	test.IfNotNil(t, err, fmt.Sprintf("Error resolving reparam: %v", err))

	context.Blockchain.Commit()

	// candidate is gone, param has changed, but stake remains
	isCandidate, _ := deployed.VotingContract.IsCandidate(nil, bytes)

	if isCandidate != false {
		t.Errorf("Expected isCandidate be false, got: %v", isCandidate)
	}

	voteBy, _ := deployed.ParameterizerContract.GetVoteBy(nil)
	if voteBy.Cmp(big.NewInt(100000)) != 0 {
		t.Errorf("Expected voteBy to be 100000, got: %v", voteBy)
	}

	stake, _ := deployed.VotingContract.GetStake(nil, bytes, context.AuthUser1.From)

	if stake.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected stake to be > 0, got: %v", stake)
	}
}
