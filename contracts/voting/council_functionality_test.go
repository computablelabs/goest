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
		From:     context.AuthInvesting.From,
		Signer:   context.AuthInvesting.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 100000,
	}, member)

	if err != nil {
		t.Fatalf("Error adding council member: %v", err)
	}

	context.Blockchain.Commit()
}

func TestGetCouncilCount(t *testing.T) {
	count, _ := deployed.VotingContract.GetCouncilCount(nil)

	if count.Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("Expected council count to be > 0, got: %v", count)
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
	// what is the count right now?
	count, _ := deployed.VotingContract.GetCouncilCount(nil)
	member := common.HexToAddress("0xfgh")

	_, err := deployed.VotingContract.RemoveFromCouncil(&bind.TransactOpts{
		From:     context.AuthInvesting.From,
		Signer:   context.AuthInvesting.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 100000,
	}, member)

	if err != nil {
		t.Fatalf("Error removing council member: %v", err)
	}

	context.Blockchain.Commit()

	// array should be empty now
	countNow, _ := deployed.VotingContract.GetCouncilCount(nil)

	if count.Cmp(countNow) != 1 {
		t.Fatalf("Expected council count, %v, to be less than it was, %v", countNow, count)
	}

	// member is not detected as council
	res, _ := deployed.VotingContract.InCouncil(nil, member)

	if res != false {
		t.Fatalf("Expected inCouncil call to be false, got: %v", res)
	}
}
