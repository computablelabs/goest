package voting

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
	AuthMarket        *bind.TransactOpts
	AuthOwner         *bind.TransactOpts
	AuthParameterizer *bind.TransactOpts
	AuthVoter         *bind.TransactOpts
	AuthChallenger    *bind.TransactOpts
	Blockchain        *backends.SimulatedBackend
}

type dep struct {
	VotingAddress     common.Address
	VotingContract    *Voting
	VotingTransaction *types.Transaction
}

func Deploy(initialBalance *big.Int, c *ctx) (*dep, error) {
	votingAddr, votingTrans, votingCont, votingErr := DeployVoting(
		c.AuthOwner,
		c.Blockchain,
	)

	if votingErr != nil {
		return nil, votingErr
	}

	c.Blockchain.Commit()

	return &dep{
		VotingAddress:     votingAddr,
		VotingContract:    votingCont,
		VotingTransaction: votingTrans,
	}, nil
}

// helper for creating the solidity bytes32 listing hash
func genBytes32(name string) [32]byte {
	bytes := [32]byte{}
	copy(bytes[:], []byte(name))
	return bytes
}

func SetupBlockchain(accountBalance *big.Int) *ctx {
	// generate a new key, toss the error for now as it shouldnt happen
	keyMarket, _ := crypto.GenerateKey()
	keyOwner, _ := crypto.GenerateKey()
	keyParameterizer, _ := crypto.GenerateKey()
	keyChallenger, _ := crypto.GenerateKey()
	keyVoter, _ := crypto.GenerateKey()
	authMarket := bind.NewKeyedTransactor(keyMarket)
	authOwner := bind.NewKeyedTransactor(keyOwner)
	authParameterizer := bind.NewKeyedTransactor(keyParameterizer)
	authChallenger := bind.NewKeyedTransactor(keyChallenger)
	authVoter := bind.NewKeyedTransactor(keyVoter)
	alloc := make(core.GenesisAlloc)
	alloc[authMarket.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authOwner.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authParameterizer.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authChallenger.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authVoter.From] = core.GenesisAccount{Balance: accountBalance}
	// 2nd arg is a gas limit, a uint64. we'll use 4.7 million
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &ctx{
		AuthMarket:        authMarket,
		AuthOwner:         authOwner,
		AuthParameterizer: authParameterizer,
		AuthChallenger:    authChallenger,
		AuthVoter:         authVoter, // a councilmember
		Blockchain:        bc,
	}
}
