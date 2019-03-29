package parameterizer

import (
	"github.com/computablelabs/goest/contracts/markettoken"
	"github.com/computablelabs/goest/contracts/voting"
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

// to match the p11r's internal params
var STAKE = big.NewInt(1)
var CONVERSION_RATE = big.NewInt(2)
var INVEST_D = big.NewInt(3)
var INVEST_N = big.NewInt(4)
var LIST_REWARD = big.NewInt(5)
var QUORUM = big.NewInt(6)
var VOTE_BY = big.NewInt(7)

type ctx struct {
	// 	Alloc          core.GenesisAlloc
	AuthInvesting *bind.TransactOpts
	AuthListing   *bind.TransactOpts
	AuthDatatrust *bind.TransactOpts
	AuthFactory   *bind.TransactOpts
	AuthMember1   *bind.TransactOpts
	AuthMember2   *bind.TransactOpts
	AuthMember3   *bind.TransactOpts
	Blockchain    *backends.SimulatedBackend
}

type dep struct {
	MarketTokenAddress       common.Address
	MarketTokenContract      *markettoken.MarketToken
	MarketTokenTransaction   *types.Transaction
	VotingAddress            common.Address
	ParameterizerAddress     common.Address
	VotingContract           *voting.Voting
	ParameterizerContract    *Parameterizer
	VotingTransaction        *types.Transaction
	ParameterizerTransaction *types.Transaction
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

	c.Blockchain.Commit()

	votingAddr, votingTrans, votingCont, votingErr := voting.DeployVoting(
		c.AuthFactory,
		c.Blockchain,
		marketTokenAddr,
	)

	if votingErr != nil {
		return nil, votingErr
	}

	c.Blockchain.Commit()

	paramAddr, paramTrans, paramCont, paramErr := DeployParameterizer(
		c.AuthFactory,
		c.Blockchain,
		votingAddr,
		big.NewInt(ONE_GWEI),  // conversionRate of 1e9
		big.NewInt(100),       // InvestDenominator, a scaling factor
		big.NewInt(110),       // InvestNumerator, a scaling factor
		big.NewInt(ONE_WEI),   // listReward
		big.NewInt(ONE_WEI/2), // access reward
		big.NewInt(ONE_GWEI),  // stake
		big.NewInt(40),        // voteBy of 20 seconds for specs
		big.NewInt(50),        // quorum
		big.NewInt(30),        // backend payment percent
		big.NewInt(50),        // maker payment percent
		big.NewInt(ONE_WEI/4), // cost per byte
	)

	if paramErr != nil {
		return nil, paramErr
	}

	c.Blockchain.Commit()

	return &dep{
		MarketTokenAddress:       marketTokenAddr,
		MarketTokenContract:      marketTokenCont,
		MarketTokenTransaction:   marketTokenTrans,
		VotingAddress:            votingAddr,
		VotingContract:           votingCont,
		VotingTransaction:        votingTrans,
		ParameterizerAddress:     paramAddr,
		ParameterizerContract:    paramCont,
		ParameterizerTransaction: paramTrans,
	}, nil
}

// parameterizer address is passed in
func SetupBlockchain(accountBalance *big.Int) *ctx {
	// generate a new key, toss the error for now as it shouldnt happen
	keyData, _ := crypto.GenerateKey()
	keyList, _ := crypto.GenerateKey()
	keyInv, _ := crypto.GenerateKey()
	keyFac, _ := crypto.GenerateKey()
	keyMem1, _ := crypto.GenerateKey()
	keyMem2, _ := crypto.GenerateKey()
	keyMem3, _ := crypto.GenerateKey()
	authData := bind.NewKeyedTransactor(keyData)
	authList := bind.NewKeyedTransactor(keyList)
	authInv := bind.NewKeyedTransactor(keyInv)
	authFac := bind.NewKeyedTransactor(keyFac)
	authMem1 := bind.NewKeyedTransactor(keyMem1)
	authMem2 := bind.NewKeyedTransactor(keyMem2)
	authMem3 := bind.NewKeyedTransactor(keyMem3)
	alloc := make(core.GenesisAlloc)
	alloc[authData.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authList.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authInv.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authFac.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem1.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem2.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem3.From] = core.GenesisAccount{Balance: accountBalance}
	// 2nd arg is a gas limit, a uint64. we'll use 4.7 million
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &ctx{
		// 	Alloc:          alloc,
		AuthInvesting: authInv,
		AuthListing:   authList,
		AuthDatatrust: authData,
		AuthFactory:   authFac,
		AuthMember1:   authMem1,
		AuthMember2:   authMem2,
		AuthMember3:   authMem3,
		Blockchain:    bc,
	}
}
