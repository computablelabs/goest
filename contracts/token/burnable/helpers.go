package burnable

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type context struct {
	auth       *bind.TransactOpts
	blockchain *backends.SimulatedBackend
}

type deployed struct {
	address     common.Address
	contract    *ConstructableBurnable
	transaction *types.Transaction
}

func Deploy(initialBalance *big.Int, c *context) (*deployed, error) {
	// this method is generated in the burnable.go compiled sol class
	addr, trans, cont, err := DeployConstructableBurnable(
		c.auth,
		c.blockchain,
		c.auth.From,
		initialBalance,
	)

	if err != nil {
		return nil, err
	}

	// commit the deploy before returning
	c.blockchain.Commit()

	return &deployed{address: addr, contract: cont, transaction: trans}, nil
}

func SetupBlockchain(accountBalance *big.Int) *context {
	// generate a new key, toss the error for now as it shouldnt happen
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: accountBalance}
	// 2nd arg is a gas limit, a uint64. we'll use 5 million
	bc := backends.NewSimulatedBackend(alloc, 5000000)

	return &context{auth: auth, blockchain: bc}
}
