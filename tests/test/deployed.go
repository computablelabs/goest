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
// The initialBal argument is an amount of token(in wei) credited to the AuthOwner
// in the deployed MarketToken and EtherToken Contracts.
// Also passed a hydrated Ctx object, used to aid the deploying.
// Returns a hydrated Dep object or any error incurred.
func Deploy(initialBal *big.Int, c *Ctx) (*Dep, error) {
	// Ether Token: { consumes: [none], privileged: [none] }
	etherTokenAddr, etherTokenTrans, etherTokenCont, etherTokenErr := ethertoken.DeployEtherToken(
		c.AuthOwner,
		c.Blockchain,
		c.AuthOwner.From,
		initialBal,
	)

	if etherTokenErr != nil {
		return nil, etherTokenErr
	}

	// Market Token: { consumes: [none], privileged: [listing, investing] }
	marketTokenAddr, marketTokenTrans, marketTokenCont, marketTokenErr := markettoken.DeployMarketToken(
		c.AuthOwner,
		c.Blockchain,
		c.AuthOwner.From,
		initialBal,
	)

	if marketTokenErr != nil {
		return nil, marketTokenErr
	}
	c.Blockchain.Commit()

	// Voting: { consumes: [market token], privileged:[parameterizer, datatrust, listing, investing] }
	votingAddr, votingTrans, votingCont, votingErr := voting.DeployVoting(
		c.AuthOwner,
		c.Blockchain,
		marketTokenAddr,
	)

	if votingErr != nil {
		return nil, votingErr
	}
	c.Blockchain.Commit()

	// Parameterizer: { consumes: [voting], privileged: [none] }
	paramAddr, paramTrans, paramCont, paramErr := parameterizer.DeployParameterizer(
		c.AuthOwner,
		c.Blockchain,
		votingAddr,
		big.NewInt(ONE_GWEI),     // conversionRate: stipulation is that market token should be, at least, as val as eth
		big.NewInt(110),          // spread, a scaling factor
		big.NewInt(ONE_WEI),      // listReward (one token)
		big.NewInt(ONE_GWEI),     // Stake
		big.NewInt(100),          // voteBy
		big.NewInt(50),           // quorum
		big.NewInt(25),           // backend payment percent
		big.NewInt(50),           // maker payment percent
		big.NewInt(ONE_FINNEY*6), // cost per byte at 5000 wei
	)

	if paramErr != nil {
		return nil, paramErr
	}
	c.Blockchain.Commit()

	// Investing: { consumes: [ether token, market token, parameterizer], privileged: [none] }
	investAddr, investTrans, investCont, investErr := investing.DeployInvesting(
		c.AuthOwner,
		c.Blockchain,
		etherTokenAddr,
		marketTokenAddr,
		paramAddr,
	)

	if investErr != nil {
		return nil, investErr
	}
	c.Blockchain.Commit()

	// Datatrust: { consumes: [ether token, voting, parameterizer], privileged: [listing] }
	dataAddr, dataTrans, dataCont, dataErr := datatrust.DeployDatatrust(
		c.AuthOwner,
		c.Blockchain,
		etherTokenAddr,
		votingAddr,
		paramAddr,
		investAddr, // does not consume the investing contract, simply references the address to transfer to reserve
	)

	if dataErr != nil {
		return nil, dataErr
	}
	c.Blockchain.Commit()

	// Listing: { consumes: [market token, voting, parameterizer, datatrust, investing], privileged: [none] }
	listingAddr, listingTrans, listingCont, listingErr := listing.DeployListing(
		c.AuthOwner,
		c.Blockchain,
		marketTokenAddr,
		votingAddr,
		paramAddr,
		dataAddr,
		investAddr,
	)

	if listingErr != nil {
		return nil, listingErr
	}
	c.Blockchain.Commit()

	// Set privileged addresses now that contracts are deployed. First: Market Token
	_, mtPrivErr := marketTokenCont.SetPrivileged(GetTxOpts(c.AuthOwner, nil, big.NewInt(ONE_GWEI*2), 100000),
		listingAddr, investAddr)

	if mtPrivErr != nil {
		return nil, mtPrivErr
	}

	// Voting
	_, vtPrivErr := votingCont.SetPrivileged(GetTxOpts(c.AuthOwner, nil, big.NewInt(ONE_GWEI*2), 150000),
		paramAddr, dataAddr, listingAddr, investAddr)

	if vtPrivErr != nil {
		return nil, vtPrivErr
	}

	_, dtPrivErr := dataCont.SetPrivileged(GetTxOpts(c.AuthOwner, nil, big.NewInt(ONE_GWEI*2), 150000), listingAddr)

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
