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
	AuthOwner      *bind.TransactOpts
	AuthVoter      *bind.TransactOpts
	AuthChallenger *bind.TransactOpts
	Blockchain     *backends.SimulatedBackend
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
		c.AuthOwner,
		c.Blockchain,
		c.AuthOwner.From,
		initialBalance,
	)

	if networkTokenErr != nil {
		return nil, networkTokenErr
	}

	marketTokenAddr, marketTokenTrans, marketTokenCont, marketTokenErr := markettoken.DeployMarketToken(
		c.AuthOwner,
		c.Blockchain,
		c.AuthOwner.From,
		initialBalance,
	)

	if marketTokenErr != nil {
		return nil, marketTokenErr
	}

	// commit the deploy before deploying voting
	c.Blockchain.Commit()

	votingAddr, votingTrans, votingCont, votingErr := voting.DeployVoting(
		c.AuthOwner,
		c.Blockchain,
	)

	if votingErr != nil {
		return nil, votingErr
	}

	c.Blockchain.Commit()

	paramAddr, paramTrans, paramCont, paramErr := parameterizer.DeployParameterizer(
		c.AuthOwner,
		c.Blockchain,
		votingAddr,
		big.NewInt(1000000000000000000), // challengeStake in tokenWei (10**18 == 1 token)
		big.NewInt(10000000000000000),   // conversionRate tokenWei, .1 of a token (10**16)
		big.NewInt(101),                 // conversionSlopeDenominator, a scaling factor
		big.NewInt(100),                 // conversionSlopeNumerator, a scaling factor
		big.NewInt(1000000000000000000), // listReward (one token)
		big.NewInt(50),                  // quorum
		big.NewInt(300),                 // voteBy (5 mins)
	)

	if paramErr != nil {
		return nil, paramErr
	}

	c.Blockchain.Commit()

	// finally the market...
	marketAddr, marketTrans, marketCont, marketErr := DeployMarket(
		c.AuthOwner,
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
	keyOwner, _ := crypto.GenerateKey()
	keyChallenger, _ := crypto.GenerateKey()
	keyVoter, _ := crypto.GenerateKey()

	authOwner := bind.NewKeyedTransactor(keyOwner)
	authOwner.GasPrice = big.NewInt(2000000000) // 2 Gwei
	authOwner.GasLimit = 1000000

	authChallenger := bind.NewKeyedTransactor(keyChallenger)
	authVoter := bind.NewKeyedTransactor(keyVoter)
	alloc := make(core.GenesisAlloc)
	alloc[authOwner.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authChallenger.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authVoter.From] = core.GenesisAccount{Balance: accountBalance}
	// 2nd arg is a gas limit, a uint64. we'll use 4.7M
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &ctx{
		AuthOwner:      authOwner,
		AuthChallenger: authChallenger,
		AuthVoter:      authVoter,
		Blockchain:     bc,
	}
}
