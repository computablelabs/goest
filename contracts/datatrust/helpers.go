package datatrust

import (
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
	REGISTRATION = 4
)

type ctx struct {
	ListingAddress   common.Address
	InvestingAddress common.Address
	AuthFactory      *bind.TransactOpts
	AuthBackend      *bind.TransactOpts
	AuthMember1      *bind.TransactOpts
	AuthMember2      *bind.TransactOpts
	AuthMember3      *bind.TransactOpts
	Blockchain       *backends.SimulatedBackend
}

type dep struct {
	DatatrustAddress         common.Address
	ParameterizerAddress     common.Address
	VotingAddress            common.Address
	DatatrustContract        *Datatrust
	ParameterizerContract    *parameterizer.Parameterizer
	VotingContract           *voting.Voting
	DatatrustTransaction     *types.Transaction
	ParameterizerTransaction *types.Transaction
	VotingTransaction        *types.Transaction
}

func Deploy(initialBalance *big.Int, c *ctx) (*dep, error) {
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
		big.NewInt(ONE_WEI),  // challengeStake
		big.NewInt(ONE_GWEI), // conversionRate: stipulation is that market token should be, at least, as val as eth
		big.NewInt(100),      // investDenominator, a scaling factor
		big.NewInt(110),      // investNumerator, a scaling factor
		big.NewInt(ONE_WEI),  // listReward (one token)
		big.NewInt(50),       // quorum
		big.NewInt(100),      // voteBy
	)

	if paramErr != nil {
		return nil, paramErr
	}

	c.Blockchain.Commit()

	// finally the datatrust...
	dataAddr, dataTrans, dataCont, dataErr := DeployDatatrust(
		c.AuthFactory,
		c.Blockchain,
		votingAddr,
		paramAddr,
	)

	if dataErr != nil {
		return nil, dataErr
	}

	c.Blockchain.Commit()

	return &dep{
		DatatrustAddress:         dataAddr,
		DatatrustContract:        dataCont,
		DatatrustTransaction:     dataTrans,
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
	keyBac, _ := crypto.GenerateKey()
	keyMem1, _ := crypto.GenerateKey()
	keyMem2, _ := crypto.GenerateKey()
	keyMem3, _ := crypto.GenerateKey()

	authFac := bind.NewKeyedTransactor(keyFac)
	authBac := bind.NewKeyedTransactor(keyBac)
	authFac.GasPrice = big.NewInt(ONE_GWEI * 2)
	// authFac.GasLimit = 4000000             // setting a gas limit here causes the "silent simulated backed fail"... TODO PR

	authMem1 := bind.NewKeyedTransactor(keyMem1)
	authMem2 := bind.NewKeyedTransactor(keyMem2)
	authMem3 := bind.NewKeyedTransactor(keyMem3)

	alloc := make(core.GenesisAlloc)
	alloc[authFac.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authBac.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem1.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem2.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem3.From] = core.GenesisAccount{Balance: accountBalance}
	// 2nd arg is a gas limit, a uint64. we'll use 4.7M
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &ctx{
		ListingAddress:   common.HexToAddress("0xlist"),
		InvestingAddress: common.HexToAddress("0xinv"),
		AuthFactory:      authFac,
		AuthBackend:      authBac,
		AuthMember1:      authMem1,
		AuthMember2:      authMem2,
		AuthMember3:      authMem3,
		Blockchain:       bc,
	}
}
