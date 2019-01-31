package voting

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestAddToCouncil(t *testing.T) {
	member := common.HexToAddress("0xfgh")

	_, err := deployed.VotingContract.AddToCouncil(&bind.TransactOpts{
		From:     context.AuthMarket.From,
		Signer:   context.AuthMarket.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 100000,
	}, member)

	if err != nil {
		t.Fatalf("Error adding council member: %v", err)
	}

	context.Blockchain.Commit()
}

func TestGetCouncil(t *testing.T) {
	council, _ := deployed.VotingContract.GetCouncil(nil)

	if !(len(council) > 0) {
		t.Fatalf("Expected length of council array to be > 0, got: %v", len(council))
	}
}

func TestInCouncil(t *testing.T) {
	member := common.HexToAddress("0xfgh")
	res, _ := deployed.VotingContract.InCouncil(nil, member)

	if res != true {
		t.Fatalf("Expected inCouncil call to be true, got: %v", res)
	}
}

func TestRemoveFromCouncil(t *testing.T) {
	// what is the array length right now?
	council, _ := deployed.VotingContract.GetCouncil(nil)
	length := len(council)

	member := common.HexToAddress("0xfgh")

	_, err := deployed.VotingContract.RemoveFromCouncil(&bind.TransactOpts{
		From:     context.AuthMarket.From,
		Signer:   context.AuthMarket.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 100000,
	}, member)

	if err != nil {
		t.Fatalf("Error removing council member: %v", err)
	}

	context.Blockchain.Commit()

	// array should be empty now
	councilNow, _ := deployed.VotingContract.GetCouncil(nil)
	lengthNow := len(councilNow)

	if !(length > lengthNow) {
		t.Fatalf("Expected length of council array, %v, to be less than it was, %v", lengthNow, length)
	}

	// member is not detected as council
	res, _ := deployed.VotingContract.InCouncil(nil, member)

	if res != false {
		t.Fatalf("Expected inCouncil call to be false, got: %v", res)
	}
}
