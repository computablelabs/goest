package parameterizertest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
	"time"
)

func TestParameterize(t *testing.T) {
	// auth member will need at least the stake
	transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser1.From,
		big.NewInt(2*test.ONE_GWEI))
	test.IfNotNil(t, transErr, "Error maybe transferring market tokens")

	// member will need to have approved the voting contract to spend at least the stake
	incErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser1, deployed.VotingAddress,
		big.NewInt(2*test.ONE_GWEI))
	test.IfNotNil(t, incErr, "Error maybe transferring market token approval")

	_, err := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), test.VOTE_BY, big.NewInt(2*test.MIN_VOTE_BY))
	test.IfNotNil(t, err, fmt.Sprintf("Error creating proposal: %v", err))

	context.Blockchain.Commit()

	// we should see a candidate now
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(2*test.MIN_VOTE_BY))
	isCan, _ := deployed.VotingContract.IsCandidate(nil, paramHash)

	if !isCan {
		t.Errorf("Expected isCandidate to be true, got: %v", isCan)
	}
}

func GetReparam(t *testing.T) {
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(2*test.MIN_VOTE_BY))
	name, val, _ := deployed.ParameterizerContract.GetReparam(nil, paramHash)

	if name.Cmp(test.VOTE_BY) != 0 {
		t.Errorf("Expected name to be 'voteBy', got: %v", name)
	}

	if val.Cmp(big.NewInt(2*test.MIN_VOTE_BY)) != 0 {
		t.Errorf("Expected value to be %v, got: %v", 2*test.MIN_VOTE_BY, val)
	}
}

func ResolveReparam(t *testing.T) {
	oldVoteBy, _ := deployed.ParameterizerContract.GetVoteBy(nil)

	// make sure its not the proposed candidate's amount (2*test.MIN_VOTE_BY)
	if oldVoteBy.Cmp(big.NewInt(2*test.MIN_VOTE_BY)) == 0 {
		t.Errorf("Expected the old voteBy param to not be %v, got: %v", 2*test.MIN_VOTE_BY, oldVoteBy)
	}

	// how we will fetch this proposal
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(2*test.MIN_VOTE_BY))

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

	// voteBy should be changed now
	newVoteBy, _ := deployed.ParameterizerContract.GetVoteBy(nil)

	if newVoteBy.Cmp(big.NewInt(2*test.MIN_VOTE_BY)) != 0 {
		t.Errorf("Expected the new voteBy param to be %v, got: %v", 2*test.MIN_VOTE_BY, newVoteBy)
	}

	// should have cleaned up the voting candidate
	isCanNow, _ := deployed.VotingContract.IsCandidate(nil, paramHash)

	if isCanNow {
		t.Errorf("Expected isCandidate to be false, got: %v", isCanNow)
	}

	// should have cleaned up the proposal
	name, val, _ := deployed.ParameterizerContract.GetReparam(nil, paramHash)

	if name.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected name to be falsy, got: %v", name)
	}

	if val.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected value to be 0, got: %v", val)
	}
}
