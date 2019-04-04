package test

import (
	"github.com/computablelabs/goest/contracts/datatrust"
	"github.com/computablelabs/goest/contracts/ethertoken"
	"github.com/computablelabs/goest/contracts/investing"
	"github.com/computablelabs/goest/contracts/listing"
	"github.com/computablelabs/goest/contracts/markettoken"
	"github.com/computablelabs/goest/contracts/parameterizer"
	"github.com/computablelabs/goest/contracts/voting"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// The deployed struct holds references to the instantiated contracts on the
// simulated backend. Also, refernces to each contract's adress and the
// deployment transaction itself, in case we want to inspect it.
type Dep struct {
	EtherTokenAddress        common.Address
	EtherTokenContract       *ethertoken.EtherToken
	EtherTokenTransaction    *types.Transaction
	MarketTokenAddress       common.Address
	MarketTokenContract      *markettoken.MarketToken
	MarketTokenTransaction   *types.Transaction
	VotingAddress            common.Address
	VotingContract           *voting.Voting
	VotingTransaction        *types.Transaction
	ParameterizerAddress     common.Address
	ParameterizerContract    *parameterizer.Parameterizer
	ParameterizerTransaction *types.Transaction
	DatatrustAddress         common.Address
	DatatrustContract        *datatrust.Datatrust
	DatatrustTransaction     *types.Transaction
	ListingAddress           common.Address
	ListingContract          *listing.Listing
	ListingTransaction       *types.Transaction
	InvestingAddress         common.Address
	InvestingContract        *investing.Investing
	InvestingTransaction     *types.Transaction
}

// Deploy function which, well, deploys our contracts - returning
// a hydrated Dep object ready for use.
// The initialBal argument is an amount of token(in wei) credited to the AuthFactory
// in the deployed MarketToken and EtherToken Contracts.
// Also passed a hydrated Ctx object, used to aid the deploying.
// Returns a hydrated Dep object or any error incurred.
func Deploy(initialBal *big.Int, c *Ctx) (*Dep, error) {
	// Ether Token: consumes: none, privileged: none
	etherTokenAddr, etherTokenTrans, etherTokenCont, etherTokenErr := ethertoken.DeployEtherToken(
		c.AuthFactory,
		c.Blockchain,
		c.AuthFactory.From,
		initialBal,
	)

	if etherTokenErr != nil {
		return nil, etherTokenErr
	}

	// Market Token: consumes: none, privileged: Listing, Investing
	marketTokenAddr, marketTokenTrans, marketTokenCont, marketTokenErr := markettoken.DeployMarketToken(
		c.AuthFactory,
		c.Blockchain,
		c.AuthFactory.From,
		initialBal,
	)

	if marketTokenErr != nil {
		return nil, marketTokenErr
	}

	c.Blockchain.Commit()

	// Voting: consumes: Market Token, privileged: Parameterizer, Datatrust, Listing, Investing
	votingAddr, votingTrans, votingCont, votingErr := voting.DeployVoting(
		c.AuthFactory,
		c.Blockchain,
		marketTokenAddr,
	)

	if votingErr != nil {
		return nil, votingErr
	}

	c.Blockchain.Commit()

	// Parameterizer: consumes: Voting, privileged: none
	paramAddr, paramTrans, paramCont, paramErr := parameterizer.DeployParameterizer(
		c.AuthFactory,
		c.Blockchain,
		votingAddr,
		big.NewInt(ONE_GWEI),   // conversionRate: stipulation is that market token should be, at least, as val as eth
		big.NewInt(100),        // investDenominator, a scaling factor
		big.NewInt(110),        // investNumerator, a scaling factor
		big.NewInt(ONE_WEI),    // listReward (one token)
		big.NewInt(ONE_GWEI),   // access reward
		big.NewInt(ONE_GWEI),   // Stake
		big.NewInt(100),        // voteBy
		big.NewInt(50),         // quorum
		big.NewInt(30),         // backend payment percent
		big.NewInt(50),         // maker payment percent
		big.NewInt(ONE_FINNEY), // cost per byte
	)

	if paramErr != nil {
		return nil, paramErr
	}

	c.Blockchain.Commit()

	dataAddr, dataTrans, dataCont, dataErr := datatrust.DeployDatatrust(
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

	listingAddr, listingTrans, listingCont, listingErr := listing.DeployListing(
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

	investAddr, investTrans, investCont, investErr := investing.DeployInvesting(
		c.AuthFactory,
		c.Blockchain,
		etherTokenAddr,
		marketTokenAddr,
		paramAddr,
	)

	if investErr != nil {
		return nil, investErr
	}

	c.Blockchain.Commit()

	// Set privileged addresses now that contracts are deployed. First: Market Token
	_, mtPrivErr := marketTokenCont.SetPrivileged(GetTxOpts(c.AuthFactory, nil, big.NewInt(ONE_GWEI*2), 100000),
		listingAddr, investAddr)

	if mtPrivErr != nil {
		return nil, mtPrivErr
	}

	// Voting
	_, vtPrivErr := votingCont.SetPrivileged(GetTxOpts(c.AuthFactory, nil, big.NewInt(ONE_GWEI*2), 150000),
		paramAddr, dataAddr, listingAddr, investAddr)

	if vtPrivErr != nil {
		return nil, vtPrivErr
	}

	_, dtPrivErr := dataCont.SetPrivileged(GetTxOpts(c.AuthFactory, nil, big.NewInt(ONE_GWEI*2), 150000), listingAddr)

	if dtPrivErr != nil {
		return nil, dtPrivErr
	}

	c.Blockchain.Commit()

	return &Dep{
		EtherTokenAddress:        etherTokenAddr,
		EtherTokenTransaction:    etherTokenTrans,
		EtherTokenContract:       etherTokenCont,
		MarketTokenAddress:       marketTokenAddr,
		MarketTokenContract:      marketTokenCont,
		MarketTokenTransaction:   marketTokenTrans,
		VotingAddress:            votingAddr,
		VotingContract:           votingCont,
		VotingTransaction:        votingTrans,
		ParameterizerAddress:     paramAddr,
		ParameterizerContract:    paramCont,
		ParameterizerTransaction: paramTrans,
		DatatrustAddress:         dataAddr,
		DatatrustContract:        dataCont,
		DatatrustTransaction:     dataTrans,
		ListingAddress:           listingAddr,
		ListingContract:          listingCont,
		ListingTransaction:       listingTrans,
		InvestingAddress:         investAddr,
		InvestingContract:        investCont,
		InvestingTransaction:     investTrans,
	}, nil
}
