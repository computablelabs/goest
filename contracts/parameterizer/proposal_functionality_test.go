package parameterizer

import (
	// "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	// "github.com/ethereum/go-ethereum/crypto/sha3"
	"math/big"
	"testing"
)

func TestPropExists(t *testing.T) {
	t.Log("Parameterizer returns falsy for a non-existing proposal")

	// in go the bytes32 is a byte slice of 32 len
	notThere := "this is not there"
	var bytes32 [32]byte
	copy(bytes32[:], notThere) // TODO we could likely just generate a random one...

	exists, err := deployed.ParameterizerContract.PropExists(&bind.CallOpts{
		From: context.AuthOwner.From,
	}, bytes32)

	if err != nil {
		t.Fatalf("Error checking for proposal: %v", err)
	}

	if exists != false {
		t.Fatalf("Expected propExists to be false, got : %v", exists)
	}
}

func TestProposeReparameterization(t *testing.T) {
	t.Log("Parameterizer can create a new proposal")

	_, err := deployed.ParameterizerContract.ProposeReparameterization(&bind.TransactOpts{
		From:     context.AuthOwner.From,
		Signer:   context.AuthOwner.Signer,
		GasLimit: 1000000,
	}, "voteQuorum", big.NewInt(51))

	if err != nil {
		t.Fatalf("Error creating proposal: %v", err)
	}

	context.Blockchain.Commit()

	// NOTE: leaving for later reference - this has to do with attempting to recreate the
	// propHash created by solidity for a proposal id. Instead of this, what I think I'll do
	// is setup an event listener and retain the propHash from there, which is included...

	// to check, we have to re-create the solidity keccak, which is a hash of both vals
	// TODO abstract into a helper
	// nameArg, _ := abi.NewType("string")
	// valArg, _ := abi.NewType("uint")
	// args := abi.Arguments{{Type: nameArg}, {Type: valArg}}
	// bytes, _ := args.Pack(
	// common.HexToAddress("0x0000000000000000000000000000000000000000"),
	// [32]byte{'I', 'D', '1'},
	// big.NewInt(42),
	// )

	// var buf []byte
	// hash := sha3.NewKeccak256()
	// hash.Write(bytes)
	// buf = hash.Sum(buf)
	// fuckin finally...
	// propHash := hexutil.Encode(bug)) // have to then cast as [32]byte
}
