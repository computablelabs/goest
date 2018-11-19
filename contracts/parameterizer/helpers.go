package parameterizer

import (
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
	VotingAddress            common.Address
	ParameterizerAddress     common.Address
	VotingContract           *voting.Voting
	ParameterizerContract    *Parameterizer
	VotingTransaction        *types.Transaction
	ParameterizerTransaction *types.Transaction
}

func Deploy(initialBalance *big.Int, c *ctx) (*dep, error) {
	votingAddr, votingTrans, votingCont, votingErr := voting.DeployVoting(
		c.AuthOwner,
		c.Blockchain,
	)

	if votingErr != nil {
		return nil, votingErr
	}

	c.Blockchain.Commit()

	paramAddr, paramTrans, paramCont, paramErr := DeployParameterizer(
		c.AuthOwner,
		c.Blockchain,
		votingAddr,
		big.NewInt(1000000000000000000), // challengeStake in tokenWei (10^18 == 1 token)
		big.NewInt(1),                   // conversionRate TODO tokenWei?
		big.NewInt(1),                   // conversionSlope TODO tokenWei?
		big.NewInt(1000000000000000000), // listReward (one token)
		big.NewInt(50),                  // quorum
		big.NewInt(300),                 // voteBy (5 mins)
	)

	if paramErr != nil {
		return nil, paramErr
	}

	c.Blockchain.Commit()

	return &dep{
		VotingAddress:            votingAddr,
		VotingContract:           votingCont,
		VotingTransaction:        votingTrans,
		ParameterizerAddress:     paramAddr,
		ParameterizerContract:    paramCont,
		ParameterizerTransaction: paramTrans,
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
	// 2nd arg is a gas limit, a uint64. we'll use 4.7 million
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &ctx{
		AuthOwner:      authOwner,
		AuthChallenger: authChallenger,
		AuthVoter:      authVoter,
		Blockchain:     bc,
	}
}
