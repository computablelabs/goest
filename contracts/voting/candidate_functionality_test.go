package voting

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"testing"
)

func TestAddCandidate(t *testing.T) {
	t.Log("Voting contract should add a candidate on demand")

	bytes := genBytes32("iCanHazListing")

	_, err := deployed.VotingContract.AddCandidate(&bind.TransactOpts{
		From:     context.AuthMarket.From,
		Signer:   context.AuthMarket.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 150000,
	}, "application", bytes, big.NewInt(300))

	if err != nil {
		t.Fatalf("Error adding candidate: %v", err)
	}

	context.Blockchain.Commit()
}

func TestGetCandidates(t *testing.T) {
	t.Log("Voting contract should fetch all candidates on demand")
	// we should have at least one candidate now, fetch the array
	candidates, _ := deployed.VotingContract.GetCandidates(nil)
	// it's an array of the bytes32s
	length := len(candidates)

	if !(length > 0) {
		t.Fatalf("Expected candidates length to be > 0, got: %v", length)
	}
}

func TestGetCandidate(t *testing.T) {
	t.Log("Voting contract should fetch a candidate on demand")
	bytes := genBytes32("iCanHazListing")

	kind, voteBy, votes, err := deployed.VotingContract.GetCandidate(nil, bytes)

	if err != nil {
		t.Fatalf("Error getting candidate: %v", err)
	}

	if kind != "application" {
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
	t.Log("Voting contract should correctly check the candidate type")
	bytes := genBytes32("iCanHazListing")

	// it's not a challenge...
	shouldBeFalse, _ := deployed.VotingContract.CandidateIs(nil, bytes, "challenge")

	if shouldBeFalse != false {
		t.Fatalf("Expected kind == challenge to be false, got: %v", shouldBeFalse)
	}

	shouldBeTrue, _ := deployed.VotingContract.CandidateIs(nil, bytes, "application")

	if shouldBeTrue != true {
		t.Fatalf("Expected kind == application to be true, got: %v", shouldBeTrue)
	}

	// test the case where length is equal (11), but the word itself is wrong
	alsoBeFalse, _ := deployed.VotingContract.CandidateIs(nil, bytes, "beetlejuice")

	if alsoBeFalse != false {
		t.Fatalf("Expected kind == beetlejuice to be false, got: %v", alsoBeFalse)
	}
}

func TestRemoveCandidate(t *testing.T) {
	t.Log("Voting contract should correctly remove a candidate on demand")
	candidates, _ := deployed.VotingContract.GetCandidates(nil)
	// it's an array of the bytes32s
	length := len(candidates)

	if !(length > 0) {
		t.Fatalf("Expected candidate length to be > 0, got : %v", length)
	}

	bytes := genBytes32("iCanHazListing")

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
