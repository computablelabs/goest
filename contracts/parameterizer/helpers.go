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

const ONE_WEI = 1000000000000000000
const ONE_GWEI = 1000000000

// to match the p11r's internal params enum
const (
	UNDEFINED uint8 = iota
	CHALLENGE_STAKE
	CONVERSION_RATE
	CONVERSION_SLOPE_D
	CONVERSION_SLOPE_N
	LIST_REWARD
	QUORUM
	VOTE_BY
)

type ctx struct {
	// 	Alloc          core.GenesisAlloc
	AuthMarket  *bind.TransactOpts
	AuthFactory *bind.TransactOpts
	AuthMember1 *bind.TransactOpts
	AuthMember2 *bind.TransactOpts
	Blockchain  *backends.SimulatedBackend
}

// NOTE: there is no marketAddress present as we don't need an actual market here
type dep struct {
	VotingAddress            common.Address
	ParameterizerAddress     common.Address
	VotingContract           *voting.Voting
	ParameterizerContract    *Parameterizer
	VotingTransaction        *types.Transaction
	ParameterizerTransaction *types.Transaction
}

func Deploy(c *ctx) (*dep, error) {
	votingAddr, votingTrans, votingCont, votingErr := voting.DeployVoting(
		c.AuthFactory,
		c.Blockchain,
	)

	if votingErr != nil {
		return nil, votingErr
	}

	c.Blockchain.Commit()

	paramAddr, paramTrans, paramCont, paramErr := DeployParameterizer(
		c.AuthFactory,
		c.Blockchain,
		votingAddr,
		big.NewInt(ONE_WEI),  // challengeStake
		big.NewInt(ONE_GWEI), // conversionRate of 1e9
		big.NewInt(101),      // conversionSlopeDenominator, a scaling factor
		big.NewInt(100),      // conversionSlopeNumerator, a scaling factor
		big.NewInt(ONE_WEI),  // listReward
		big.NewInt(50),       // quorum
		big.NewInt(20),       // voteBy of 20 seconds for specs
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

// parameterizer address is passed in
func SetupBlockchain(accountBalance *big.Int) *ctx {
	// generate a new key, toss the error for now as it shouldnt happen
	keyMark, _ := crypto.GenerateKey()
	keyFac, _ := crypto.GenerateKey()
	keyMem1, _ := crypto.GenerateKey()
	keyMem2, _ := crypto.GenerateKey()
	authMark := bind.NewKeyedTransactor(keyMark)
	authFac := bind.NewKeyedTransactor(keyFac)
	authMem1 := bind.NewKeyedTransactor(keyMem1)
	authMem2 := bind.NewKeyedTransactor(keyMem2)
	alloc := make(core.GenesisAlloc)
	alloc[authMark.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authFac.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem1.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem2.From] = core.GenesisAccount{Balance: accountBalance}
	// 2nd arg is a gas limit, a uint64. we'll use 4.7 million
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &ctx{
		// 	Alloc:          alloc,
		AuthMarket:  authMark,
		AuthFactory: authFac,
		AuthMember1: authMem1,
		AuthMember2: authMem2,
		Blockchain:  bc,
	}
}
