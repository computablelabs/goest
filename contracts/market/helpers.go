package market

import (
	"github.com/computablelabs/goest/contracts/markettoken"
	"github.com/computablelabs/goest/contracts/parameterizer"
	"github.com/computablelabs/goest/contracts/plcrvoting"
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
	TokenAddress             common.Address
	VotingAddress            common.Address
	ParameterizerAddress     common.Address
	MarketAddress            common.Address
	TokenContract            *markettoken.MarketToken
	VotingContract           *plcrvoting.PLCRVoting
	ParameterizerContract    *parameterizer.Parameterizer
	MarketContract           *Market
	TokenTransaction         *types.Transaction // TODO we may not ever ref these TXs: possibly remove
	VotingTransaction        *types.Transaction
	ParameterizerTransaction *types.Transaction
	MarketTransaction        *types.Transaction
}

func Deploy(initialBalance *big.Int, c *ctx) (*dep, error) {
	tokenAddr, tokenTrans, tokenCont, tokenErr := markettoken.DeployMarketToken(
		c.AuthOwner,
		c.Blockchain,
		c.AuthOwner.From,
		initialBalance,
	)

	if tokenErr != nil {
		return nil, tokenErr
	}

	// commit the deploy before deploying voting
	c.Blockchain.Commit()

	votingAddr, votingTrans, votingCont, votingErr := plcrvoting.DeployPLCRVoting(
		c.AuthOwner,
		c.Blockchain,
		tokenAddr,
	)

	if votingErr != nil {
		return nil, votingErr
	}

	c.Blockchain.Commit()

	paramAddr, paramTrans, paramCont, paramErr := parameterizer.DeployParameterizer(
		c.AuthOwner,
		c.Blockchain,
		tokenAddr,
		votingAddr,
		big.NewInt(10), // minDeposit
		big.NewInt(60), // applyStageLen in sec
		big.NewInt(60), // commitStageLen in sec
		big.NewInt(60), // revealStageLen in sec
		big.NewInt(50), // dispensation pct
		big.NewInt(50), // voteQuorum
		big.NewInt(1),  // listReward
	)

	if paramErr != nil {
		return nil, paramErr
	}

	c.Blockchain.Commit()

	// finally the market...
	marketAddr, marketTrans, marketCont, marketErr := DeployMarket(
		c.AuthOwner,
		c.Blockchain,
		tokenAddr, // TODO will need to become market token specific once network token is in
		votingAddr,
		paramAddr,
		"FooMarket",
	)

	if marketErr != nil {
		return nil, marketErr
	}

	c.Blockchain.Commit()

	return &dep{
		TokenAddress:             tokenAddr,
		TokenContract:            tokenCont,
		TokenTransaction:         tokenTrans,
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
	authChallenger := bind.NewKeyedTransactor(keyChallenger)
	authVoter := bind.NewKeyedTransactor(keyVoter)
	alloc := make(core.GenesisAlloc)
	alloc[authOwner.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authChallenger.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authVoter.From] = core.GenesisAccount{Balance: accountBalance}
	// 2nd arg is a gas limit, a uint64. we'll use 7 million
	bc := backends.NewSimulatedBackend(alloc, 7000000)

	return &ctx{
		AuthOwner:      authOwner,
		AuthChallenger: authChallenger,
		AuthVoter:      authVoter,
		Blockchain:     bc,
	}
}
