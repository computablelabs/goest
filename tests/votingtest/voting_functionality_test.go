package votingtest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
	"time"
)

func TestVote(t *testing.T) {
	// auth member will need at least the stake * 2 (if we use them for reparam as well)
	transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser1.From,
		big.NewInt(test.ONE_GWEI*2))
	test.IfNotNil(t, transErr, "Error maybe transferring market tokens")

	// member will need to have approved the voting contract to spend at least the stake
	incErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser1, deployed.VotingAddress,
		big.NewInt(test.ONE_GWEI*2))
	test.IfNotNil(t, incErr, "Error maybe transferring market token approval")

	// we need a candidate, a privileged contract must add it...(NOTE: reparams are staked now)
	_, candErr := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 500000), test.VOTE_BY, big.NewInt(2*test.MIN_VOTE_BY))
	test.IfNotNil(t, candErr, fmt.Sprintf("Error adding candidate: %v", candErr))

	context.Blockchain.Commit()

	bytes, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(2*test.MIN_VOTE_BY))

	// should be no votes atm
	_, _, _, _, preYea, _, _ := deployed.VotingContract.GetCandidate(nil, bytes)

	if preYea.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected number of votes to be 0, got: %v", preYea)
	}

	// cast a yea vote
	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 500000), bytes, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))

	context.Blockchain.Commit()

	// has been recorded in the candidate's votes
	_, _, _, _, yea, _, _ := deployed.VotingContract.GetCandidate(nil, bytes)

	if yea.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected number of votes to be > 0, got: %v", yea)
	}

	// should see the stake bal at voting.stakes.address.hash
	stake, _ := deployed.VotingContract.GetStake(nil, bytes, context.AuthUser1.From)

	if stake.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected stake to be > 0, got: %v", stake)
	}
}

func TestPollClosed(t *testing.T) {
	// should still be open
	bytes, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(2*test.MIN_VOTE_BY))
	closed, _ := deployed.VotingContract.PollClosed(nil, bytes)

	if closed != false {
		t.Errorf("Expected pollClosed to be false, got: %v", closed)
	}

	// now move time forward so that voteBy has elapsed
	context.Blockchain.AdjustTime(test.MIN_VOTE_BY * time.Second)
	context.Blockchain.Commit()

	updated, _ := deployed.VotingContract.PollClosed(nil, bytes)

	if updated != true {
		t.Errorf("Expected pollClosed to be true, got: %v", updated)
	}
}

func TestDidPass(t *testing.T) {
	bytes, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(2*test.MIN_VOTE_BY))
	// the one total vote will do it
	passed, _ := deployed.VotingContract.DidPass(nil, bytes, big.NewInt(50))

	if passed != true {
		t.Errorf("Expected didPass to be true, got: %v", passed)
	}
}

// test that the voting stake remains after reparam is resolved...
func TestResolveReparam(t *testing.T) {
	bytes, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(2*test.MIN_VOTE_BY))

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
	if voteBy.Cmp(big.NewInt(2*test.MIN_VOTE_BY)) != 0 {
		t.Errorf("Expected voteBy to be %v, got: %v", 2*test.MIN_VOTE_BY, voteBy)
	}

	stake, _ := deployed.VotingContract.GetStake(nil, bytes, context.AuthUser1.From)

	if stake.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected stake to be > 0, got: %v", stake)
	}
}
