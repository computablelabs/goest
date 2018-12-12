package voting

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
)

func TestAddCandidate(t *testing.T) {
	bytes := GenBytes32("iCanHazListing")

	_, err := deployed.VotingContract.AddCandidate(&bind.TransactOpts{
		From:     context.AuthMarket.From,
		Signer:   context.AuthMarket.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 150000,
	}, bytes, APPLICATION, big.NewInt(2))

	if err != nil {
		t.Fatalf("Error adding candidate: %v", err)
	}

	context.Blockchain.Commit()
}

func TestGetCandidates(t *testing.T) {
	// we should have at least one candidate now, fetch the array
	candidates, _ := deployed.VotingContract.GetCandidates(nil)
	// it's an array of the bytes32s
	length := len(candidates)

	if !(length > 0) {
		t.Fatalf("Expected candidates length to be > 0, got: %v", length)
	}
}

func TestGetCandidate(t *testing.T) {
	bytes := GenBytes32("iCanHazListing")

	kind, voteBy, votes, err := deployed.VotingContract.GetCandidate(nil, bytes)

	if err != nil {
		t.Fatalf("Error getting candidate: %v", err)
	}

	if kind != APPLICATION {
		t.Fatalf("Expected kind to be application, got: %v", kind)
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
	alsoBeFalse, _ := deployed.VotingContract.CandidateIs(nil, bytes, 4)

	if alsoBeFalse != false {
		t.Fatalf("Expected kind == 4 to be false, got: %v", alsoBeFalse)
	}
}

func TestRemoveCandidate(t *testing.T) {
	candidates, _ := deployed.VotingContract.GetCandidates(nil)
	// it's an array of the bytes32s
	length := len(candidates)

	if !(length > 0) {
		t.Fatalf("Expected candidate length to be > 0, got : %v", length)
	}

	bytes := GenBytes32("iCanHazListing")

	_, err := deployed.VotingContract.RemoveCandidate(&bind.TransactOpts{
		From:     context.AuthMarket.From,
		Signer:   context.AuthMarket.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 150000,
	}, bytes)

	if err != nil {
		t.Fatalf("Error removing candidate: %v", err)
	}

	context.Blockchain.Commit()

	// length should now be one less than it was
	noCandidates, _ := deployed.VotingContract.GetCandidates(nil)
	// it's an array of the bytes32s
	noLength := len(noCandidates)

	if noLength >= length {
		t.Fatalf("Expected candidates length of %v to be 1 less than %v", noLength, length)
	}
}
