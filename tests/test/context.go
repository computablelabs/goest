// Package containing objects and functions which aid in the contruction and
// execution of our Geth-based spec suite
package test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"math/big"
)

// Ctx, holds Auth objects capable of signing transactions.
// Not every spec will use every Auth object but they all should be available here,
// rather than scattered about many specs.
// Also holds the Geth simulated backend.
type Ctx struct {
	AuthOwner   *bind.TransactOpts
	AuthBackend *bind.TransactOpts
	AuthUser1   *bind.TransactOpts
	AuthUser2   *bind.TransactOpts
	AuthUser3   *bind.TransactOpts
	Blockchain  *backends.SimulatedBackend
}

// GetContext returns a hydrated Ctx struct, ready for use.
// Given a bal argument, it assigns this as the wallet balance for
// each authorization object in the Ctx
func GetContext(bal *big.Int) *Ctx {
	authOwn := getAuthObject()
	authBac := getAuthObject()
	authU1 := getAuthObject()
	authU2 := getAuthObject()
	authU3 := getAuthObject()
	alloc := make(core.GenesisAlloc)
	alloc[authOwn.From] = core.GenesisAccount{Balance: bal}
	alloc[authBac.From] = core.GenesisAccount{Balance: bal}
	alloc[authU1.From] = core.GenesisAccount{Balance: bal}
	alloc[authU2.From] = core.GenesisAccount{Balance: bal}
	alloc[authU3.From] = core.GenesisAccount{Balance: bal}
	// 2nd arg is a gas limit, a uint64. we'll use 4.7 million
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &Ctx{
		AuthOwner:   authOwn,
		AuthBackend: authBac,
		AuthUser1:   authU1,
		AuthUser2:   authU2,
		AuthUser3:   authU3,
		Blockchain:  bc,
	}
}
