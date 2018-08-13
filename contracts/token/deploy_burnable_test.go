package token

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

// Expect to deploy correctly
func TestDeployConstructableBurnable(t *testing.T) {
	// setup our simulated block chain TODO abstract into test helper
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc)

	// do the actual deploy. args go [auth, bc, ...args to the contructor]
	address, _, _, err := DeployConstructableBurnable(
		auth,
		blockchain,
		auth.From,
		big.NewInt(1000000),
	)

	// do stuff
	blockchain.Commit()

	if err != nil {
		t.Fatalf("Failed to deploy the burnable token contract: %v", err)
	}

	if len(address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address to be returned from deploy, got empty byte array instead")
	}
}
