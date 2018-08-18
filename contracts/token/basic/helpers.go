package basic

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type ctx struct {
	Auth       *bind.TransactOpts
	Blockchain *backends.SimulatedBackend
}

type dep struct {
	Address     common.Address
	Contract    *ConstructableBasic
	Transaction *types.Transaction
}

func Deploy(initialBalance *big.Int, c *ctx) (*dep, error) {
	// this method is generated in the burnable.go compiled sol class
	addr, trans, cont, err := DeployConstructableBasic(
		c.Auth,
		c.Blockchain,
		c.Auth.From,
		initialBalance,
	)

	if err != nil {
		return nil, err
	}

	// commit the deploy before returning
	c.Blockchain.Commit()

	return &dep{Address: addr, Contract: cont, Transaction: trans}, nil
}

func SetupBlockchain(accountBalance *big.Int) *ctx {
	// generate a new key, toss the error for now as it shouldnt happen
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: accountBalance}
	// 2nd arg is a gas limit, a uint64. we'll use 5 million
	bc := backends.NewSimulatedBackend(alloc, 5000000)

	return &ctx{Auth: auth, Blockchain: bc}
}
