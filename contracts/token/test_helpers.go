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
	// 2nd arg is a gas limit, a uint64. we'll use 5 million
	blockchain := backends.NewSimulatedBackend(alloc, 5000000)

	return auth, blockchain
}
