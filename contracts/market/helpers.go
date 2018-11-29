package market

import (
	"github.com/computablelabs/goest/contracts/markettoken"
	"github.com/computablelabs/goest/contracts/networktoken"
	"github.com/computablelabs/goest/contracts/parameterizer"
	"github.com/computablelabs/goest/contracts/voting"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type ctx struct {
	AuthFactory *bind.TransactOpts
	AuthMember1 *bind.TransactOpts
	AuthMember2 *bind.TransactOpts
	Blockchain  *backends.SimulatedBackend
}

type dep struct {
	NetworkTokenAddress      common.Address
	MarketTokenAddress       common.Address
	VotingAddress            common.Address
	ParameterizerAddress     common.Address
	MarketAddress            common.Address
	NetworkTokenContract     *networktoken.NetworkToken
	MarketTokenContract      *markettoken.MarketToken
	VotingContract           *voting.Voting
	ParameterizerContract    *parameterizer.Parameterizer
	MarketContract           *Market
	NetworkTokenTransaction  *types.Transaction
	MarketTokenTransaction   *types.Transaction
	VotingTransaction        *types.Transaction
	ParameterizerTransaction *types.Transaction
	MarketTransaction        *types.Transaction
}

func Deploy(initialBalance *big.Int, c *ctx) (*dep, error) {
	networkTokenAddr, networkTokenTrans, networkTokenCont, networkTokenErr := networktoken.DeployNetworkToken(
		c.AuthFactory,
		c.Blockchain,
		c.AuthFactory.From,
		initialBalance,
	)

	if networkTokenErr != nil {
		return nil, networkTokenErr
	}

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

	votingAddr, votingTrans, votingCont, votingErr := voting.DeployVoting(
		c.AuthFactory,
		c.Blockchain,
	)

	if votingErr != nil {
		return nil, votingErr
	}

	c.Blockchain.Commit()

	paramAddr, paramTrans, paramCont, paramErr := parameterizer.DeployParameterizer(
		c.AuthFactory,
		c.Blockchain,
		votingAddr,
		big.NewInt(1000000000000000000), // challengeStake in tokenWei (10**18 == 1 token)
		big.NewInt(10000000000000000),   // conversionRate tokenWei, .1 of a token (10**16)
		big.NewInt(101),                 // conversionSlopeDenominator, a scaling factor
		big.NewInt(100),                 // conversionSlopeNumerator, a scaling factor
		big.NewInt(1000000000000000000), // listReward (one token)
		big.NewInt(50),                  // quorum
		big.NewInt(20),                  // voteBy (20 seconds)
	)

	if paramErr != nil {
		return nil, paramErr
	}

	c.Blockchain.Commit()

	// finally the market...
	marketAddr, marketTrans, marketCont, marketErr := DeployMarket(
		c.AuthFactory,
		c.Blockchain,
		"FooMarket",
		networkTokenAddr,
		marketTokenAddr,
		votingAddr,
		paramAddr,
	)

	if marketErr != nil {
		return nil, marketErr
	}

	c.Blockchain.Commit()

	return &dep{
		NetworkTokenAddress:      networkTokenAddr,
		NetworkTokenTransaction:  networkTokenTrans,
		NetworkTokenContract:     networkTokenCont,
		MarketTokenAddress:       marketTokenAddr,
		MarketTokenContract:      marketTokenCont,
		MarketTokenTransaction:   marketTokenTrans,
		VotingAddress:            votingAddr,
		VotingContract:           votingCont,
		VotingTransaction:        votingTrans,
		ParameterizerAddress:     paramAddr,
		ParameterizerContract:    paramCont,
		ParameterizerTransaction: paramTrans,
		MarketAddress:            marketAddr,
		MarketContract:           marketCont,
		MarketTransaction:        marketTrans,
	}, nil
}

func SetupBlockchain(accountBalance *big.Int) *ctx {
	// generate a new key, toss the error for now as it shouldnt happen
	keyFac, _ := crypto.GenerateKey()
	keyMem1, _ := crypto.GenerateKey()
	keyMem2, _ := crypto.GenerateKey()

	authFac := bind.NewKeyedTransactor(keyFac)
	authFac.GasPrice = big.NewInt(2000000000) // 2 Gwei
	authFac.GasLimit = 1000000

	authMem1 := bind.NewKeyedTransactor(keyMem1)
	authMem2 := bind.NewKeyedTransactor(keyMem2)
	alloc := make(core.GenesisAlloc)
	alloc[authFac.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem1.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem2.From] = core.GenesisAccount{Balance: accountBalance}
	// 2nd arg is a gas limit, a uint64. we'll use 4.7M
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &ctx{
		AuthFactory: authFac,
		AuthMember1: authMem1,
		AuthMember2: authMem2,
		Blockchain:  bc,
	}
}
