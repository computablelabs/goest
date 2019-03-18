package datatrust

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"strings"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	_, regErr := deployed.DatatrustContract.Register(&bind.TransactOpts{
		From:     context.AuthBackend.From,
		Signer:   context.AuthBackend.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 500000,
	}, "http://www.icanhazbackend.com")

	if regErr != nil {
		t.Fatalf("Error registering for backend status: %v", regErr)
	}

	context.Blockchain.Commit()

	// url will be stored as we wait on the voting
	url, _ := deployed.DatatrustContract.GetBackendUrl(nil)
	if !strings.Contains(url, "icanhazbackend") {
		t.Fatalf("expected url to be %v, got: %v", "http://www.icanhazbackend.com", url)
	}

	// we should have the candidate
	hash, _ := deployed.DatatrustContract.GetHash(nil, "http://www.icanhazbackend.com")
	isReg, _ := deployed.VotingContract.CandidateIs(nil, hash, big.NewInt(REGISTRATION))
	if isReg != true {
		t.Fatalf("Expected is registered to be true, got: %v", isReg)
	}
}

func TestResolveRegistration(t *testing.T) {
	// we should not have an address yet
	nodress, _ := deployed.DatatrustContract.GetBackendAddress(nil)

	hash, _ := deployed.DatatrustContract.GetHash(nil, "http://www.icanhazbackend.com")
	// we'll need a council member
	isMember, _ := deployed.VotingContract.InCouncil(nil, context.AuthMember1.From)
	if isMember != true {
		_, councilErr := deployed.VotingContract.AddToCouncil(&bind.TransactOpts{
			From:     context.AuthListing.From,
			Signer:   context.AuthListing.Signer,
			GasPrice: big.NewInt(ONE_GWEI * 2),
			GasLimit: 100000,
		}, context.AuthMember1.From)

		if councilErr != nil {
			t.Fatal("Error adding member to council")
		}
	}

	context.Blockchain.Commit()

	// cast a vote for
	_, voteErr := deployed.VotingContract.Vote(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 100000,
	}, hash)

	if voteErr != nil {
		t.Fatalf("Error voting for candidate: %v", voteErr)
	}

	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	// any council member can call for resolution
	_, resolveErr := deployed.DatatrustContract.ResolveRegistration(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, hash)

	if resolveErr != nil {
		t.Fatalf("Error resolving application: %v", resolveErr)
	}

	context.Blockchain.Commit()

	// should be listed our backend now, with an address
	address, _ := deployed.DatatrustContract.GetBackendAddress(nil)
	if bytes.Equal(nodress.Bytes(), address.Bytes()) {
		t.Fatalf("Expected backend address to be set, got: %v", address)
	}
}

func TestSetDataHash(t *testing.T) {
	// create a bogus listing hash, we can use datatrust's method for this
	listing_hash, _ := deployed.DatatrustContract.GetHash(nil, "look-at-me-im-a-listing")
	// it's empty atm
	no_hash, _ := deployed.DatatrustContract.GetDataHash(nil, listing_hash)
	// make a dummy data_hash
	data_hash, _ := deployed.DatatrustContract.GetHash(nil, "look-at-me-im-data")
	_, setErr := deployed.DatatrustContract.SetDataHash(&bind.TransactOpts{
		From:     context.AuthBackend.From,
		Signer:   context.AuthBackend.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 500000,
	}, listing_hash, data_hash)

	if setErr != nil {
		t.Fatalf("Error setting data hash: %v", setErr)
	}

	context.Blockchain.Commit()

	// no longer empty
	hash, _ := deployed.DatatrustContract.GetDataHash(nil, listing_hash)
	if hash == no_hash {
		t.Fatalf("Expected %v not to be %v", hash, no_hash)
	}
}
