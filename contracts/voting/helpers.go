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
	AuthFactory       *bind.TransactOpts
	AuthParameterizer *bind.TransactOpts
	AuthMember1       *bind.TransactOpts
	AuthMember2       *bind.TransactOpts
	Blockchain        *backends.SimulatedBackend
}

type dep struct {
	VotingAddress     common.Address
	VotingContract    *Voting
	VotingTransaction *types.Transaction
}

func Deploy(c *ctx) (*dep, error) {
	votingAddr, votingTrans, votingCont, votingErr := DeployVoting(
		c.AuthFactory,
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
	keyMark, _ := crypto.GenerateKey()
	keyFac, _ := crypto.GenerateKey()
	keyParam, _ := crypto.GenerateKey()
	keyMem1, _ := crypto.GenerateKey()
	keyMem2, _ := crypto.GenerateKey()
	authMark := bind.NewKeyedTransactor(keyMark)
	authFac := bind.NewKeyedTransactor(keyFac)
	authParam := bind.NewKeyedTransactor(keyParam)
	authMem1 := bind.NewKeyedTransactor(keyMem1)
	authMem2 := bind.NewKeyedTransactor(keyMem2)
	alloc := make(core.GenesisAlloc)
	alloc[authMark.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authFac.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authParam.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem1.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem2.From] = core.GenesisAccount{Balance: accountBalance}
	// 2nd arg is a gas limit, a uint64. we'll use 4.7 million
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &ctx{
		AuthMarket:        authMark,
		AuthFactory:       authFac,
		AuthParameterizer: authParam,
		AuthMember1:       authMem1,
		AuthMember2:       authMem2, // a councilmember
		Blockchain:        bc,
	}
}
