package test

import (
	"github.com/computablelabs/goest/contracts/datatrust"
	"github.com/computablelabs/goest/contracts/ethertoken"
	"github.com/computablelabs/goest/contracts/listing"
	"github.com/computablelabs/goest/contracts/markettoken"
	"github.com/computablelabs/goest/contracts/parameterizer"
	"github.com/computablelabs/goest/contracts/reserve"
	"github.com/computablelabs/goest/contracts/voting"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// the params struct holds values that will be set in the P11r at init during deploy
type Params struct {
	PriceFloor  *big.Int
	Spread      *big.Int
	ListReward  *big.Int
	Stake       *big.Int
	VoteBy      *big.Int
	Plurality   *big.Int
	BackendPct  *big.Int
	MakerPct    *big.Int
	CostPerByte *big.Int
}

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
	ReserveAddress           common.Address
	ReserveContract          *reserve.Reserve
	ReserveTransaction       *types.Transaction
}

// Deploy function which, well, deploys our contracts - returning
// a hydrated Dep object ready for use.
// The mtBal argument is an amount of market token(in wei) credited to the AuthOwner
// Also passed a hydrated Ctx object, used to aid the deploying.
// Params arg is hydrated with values for the parameterizer.
// Returns a hydrated Dep object or any error incurred.
func Deploy(mtBal *big.Int, c *Ctx, p *Params) (*Dep, error) {
	// Ether Token: { consumes: [none], privileged: [none] }
	etherTokenAddr, etherTokenTrans, etherTokenCont, etherTokenErr := ethertoken.DeployEtherToken(
		c.AuthOwner,
		c.Blockchain,
	)

	if etherTokenErr != nil {
		return nil, etherTokenErr
	}

	// Market Token: { consumes: [none], privileged: [listing, reserve] }
	marketTokenAddr, marketTokenTrans, marketTokenCont, marketTokenErr := markettoken.DeployMarketToken(
		c.AuthOwner,
		c.Blockchain,
		c.AuthOwner.From,
		mtBal,
	)

	if marketTokenErr != nil {
		return nil, marketTokenErr
	}
	c.Blockchain.Commit()

	// Voting: { consumes: [market token], privileged:[parameterizer, datatrust, listing] }
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
		marketTokenAddr,
		votingAddr,
		p.PriceFloor,
		p.Spread,
		p.ListReward,
		p.Stake,
		p.VoteBy,
		p.Plurality,
		p.BackendPct,
		p.MakerPct,
		p.CostPerByte,
	)

	if paramErr != nil {
		return nil, paramErr
	}
	c.Blockchain.Commit()

	// Reserve: { consumes: [ether token, market token, parameterizer], privileged: [none] }
	resAddr, resTrans, resCont, resErr := reserve.DeployReserve(
		c.AuthOwner,
		c.Blockchain,
		etherTokenAddr,
		marketTokenAddr,
		paramAddr,
	)

	if resErr != nil {
		return nil, resErr
	}
	c.Blockchain.Commit()

	// Datatrust: { consumes: [ether token, voting, parameterizer], privileged: [listing] }
	dataAddr, dataTrans, dataCont, dataErr := datatrust.DeployDatatrust(
		c.AuthOwner,
		c.Blockchain,
		etherTokenAddr,
		votingAddr,
		paramAddr,
		resAddr, // does not consume the reserve contract, simply references the address to transfer to reserve
	)

	if dataErr != nil {
		return nil, dataErr
	}
	c.Blockchain.Commit()

	// Listing: { consumes: [market token, voting, parameterizer, datatrust, reserve], privileged: [none] }
	listingAddr, listingTrans, listingCont, listingErr := listing.DeployListing(
		c.AuthOwner,
		c.Blockchain,
		marketTokenAddr,
		votingAddr,
		paramAddr,
		resAddr,
		dataAddr,
	)

	if listingErr != nil {
		return nil, listingErr
	}
	c.Blockchain.Commit()

	// Set privileged addresses now that contracts are deployed. First: Market Token
	_, mtPrivErr := marketTokenCont.SetPrivileged(GetTxOpts(c.AuthOwner, nil, big.NewInt(ONE_GWEI*2), 100000),
		resAddr, listingAddr)

	if mtPrivErr != nil {
		return nil, mtPrivErr
	}

	// Voting
	_, vtPrivErr := votingCont.SetPrivileged(GetTxOpts(c.AuthOwner, nil, big.NewInt(ONE_GWEI*2), 150000),
		paramAddr, dataAddr, listingAddr)

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
		ReserveAddress:           resAddr,
		ReserveContract:          resCont,
		ReserveTransaction:       resTrans,
		DatatrustAddress:         dataAddr,
		DatatrustContract:        dataCont,
		DatatrustTransaction:     dataTrans,
		ListingAddress:           listingAddr,
		ListingContract:          listingCont,
		ListingTransaction:       listingTrans,
	}, nil
}
