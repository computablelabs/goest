package datatrusttest

import (
	"bytes"
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"strings"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	preAddress, _ := deployed.DatatrustContract.GetBackendAddress(nil)
	backend_address := context.AuthBackend.From
	// The backend is already registered
	if preAddress == backend_address {
		// Set datatrust utl
		_, setErr := deployed.DatatrustContract.SetBackendUrl(test.GetTxOpts(context.AuthBackend, nil,
			big.NewInt(test.ONE_GWEI*2), 500000), "http://www.icanhazbackend.com")
		test.IfNotNil(t, setErr, fmt.Sprintf("Error setting backend url: %v", setErr))
		context.Blockchain.Commit()
	} else {
		// auth backend will need at least the stake
		transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthBackend.From,
			big.NewInt(2*test.ONE_GWEI))
		test.IfNotNil(t, transErr, "Error maybe transferring market tokens")

		// backend will need to have approved the voting contract to spend at least the stake
		incErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthBackend, deployed.VotingAddress,
			big.NewInt(2*test.ONE_GWEI))
		test.IfNotNil(t, incErr, "Error maybe transferring market token approval")

		// Register the datatrust
		_, regErr := deployed.DatatrustContract.Register(test.GetTxOpts(context.AuthBackend, nil,
			big.NewInt(test.ONE_GWEI*2), 500000), "http://www.icanhazbackend.com")
		test.IfNotNil(t, regErr, fmt.Sprintf("Error registering for backend status: %v", regErr))
		context.Blockchain.Commit()

		// url will not be stored as we wait on the voting
		url, _ := deployed.DatatrustContract.GetBackendUrl(nil)
		if strings.Contains(url, "icanhazbackend") {
			t.Errorf("Url was improperly updated to: %v", url)
		}

		// we should have the candidate
		hash, _ := deployed.DatatrustContract.GetHash(nil, "http://www.icanhazbackend.com")
		isReg, _ := deployed.VotingContract.CandidateIs(nil, hash, test.REGISTRATION)
		if isReg != true {
			t.Errorf("Expected is registered to be true, got: %v", isReg)
		}
	}
}

func TestReRegister(t *testing.T) {
	// Test what happens when a new party tries to register as datatrust
	// auth user will need at least the stake
	transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser1.From,
		big.NewInt(2*test.ONE_GWEI))
	test.IfNotNil(t, transErr, "Error maybe transferring market tokens")

	// backend will need to have approved the voting contract to spend at least the stake
	incErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser1, deployed.VotingAddress,
		big.NewInt(2*test.ONE_GWEI))
	test.IfNotNil(t, incErr, "Error maybe transferring market token approval")

	// Try registering a new datatrust by a different user
	_, regErr := deployed.DatatrustContract.Register(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 500000), "http://www.thenewbackend.com")
	test.IfNotNil(t, regErr, fmt.Sprintf("Error registering for backend status: %v", regErr))
	context.Blockchain.Commit()

	// url will not be stored as we wait on the voting
	url, _ := deployed.DatatrustContract.GetBackendUrl(nil)
	if strings.Contains(url, "thenewbackend") {
		t.Errorf("Url was improperly updated to: %v", url)
	}

	// we should have the candidate
	hash, _ := deployed.DatatrustContract.GetHash(nil, "http://www.thenewbackend.com")
	isReg, _ := deployed.VotingContract.CandidateIs(nil, hash, test.REGISTRATION)
	if isReg != true {
		t.Errorf("Expected is registered to be true, got: %v", isReg)
	}
}

func TestResolveRegistration(t *testing.T) {
	// we may have already resolved a registration
	nodress, _ := deployed.DatatrustContract.GetBackendAddress(nil)
	if nodress != context.AuthBackend.From {
		hash, _ := deployed.DatatrustContract.GetHash(nil, "http://www.icanhazbackend.com")

		// auth member will need at least the stake
		transErr := test.MaybeTransferMarketToken(context, deployed, context.AuthOwner, context.AuthUser1.From,
			big.NewInt(test.ONE_GWEI))
		test.IfNotNil(t, transErr, "Error maybe transferring market tokens")

		// member will need to have approved the voting contract to spend at least the stake
		incErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser1, deployed.VotingAddress,
			big.NewInt(test.ONE_GWEI))
		test.IfNotNil(t, incErr, "Error maybe transferring market token approval")

		_, voteErr := deployed.VotingContract.Vote(test.GetTxOpts(context.AuthUser1, nil,
			big.NewInt(test.ONE_GWEI*2), 150000), hash, big.NewInt(1))
		test.IfNotNil(t, voteErr, fmt.Sprintf("Error voting for candidate: %v", voteErr))

		context.Blockchain.Commit()

		// move past the voteBy
		context.Blockchain.AdjustTime(100 * time.Second)
		context.Blockchain.Commit()

		_, resolveErr := deployed.DatatrustContract.ResolveRegistration(test.GetTxOpts(
			context.AuthUser1, nil, big.NewInt(test.ONE_GWEI*2), 1000000), hash)
		test.IfNotNil(t, resolveErr, fmt.Sprintf("Error resolving application: %v", resolveErr))

		context.Blockchain.Commit()

		// should be listed our backend now, with an address
		address, _ := deployed.DatatrustContract.GetBackendAddress(nil)
		if bytes.Equal(nodress.Bytes(), address.Bytes()) {
			t.Errorf("Expected backend address to be set, got: %v", address)
		}
	}
}

func TestSetDataHash(t *testing.T) {
	//// create a bogus listing hash, we can use datatrust's method for this
	listing_hash, _ := deployed.DatatrustContract.GetHash(nil, "look-at-me-im-a-listing")
	//// it's empty atm
	no_hash, _ := deployed.DatatrustContract.GetDataHash(nil, listing_hash)
	// make a dummy data_hash
	data_hash, _ := deployed.DatatrustContract.GetHash(nil, "look-at-me-im-data")
	_, setErr := deployed.DatatrustContract.SetDataHash(test.GetTxOpts(context.AuthBackend,
		nil, big.NewInt(test.ONE_GWEI*2), 500000), listing_hash, data_hash)
	test.IfNotNil(t, setErr, fmt.Sprintf("Error setting data hash: %v", setErr))

	context.Blockchain.Commit()

	// no longer empty
	hash, _ := deployed.DatatrustContract.GetDataHash(nil, listing_hash)
	if hash == no_hash {
		t.Errorf("Expected %v not to be %v", hash, no_hash)
	}
}
