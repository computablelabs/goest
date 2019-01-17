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

const ONE_WEI = 1000000000000000000
const ONE_GWEI = 1000000000

type ctx struct {
	AuthFactory *bind.TransactOpts
	AuthMember1 *bind.TransactOpts
	AuthMember2 *bind.TransactOpts
	AuthMarket  *bind.TransactOpts
	Blockchain  *backends.SimulatedBackend
}

type dep struct {
	Address     common.Address
	Contract    *MarketToken
	Transaction *types.Transaction
}

func Deploy(initialBalance *big.Int, c *ctx) (*dep, error) {
	// this method is generated in the burnable.go compiled sol class
	addr, trans, cont, err := DeployMarketToken(
		c.AuthFactory,
		c.Blockchain,
		c.AuthFactory.From,
		initialBalance,
	)

	if err != nil {
		return nil, err
	}

	// commit the deploy before returning
	// c.Blockchain.Commit()

	// the market token must have an address for the market contract
	_, setErr := cont.SetPrivilegedContracts(&bind.TransactOpts{
		From:     c.AuthFactory.From,
		Signer:   c.AuthFactory.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2), // 2 Gwei
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
	keyFac, _ := crypto.GenerateKey()
	keyMem1, _ := crypto.GenerateKey()
	keyMem2, _ := crypto.GenerateKey()
	keyMark, _ := crypto.GenerateKey()
	authFac := bind.NewKeyedTransactor(keyFac)
	authMem1 := bind.NewKeyedTransactor(keyMem1)
	authMem2 := bind.NewKeyedTransactor(keyMem2)
	authMark := bind.NewKeyedTransactor(keyMark)
	alloc := make(core.GenesisAlloc)
	alloc[authFac.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem1.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem2.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMark.From] = core.GenesisAccount{Balance: accountBalance}
	// 2nd arg is a gas limit, a uint64. we'll use 4.7 million
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &ctx{
		AuthFactory: authFac,
		AuthMember1: authMem1,
		AuthMember2: authMem2,
		AuthMarket:  authMark,
		Blockchain:  bc,
	}
}
