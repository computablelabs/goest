package voting

import (
	"github.com/computablelabs/goest/contracts/markettoken"
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

var UNDEFINED = big.NewInt(0)
var APPLICATION = big.NewInt(1)
var CHALLENGE = big.NewInt(2)
var REPARAM = big.NewInt(3)

type ctx struct {
	ListingAddress    common.Address
	InvestingAddress  common.Address
	AuthFactory       *bind.TransactOpts
	AuthParameterizer *bind.TransactOpts
	AuthDatatrust     *bind.TransactOpts
	AuthListing       *bind.TransactOpts
	AuthInvesting     *bind.TransactOpts
	AuthMember1       *bind.TransactOpts
	AuthMember2       *bind.TransactOpts
	Blockchain        *backends.SimulatedBackend
}

type dep struct {
	MarketTokenAddress     common.Address
	MarketTokenContract    *markettoken.MarketToken
	MarketTokenTransaction *types.Transaction
	VotingAddress          common.Address
	VotingContract         *Voting
	VotingTransaction      *types.Transaction
}

func Deploy(initialBalance *big.Int, c *ctx) (*dep, error) {
	marketTokenAddr, marketTokenTrans, marketTokenCont, marketTokenErr := markettoken.DeployMarketToken(
		c.AuthFactory,
		c.Blockchain,
		c.AuthFactory.From,
		initialBalance,
	)

	if marketTokenErr != nil {
		return nil, marketTokenErr
	}

	// commit the deploy before deploying voting
	c.Blockchain.Commit()

	votingAddr, votingTrans, votingCont, votingErr := DeployVoting(
		c.AuthFactory,
		c.Blockchain,
		marketTokenAddr,
	)

	if votingErr != nil {
		return nil, votingErr
	}

	c.Blockchain.Commit()

	return &dep{
		MarketTokenAddress:     marketTokenAddr,
		MarketTokenContract:    marketTokenCont,
		MarketTokenTransaction: marketTokenTrans,
		VotingAddress:          votingAddr,
		VotingContract:         votingCont,
		VotingTransaction:      votingTrans,
	}, nil
}

// helper for creating the solidity bytes32 listing hash
func GenBytes32(name string) [32]byte {
	bytes := [32]byte{}
	copy(bytes[:], []byte(name))
	return bytes
}

func SetupBlockchain(accountBalance *big.Int) *ctx {
	// generate a new key, toss the error for now as it shouldnt happen
	keyList, _ := crypto.GenerateKey()
	keyInv, _ := crypto.GenerateKey()
	keyFac, _ := crypto.GenerateKey()
	keyParam, _ := crypto.GenerateKey()
	keyData, _ := crypto.GenerateKey()
	keyMem1, _ := crypto.GenerateKey()
	keyMem2, _ := crypto.GenerateKey()
	authList := bind.NewKeyedTransactor(keyList)
	authInv := bind.NewKeyedTransactor(keyInv)
	authFac := bind.NewKeyedTransactor(keyFac)
	authParam := bind.NewKeyedTransactor(keyParam)
	authData := bind.NewKeyedTransactor(keyData)
	authMem1 := bind.NewKeyedTransactor(keyMem1)
	authMem2 := bind.NewKeyedTransactor(keyMem2)
	alloc := make(core.GenesisAlloc)
	alloc[authList.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authInv.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authFac.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authParam.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authData.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem1.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem2.From] = core.GenesisAccount{Balance: accountBalance}
	// 2nd arg is a gas limit, a uint64. we'll use 4.7 million
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &ctx{
		ListingAddress:    common.HexToAddress("0xlist"),
		InvestingAddress:  common.HexToAddress("0xinv"),
		AuthListing:       authList,
		AuthInvesting:     authInv,
		AuthFactory:       authFac,
		AuthParameterizer: authParam,
		AuthDatatrust:     authData,
		AuthMember1:       authMem1,
		AuthMember2:       authMem2, // a councilmember
		Blockchain:        bc,
	}
}
