package token

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

func Setup(accountBalance *big.Int) (
	*bind.TransactOpts,
	*backends.SimulatedBackend) {
	// generate a new key, toss the error for now as it shouldnt happen
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: accountBalance}
	blockchain := backends.NewSimulatedBackend(alloc)

	return auth, blockchain
}
