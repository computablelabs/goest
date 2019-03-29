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
	// cast a vote for, member may need funds
	memBal, _ := deployed.MarketTokenContract.BalanceOf(nil, context.AuthMember1.From)
	if memBal.Cmp(big.NewInt(ONE_GWEI)) == -1 {
		// transfer one
		_, transErr := deployed.MarketTokenContract.Transfer(&bind.TransactOpts{
			From:     context.AuthFactory.From,
			Signer:   context.AuthFactory.Signer,
			GasPrice: big.NewInt(ONE_GWEI * 2),
			GasLimit: 1000000,
		}, context.AuthMember1.From, big.NewInt(ONE_GWEI))

		if transErr != nil {
			t.Fatalf("Error transferring tokens to member: %v", transErr)
		}

		context.Blockchain.Commit()
	}

	// member will need to have approved the voting contract to spend
	allowed, _ := deployed.MarketTokenContract.Allowance(nil, context.AuthMember1.From, deployed.VotingAddress)
	if !(allowed.Cmp(big.NewInt(ONE_GWEI)) >= 0) {
		_, approveErr := deployed.MarketTokenContract.Approve(&bind.TransactOpts{
			From:     context.AuthMember1.From,
			Signer:   context.AuthMember1.Signer,
			GasPrice: big.NewInt(ONE_GWEI * 2),
			GasLimit: 1000000,
		}, deployed.VotingAddress, big.NewInt(ONE_GWEI))

		if approveErr != nil {
			t.Fatalf("Error approving market contract to spend: %v", approveErr)
		}

		context.Blockchain.Commit()
	}

	_, voteErr := deployed.VotingContract.Vote(&bind.TransactOpts{
		From:     context.AuthMember1.From,
		Signer:   context.AuthMember1.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 150000,
	}, hash, big.NewInt(1))

	if voteErr != nil {
		t.Fatalf("Error voting for candidate: %v", voteErr)
	}

	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

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
