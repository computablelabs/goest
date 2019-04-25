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
		big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, transErr, "Error maybe transferring market tokens")

	// member will need to have approved the voting contract to spend at least the stake
	incErr := test.MaybeIncreaseMarketTokenApproval(context, deployed, context.AuthUser1, deployed.VotingAddress,
		big.NewInt(test.ONE_GWEI))
	test.IfNotNil(t, incErr, "Error maybe transferring market token approval")

	// we need a candidate, a privileged contract must add it...
	_, candErr := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 200000), test.VOTE_BY, big.NewInt(101))
	test.IfNotNil(t, candErr, fmt.Sprintf("Error adding candidate: %v", candErr))

	context.Blockchain.Commit()

	bytes, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(101))

	// should be no votes atm
	_, _, _, _, preYea, _, _ := deployed.VotingContract.GetCandidate(nil, bytes)

	if preYea.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected number of votes to be 0, got: %v", preYea)
	}

	// cast a yea vote
	_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), bytes, big.NewInt(1))
	test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))

	context.Blockchain.Commit()

	// has been recorded in the candidate's votes
	_, _, _, _, yea, _, _ := deployed.VotingContract.GetCandidate(nil, bytes)

	if yea.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected number of votes to be > 0, got: %v", yea)
	}
}

func TestPollClosed(t *testing.T) {
	// should still be open
	bytes, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(101))
	closed, _ := deployed.VotingContract.PollClosed(nil, bytes)

	if closed != false {
		t.Errorf("Expected pollClosed to be false, got: %v", closed)
	}

	// now move time forward so that voteBy has elapsed
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	updated, _ := deployed.VotingContract.PollClosed(nil, bytes)

	if updated != true {
		t.Errorf("Expected pollClosed to be true, got: %v", updated)
	}
}

func TestDidPass(t *testing.T) {
	bytes, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(101))
	// the one total vote will do it
	passed, _ := deployed.VotingContract.DidPass(nil, bytes, big.NewInt(50))

	if passed != true {
		t.Errorf("Expected didPass to be true, got: %v", passed)
	}
}
