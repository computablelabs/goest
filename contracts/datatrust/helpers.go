package datatrust

import (
	"github.com/computablelabs/goest/contracts/ethertoken"
	"github.com/computablelabs/goest/contracts/markettoken"
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
	InvestingAddress common.Address
	AuthFactory      *bind.TransactOpts
	AuthListing      *bind.TransactOpts
	AuthBackend      *bind.TransactOpts
	AuthMember1      *bind.TransactOpts
	AuthMember2      *bind.TransactOpts
	AuthMember3      *bind.TransactOpts
	Blockchain       *backends.SimulatedBackend
}

type dep struct {
	EtherTokenAddress        common.Address
	MarketTokenAddress       common.Address
	DatatrustAddress         common.Address
	ParameterizerAddress     common.Address
	VotingAddress            common.Address
	EtherTokenContract       *ethertoken.EtherToken
	MarketTokenContract      *markettoken.MarketToken
	DatatrustContract        *Datatrust
	ParameterizerContract    *parameterizer.Parameterizer
	VotingContract           *voting.Voting
	EtherTokenTransaction    *types.Transaction
	MarketTokenTransaction   *types.Transaction
	DatatrustTransaction     *types.Transaction
	ParameterizerTransaction *types.Transaction
	VotingTransaction        *types.Transaction
}

func Deploy(initialBalance *big.Int, c *ctx) (*dep, error) {
	etherTokenAddr, etherTokenTrans, etherTokenCont, etherTokenErr := ethertoken.DeployEtherToken(
		c.AuthFactory,
		c.Blockchain,
		c.AuthFactory.From,
		initialBalance,
	)

	if etherTokenErr != nil {
		return nil, etherTokenErr
	}

	c.Blockchain.Commit()

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

	paramAddr, paramTrans, paramCont, paramErr := parameterizer.DeployParameterizer(
		c.AuthFactory,
		c.Blockchain,
		votingAddr,
		big.NewInt(ONE_GWEI), // conversionRate: stipulation is that market token should be, at least, as val as eth
		big.NewInt(100),      // investDenominator, a scaling factor
		big.NewInt(110),      // investNumerator, a scaling factor
		big.NewInt(ONE_WEI),  // listReward (one token)
		big.NewInt(ONE_GWEI), // access reward
		big.NewInt(ONE_GWEI), // Stake
		big.NewInt(100),      // voteBy
		big.NewInt(50),       // quorum
		big.NewInt(20),       // backend payment percent
		big.NewInt(60),       // maker payment percent
		big.NewInt(1000),     // cost per byte (1000 wei per byte)
	)

	if paramErr != nil {
		return nil, paramErr
	}

	c.Blockchain.Commit()

	// finally the datatrust...
	dataAddr, dataTrans, dataCont, dataErr := DeployDatatrust(
		c.AuthFactory,
		c.Blockchain,
		etherTokenAddr,
		votingAddr,
		paramAddr,
	)

	if dataErr != nil {
		return nil, dataErr
	}

	c.Blockchain.Commit()

	return &dep{
		EtherTokenAddress:        etherTokenAddr,
		EtherTokenContract:       etherTokenCont,
		EtherTokenTransaction:    etherTokenTrans,
		MarketTokenAddress:       marketTokenAddr,
		MarketTokenContract:      marketTokenCont,
		MarketTokenTransaction:   marketTokenTrans,
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
	keyLis, _ := crypto.GenerateKey()
	keyMem1, _ := crypto.GenerateKey()
	keyMem2, _ := crypto.GenerateKey()
	keyMem3, _ := crypto.GenerateKey()

	authFac := bind.NewKeyedTransactor(keyFac)
	authBac := bind.NewKeyedTransactor(keyBac)
	authLis := bind.NewKeyedTransactor(keyLis)
	authFac.GasPrice = big.NewInt(ONE_GWEI * 2)
	// authFac.GasLimit = 4000000             // setting a gas limit here causes the "silent simulated backed fail"... TODO PR

	authMem1 := bind.NewKeyedTransactor(keyMem1)
	authMem2 := bind.NewKeyedTransactor(keyMem2)
	authMem3 := bind.NewKeyedTransactor(keyMem3)

	alloc := make(core.GenesisAlloc)
	alloc[authFac.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authBac.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authLis.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem1.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem2.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem3.From] = core.GenesisAccount{Balance: accountBalance}
	// 2nd arg is a gas limit, a uint64. we'll use 4.7M
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &ctx{
		InvestingAddress: common.HexToAddress("0xinv"),
		AuthFactory:      authFac,
		AuthBackend:      authBac,
		AuthListing:      authLis,
		AuthMember1:      authMem1,
		AuthMember2:      authMem2,
		AuthMember3:      authMem3,
		Blockchain:       bc,
	}
}