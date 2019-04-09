package votingtest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
	"time"
)

func TestAddCandidate(t *testing.T) {
	listingHash, _ := deployed.ListingContract.GetHash(nil, "Voting Candidate Test")
	_, listErr := deployed.ListingContract.List(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), listingHash)
	test.IfNotNil(t, listErr, fmt.Sprintf("Error adding candidate: %v", listErr))

	context.Blockchain.Commit()
}

func TestIsCandidate(t *testing.T) {
	bytes, _ := deployed.ListingContract.GetHash(nil, "Voting Candidate Test")

	isCandidate, _ := deployed.VotingContract.IsCandidate(nil, bytes)

	if isCandidate != true {
		t.Errorf("Expected isCandidate be true, got: %v", isCandidate)
	}
}

func TestGetCandidate(t *testing.T) {
	bytes, _ := deployed.ListingContract.GetHash(nil, "Voting Candidate Test")

	kind, owner, stake, voteBy, yea, nay, err := deployed.VotingContract.GetCandidate(nil, bytes)
	test.IfNotNil(t, err, fmt.Sprintf("Error getting candidate: %v", err))

	if kind.Cmp(test.APPLICATION) != 0 {
		t.Errorf("Expected kind to be application, got: %v", kind)
	}

	if owner != context.AuthUser1.From {
		t.Errorf("Expected %v to be %v", owner, context.AuthUser1.From)
	}

	if stake.Cmp(big.NewInt(test.ONE_GWEI)) != 0 {
		t.Error("Expected stake to be 1e9")
	}

	// should be > 0
	if voteBy.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Expected vote by date to be > 0, got: %v", voteBy)
	}

	if yea.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected no votes to be cast yet, got: %v", yea)
	}

	if nay.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected no votes to be cast yet, got: %v", nay)
	}
}

func TestCandidateIs(t *testing.T) {
	bytes, _ := deployed.ListingContract.GetHash(nil, "Voting Candidate Test")

	// it's not a challenge...
	shouldBeFalse, _ := deployed.VotingContract.CandidateIs(nil, bytes, test.CHALLENGE)

	if shouldBeFalse != false {
		t.Errorf("Expected kind == challenge to be false, got: %v", shouldBeFalse)
	}

	shouldBeTrue, _ := deployed.VotingContract.CandidateIs(nil, bytes, test.APPLICATION)

	if shouldBeTrue != true {
		t.Errorf("Expected kind == application to be true, got: %v", shouldBeTrue)
	}

	// test the case where the option does not exist
	alsoBeFalse, _ := deployed.VotingContract.CandidateIs(nil, bytes, big.NewInt(4))

	if alsoBeFalse != false {
		t.Errorf("Expected kind == 4 to be false, got: %v", alsoBeFalse)
	}
}

func TestRemoveCandidate(t *testing.T) {
	// only privileged contracts can call for candidate removal, move past the vote by and call for resolution.
	// this should trigger the candidate removal as there were no votes (and no datatrust hash)
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	bytes, _ := deployed.ListingContract.GetHash(nil, "Voting Candidate Test")

	_, err := deployed.ListingContract.ResolveApplication(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 150000), bytes)
	test.IfNotNil(t, err, fmt.Sprintf("Error removing candidate: %v", err))

	context.Blockchain.Commit()

	notCandidate, _ := deployed.VotingContract.IsCandidate(nil, bytes)

	if notCandidate != false {
		t.Errorf("Expected isCandidate be false, got: %v", notCandidate)
	}
}
