package voting

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
)

func TestAddCandidate(t *testing.T) {
	bytes := GenBytes32("iCanHazListing")

	_, err := deployed.VotingContract.AddCandidate(&bind.TransactOpts{
		From:     context.AuthListing.From,
		Signer:   context.AuthListing.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 150000,
	}, bytes, APPLICATION, context.AuthMember1.From, big.NewInt(2))

	if err != nil {
		t.Fatalf("Error adding candidate: %v", err)
	}

	context.Blockchain.Commit()
}

func TestIsCandidate(t *testing.T) {
	bytes := GenBytes32("iCanHazListing")

	isCandidate, _ := deployed.VotingContract.IsCandidate(nil, bytes)

	if isCandidate != true {
		t.Fatalf("Expected isCandidate be true, got: %v", isCandidate)
	}
}

func TestGetCandidateCount(t *testing.T) {
	// we should have at least one candidate now, fetch the number
	count, _ := deployed.VotingContract.GetCandidateCount(nil)

	if count.Cmp(UNDEFINED) != 1 {
		t.Fatalf("Expected candidates count to be > 0, got: %v", count)
	}
}

func TestGetCandidate(t *testing.T) {
	bytes := GenBytes32("iCanHazListing")

	kind, owner, voteBy, votes, err := deployed.VotingContract.GetCandidate(nil, bytes)

	if err != nil {
		t.Fatalf("Error getting candidate: %v", err)
	}

	if kind.Cmp(APPLICATION) != 0 {
		t.Fatalf("Expected kind to be application, got: %v", kind)
	}

	if owner != context.AuthMember1.From {
		t.Fatalf("Expected %v to be %v", owner, context.AuthMember1.From)
	}

	// should be > 0
	if voteBy.Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("Expected vote by date to be > 0, got: %v", voteBy)
	}

	if votes.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("Expected no votes to be cast yet, got: %v", votes)
	}
}

func TestCandidateIs(t *testing.T) {
	bytes := GenBytes32("iCanHazListing")

	// it's not a challenge...
	shouldBeFalse, _ := deployed.VotingContract.CandidateIs(nil, bytes, CHALLENGE)

	if shouldBeFalse != false {
		t.Fatalf("Expected kind == challenge to be false, got: %v", shouldBeFalse)
	}

	shouldBeTrue, _ := deployed.VotingContract.CandidateIs(nil, bytes, APPLICATION)

	if shouldBeTrue != true {
		t.Fatalf("Expected kind == application to be true, got: %v", shouldBeTrue)
	}

	// test the case where the option does not exist
	alsoBeFalse, _ := deployed.VotingContract.CandidateIs(nil, bytes, big.NewInt(4))

	if alsoBeFalse != false {
		t.Fatalf("Expected kind == 4 to be false, got: %v", alsoBeFalse)
	}
}

func TestRemoveCandidate(t *testing.T) {
	count, _ := deployed.VotingContract.GetCandidateCount(nil)

	if count.Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("Expected candidates count to be > 0, got : %v", count)
	}

	bytes := GenBytes32("iCanHazListing")

	_, err := deployed.VotingContract.RemoveCandidate(&bind.TransactOpts{
		From:     context.AuthListing.From,
		Signer:   context.AuthListing.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 150000,
	}, bytes)

	if err != nil {
		t.Fatalf("Error removing candidate: %v", err)
	}

	context.Blockchain.Commit()

	// length should now be one less than it was
	newCount, _ := deployed.VotingContract.GetCandidateCount(nil)

	notCandidate, _ := deployed.VotingContract.IsCandidate(nil, bytes)

	if notCandidate != false {
		t.Fatalf("Expected isCandidate be false, got: %v", notCandidate)
	}

	if newCount.Cmp(count) != -1 {
		t.Fatalf("Expected candidates count of %v to be 1 less than %v", newCount, count)
	}
}
