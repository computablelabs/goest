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

const ONE_WEI = 1000000000000000000
const ONE_GWEI = 1000000000

// the candidate "kinds"
const (
	UNDEFINED uint8 = iota
	APPLICATION
	CHALLENGE
	REPARAM
)

type ctx struct {
	AuthFactory *bind.TransactOpts
	AuthMember1 *bind.TransactOpts
	AuthMember2 *bind.TransactOpts
	AuthMember3 *bind.TransactOpts
	Blockchain  *backends.SimulatedBackend
}

type dep struct {
	MarketAddress            common.Address
	MarketTokenAddress       common.Address
	NetworkTokenAddress      common.Address
	ParameterizerAddress     common.Address
	VotingAddress            common.Address
	MarketContract           *Market
	MarketTokenContract      *markettoken.MarketToken
	NetworkTokenContract     *networktoken.NetworkToken
	ParameterizerContract    *parameterizer.Parameterizer
	VotingContract           *voting.Voting
	MarketTransaction        *types.Transaction
	MarketTokenTransaction   *types.Transaction
	NetworkTokenTransaction  *types.Transaction
	ParameterizerTransaction *types.Transaction
	VotingTransaction        *types.Transaction
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

	networkTokenAddr, networkTokenTrans, networkTokenCont, networkTokenErr := networktoken.DeployNetworkToken(
		c.AuthFactory,
		c.Blockchain,
		c.AuthFactory.From,
		initialBalance,
	)

	if networkTokenErr != nil {
		return nil, networkTokenErr
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
		big.NewInt(ONE_WEI),     // challengeStake
		big.NewInt(ONE_WEI*0.1), // conversionRate tokenWei, .1 of a token (10**16)
		big.NewInt(101),         // conversionSlopeDenominator, a scaling factor
		big.NewInt(100),         // conversionSlopeNumerator, a scaling factor
		big.NewInt(ONE_WEI),     // listReward (one token)
		big.NewInt(50),          // quorum
		big.NewInt(100),         // voteBy
	)

	if paramErr != nil {
		return nil, paramErr
	}

	c.Blockchain.Commit()

	// finally the market...
	marketAddr, marketTrans, marketCont, marketErr := DeployMarket(
		c.AuthFactory,
		c.Blockchain,
		marketTokenAddr,
		networkTokenAddr,
		paramAddr,
		votingAddr,
	)

	if marketErr != nil {
		return nil, marketErr
	}

	c.Blockchain.Commit()

	return &dep{
		MarketAddress:            marketAddr,
		MarketContract:           marketCont,
		MarketTransaction:        marketTrans,
		MarketTokenAddress:       marketTokenAddr,
		MarketTokenContract:      marketTokenCont,
		MarketTokenTransaction:   marketTokenTrans,
		NetworkTokenAddress:      networkTokenAddr,
		NetworkTokenTransaction:  networkTokenTrans,
		NetworkTokenContract:     networkTokenCont,
		ParameterizerAddress:     paramAddr,
		ParameterizerContract:    paramCont,
		ParameterizerTransaction: paramTrans,
		VotingAddress:            votingAddr,
		VotingContract:           votingCont,
		VotingTransaction:        votingTrans,
	}, nil
}

func SetupBlockchain(accountBalance *big.Int) *ctx {
	// generate a new key, toss the error for now as it shouldnt happen
	keyFac, _ := crypto.GenerateKey()
	keyMem1, _ := crypto.GenerateKey()
	keyMem2, _ := crypto.GenerateKey()
	keyMem3, _ := crypto.GenerateKey()

	authFac := bind.NewKeyedTransactor(keyFac)
	authFac.GasPrice = big.NewInt(ONE_GWEI * 2)
	// authFac.GasLimit = 4000000             // setting a gas limit here causes the "silent simulated backed fail"... TODO PR

	// members 1 and 2 used in list_functionality
	authMem1 := bind.NewKeyedTransactor(keyMem1)
	authMem2 := bind.NewKeyedTransactor(keyMem2)
	// member 3 used in invest_functionality
	authMem3 := bind.NewKeyedTransactor(keyMem3)

	alloc := make(core.GenesisAlloc)
	alloc[authFac.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem1.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem2.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem3.From] = core.GenesisAccount{Balance: accountBalance}
	// 2nd arg is a gas limit, a uint64. we'll use 4.7M
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &ctx{
		AuthFactory: authFac,
		AuthMember1: authMem1,
		AuthMember2: authMem2,
		AuthMember3: authMem3,
		Blockchain:  bc,
	}
}
