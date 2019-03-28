package listing

import (
	"github.com/computablelabs/goest/contracts/datatrust"
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
	UNDEFINED uint8 = iota
	APPLICATION
	CHALLENGE
	REPARAM
)

type ctx struct {
	EtherTokenAddress common.Address
	AuthBackend       *bind.TransactOpts
	AuthFactory       *bind.TransactOpts
	AuthInvest        *bind.TransactOpts
	AuthMember1       *bind.TransactOpts
	AuthMember2       *bind.TransactOpts
	AuthMember3       *bind.TransactOpts
	Blockchain        *backends.SimulatedBackend
}

type dep struct {
	ListingAddress           common.Address
	MarketTokenAddress       common.Address
	ParameterizerAddress     common.Address
	VotingAddress            common.Address
	DatatrustAddress         common.Address
	ListingContract          *Listing
	MarketTokenContract      *markettoken.MarketToken
	ParameterizerContract    *parameterizer.Parameterizer
	VotingContract           *voting.Voting
	DatatrustContract        *datatrust.Datatrust
	ListingTransaction       *types.Transaction
	MarketTokenTransaction   *types.Transaction
	ParameterizerTransaction *types.Transaction
	VotingTransaction        *types.Transaction
	DatatrustTransaction     *types.Transaction
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

	// commit the deploy before deploying voting
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
		big.NewInt(ONE_GWEI),  // conversionRate: stipulation is that market token should be, at least, as val as eth
		big.NewInt(100),       // investDenominator, a scaling factor
		big.NewInt(110),       // investNumerator, a scaling factor
		big.NewInt(ONE_WEI),   // listReward (one token)
		big.NewInt(ONE_WEI/2), // access reward
		big.NewInt(ONE_GWEI),  // Stake
		big.NewInt(100),       // voteBy
		big.NewInt(50),        // quorum
		big.NewInt(30),        // backend payment percent
		big.NewInt(50),        // maker payment percent
		big.NewInt(ONE_WEI/4), // cost per byte
	)

	if paramErr != nil {
		return nil, paramErr
	}

	c.Blockchain.Commit()

	// the datatrust
	dataAddr, dataTrans, dataCont, dataErr := datatrust.DeployDatatrust(
		c.AuthFactory,
		c.Blockchain,
		c.EtherTokenAddress,
		votingAddr,
		paramAddr,
	)

	if dataErr != nil {
		return nil, dataErr
	}

	c.Blockchain.Commit()

	// finally the listing...
	listingAddr, listingTrans, listingCont, listingErr := DeployListing(
		c.AuthFactory,
		c.Blockchain,
		marketTokenAddr,
		votingAddr,
		paramAddr,
		dataAddr,
	)

	if listingErr != nil {
		return nil, listingErr
	}

	c.Blockchain.Commit()

	return &dep{
		ListingAddress:           listingAddr,
		ListingContract:          listingCont,
		ListingTransaction:       listingTrans,
		MarketTokenAddress:       marketTokenAddr,
		MarketTokenContract:      marketTokenCont,
		MarketTokenTransaction:   marketTokenTrans,
		ParameterizerAddress:     paramAddr,
		ParameterizerContract:    paramCont,
		ParameterizerTransaction: paramTrans,
		VotingAddress:            votingAddr,
		VotingContract:           votingCont,
		VotingTransaction:        votingTrans,
		DatatrustAddress:         dataAddr,
		DatatrustContract:        dataCont,
		DatatrustTransaction:     dataTrans,
	}, nil
}

func SetupBlockchain(accountBalance *big.Int) *ctx {
	// generate a new key, toss the error for now as it shouldnt happen
	keyBac, _ := crypto.GenerateKey()
	keyFac, _ := crypto.GenerateKey()
	keyInv, _ := crypto.GenerateKey()
	keyMem1, _ := crypto.GenerateKey()
	keyMem2, _ := crypto.GenerateKey()
	keyMem3, _ := crypto.GenerateKey()

	authBac := bind.NewKeyedTransactor(keyBac)
	authFac := bind.NewKeyedTransactor(keyFac)
	authInv := bind.NewKeyedTransactor(keyInv)
	authFac.GasPrice = big.NewInt(ONE_GWEI * 2)
	// authFac.GasLimit = 4000000             // setting a gas limit here causes the "silent simulated backed fail"... TODO PR

	// members 1 and 2 used in list_functionality
	authMem1 := bind.NewKeyedTransactor(keyMem1)
	authMem2 := bind.NewKeyedTransactor(keyMem2)
	authMem3 := bind.NewKeyedTransactor(keyMem3)

	alloc := make(core.GenesisAlloc)
	alloc[authBac.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authFac.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authInv.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem1.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem2.From] = core.GenesisAccount{Balance: accountBalance}
	alloc[authMem3.From] = core.GenesisAccount{Balance: accountBalance}
	// 2nd arg is a gas limit, a uint64. we'll use 4.7M
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &ctx{
		EtherTokenAddress: common.HexToAddress("0xether"),
		AuthBackend:       authBac,
		AuthFactory:       authFac,
		AuthInvest:        authInv,
		AuthMember1:       authMem1,
		AuthMember2:       authMem2,
		AuthMember3:       authMem3,
		Blockchain:        bc,
	}
}
