package datatrust

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"strings"
	"testing"
	// "time"
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
