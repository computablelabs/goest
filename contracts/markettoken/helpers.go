package markettoken

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
	AuthOwner  *bind.TransactOpts
	AuthUser   *bind.TransactOpts
	AuthOther  *bind.TransactOpts
	AuthMarket *bind.TransactOpts
	Blockchain *backends.SimulatedBackend
}

type dep struct {
	Address     common.Address
	Contract    *MarketToken
	Transaction *types.Transaction
}

func Deploy(initialBalance *big.Int, c *ctx) (*dep, error) {
	// this method is generated in the burnable.go compiled sol class
	addr, trans, cont, err := DeployMarketToken(
		c.AuthOwner,
		c.Blockchain,
		c.AuthOwner.From,
		initialBalance,
	)

	if err != nil {
		return nil, err
	}

	// commit the deploy before returning
	// c.Blockchain.Commit()

	// the market token must have an address for the market contract
	_, setErr := cont.SetMarket(&bind.TransactOpts{
		From:     c.AuthOwner.From,
		Signer:   c.AuthOwner.Signer,
		GasPrice: big.NewInt(2000000000), // 2 Gwei
		GasLimit: 100000,
	}, c.AuthMarket.From)

	if setErr != nil {
		return nil, setErr
	}

	c.Blockchain.Commit()

	return &dep{Address: addr, Contract: cont, Transaction: trans}, nil
}

func SetupBlockchain(accountBalance *big.Int) *ctx {
	// generate a new key, toss the error for now as it shouldnt happen
	keyOwner, _ := crypto.GenerateKey()
	keyUser, _ := crypto.GenerateKey()
	keyOther, _ := crypto.GenerateKey()
	keyMarket, _ := crypto.GenerateKey()
	authOwner := bind.NewKeyedTransactor(keyOwner)
	authUser := bind.NewKeyedTransactor(keyUser)
	authOther := bind.NewKeyedTransactor(keyOther)
	authMarket := bind.NewKeyedTransactor(keyMarket)
	alloc := make(core.GenesisAlloc)
	alloc[authOwner.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authUser.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authOther.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMarket.From] = core.GenesisAccount{Balance: accountBalance}
	// 2nd arg is a gas limit, a uint64. we'll use 4.7 million
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &ctx{
		AuthOwner:  authOwner,
		AuthUser:   authUser,
		AuthOther:  authOther,
		AuthMarket: authMarket,
		Blockchain: bc,
	}
}
