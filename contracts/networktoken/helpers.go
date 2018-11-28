package networktoken

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
	Alloc         core.GenesisAlloc // a map of accounts as { address: account }
	AuthFactory   *bind.TransactOpts
	AuthMember    *bind.TransactOpts
	Blockchain    *backends.SimulatedBackend
	OtherUser     common.Address
	OtherContract common.Address
}

type dep struct {
	Address     common.Address
	Contract    *NetworkToken
	Transaction *types.Transaction
}

func Deploy(initialBalance *big.Int, c *ctx) (*dep, error) {
	// this method is generated in the burnable.go compiled sol class
	addr, trans, cont, err := DeployNetworkToken(
		c.AuthFactory,
		c.Blockchain,
		c.AuthFactory.From,
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
	keyFac, _ := crypto.GenerateKey()
	authFac := bind.NewKeyedTransactor(keyFac)
	keyMem, _ := crypto.GenerateKey()
	authMem := bind.NewKeyedTransactor(keyMem)
	alloc := make(core.GenesisAlloc)
	alloc[authFac.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem.From] = core.GenesisAccount{Balance: accountBalance}
	// 2nd arg is a gas limit, a uint64. we'll use 4.7M
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &ctx{
		Alloc:         alloc,
		AuthFactory:   authFac,
		AuthMember:    authMem,
		Blockchain:    bc,
		OtherUser:     common.HexToAddress("0xabc"),
		OtherContract: common.HexToAddress("0xdef"),
	}
}
