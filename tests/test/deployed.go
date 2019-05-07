package test

import (
	"github.com/computablelabs/goest/contracts/datatrust"
	"github.com/computablelabs/goest/contracts/ethertoken"
	"github.com/computablelabs/goest/contracts/investing"
	"github.com/computablelabs/goest/contracts/listing"
	"github.com/computablelabs/goest/contracts/markettoken"
	"github.com/computablelabs/goest/contracts/parameterizer"
	"github.com/computablelabs/goest/contracts/voting"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// the params struct holds values that will be set in the P11r at init during deploy
type Params struct {
	ConversionRate *big.Int
	Spread         *big.Int
	ListReward     *big.Int
	Stake          *big.Int
	VoteBy         *big.Int
	Quorum         *big.Int
	BackendPct     *big.Int
	MakerPct       *big.Int
	CostPerByte    *big.Int
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
	InvestingAddress         common.Address
	InvestingContract        *investing.Investing
	InvestingTransaction     *types.Transaction
}

// Deploy function which, well, deploys our contracts - returning
// a hydrated Dep object ready for use.
// The etBal argument is an amount of ether token(in wei) credited to the authOwner
// The mtBal argument is an amount of market token(in wei) credited to the authOwner
// Also passed a hydrated Ctx object, used to aid the deploying.
// Params arg is hydrated with values for the parameterizer.
// Returns a hydrated Dep object or any error incurred.
//func Deploy(etBal *big.Int, mtBal *big.Int, c *Ctx, p *Params) (*Dep, error) {
func Deploy(etBal *big.Int, mtBal *big.Int, authOwner *bind.TransactOpts, blockchain *backends.SimulatedBackend, p *Params) (*Dep, error) {
	// Ether Token: { consumes: [none], privileged: [none] }
	etherTokenAddr, etherTokenTrans, etherTokenCont, etherTokenErr := ethertoken.DeployEtherToken(
		authOwner,
		blockchain,
		authOwner.From,
		etBal,
	)

	if etherTokenErr != nil {
		return nil, etherTokenErr
	}

	// Market Token: { consumes: [none], privileged: [listing, investing] }
	marketTokenAddr, marketTokenTrans, marketTokenCont, marketTokenErr := markettoken.DeployMarketToken(
		authOwner,
		blockchain,
		authOwner.From,
		mtBal,
	)

	if marketTokenErr != nil {
		return nil, marketTokenErr
	}
	blockchain.Commit()

	// Voting: { consumes: [market token], privileged:[parameterizer, datatrust, listing, investing] }
	votingAddr, votingTrans, votingCont, votingErr := voting.DeployVoting(
		authOwner,
		blockchain,
		marketTokenAddr,
	)

	if votingErr != nil {
		return nil, votingErr
	}
	blockchain.Commit()

	// Parameterizer: { consumes: [voting], privileged: [none] }
	paramAddr, paramTrans, paramCont, paramErr := parameterizer.DeployParameterizer(
		authOwner,
		blockchain,
		votingAddr,
		p.ConversionRate,
		p.Spread,
		p.ListReward,
		p.Stake,
		p.VoteBy,
		p.Quorum,
		p.BackendPct,
		p.MakerPct,
		p.CostPerByte,
	)

	if paramErr != nil {
		return nil, paramErr
	}
	blockchain.Commit()

	// Investing: { consumes: [ether token, market token, parameterizer], privileged: [none] }
	investAddr, investTrans, investCont, investErr := investing.DeployInvesting(
		authOwner,
		blockchain,
		etherTokenAddr,
		marketTokenAddr,
		paramAddr,
	)

	if investErr != nil {
		return nil, investErr
	}
	blockchain.Commit()

	// Datatrust: { consumes: [ether token, voting, parameterizer], privileged: [listing] }
	dataAddr, dataTrans, dataCont, dataErr := datatrust.DeployDatatrust(
		authOwner,
		blockchain,
		etherTokenAddr,
		votingAddr,
		paramAddr,
		investAddr, // does not consume the investing contract, simply references the address to transfer to reserve
	)

	if dataErr != nil {
		return nil, dataErr
	}
	blockchain.Commit()

	// Listing: { consumes: [market token, voting, parameterizer, datatrust, investing], privileged: [none] }
	listingAddr, listingTrans, listingCont, listingErr := listing.DeployListing(
		authOwner,
		blockchain,
		marketTokenAddr,
		votingAddr,
		paramAddr,
		dataAddr,
		investAddr,
	)

	if listingErr != nil {
		return nil, listingErr
	}
	blockchain.Commit()

	// Set privileged addresses now that contracts are deployed. First: Market Token
	_, mtPrivErr := marketTokenCont.SetPrivileged(GetTxOpts(authOwner, nil, big.NewInt(ONE_GWEI*2), 100000),
		listingAddr, investAddr)

	if mtPrivErr != nil {
		return nil, mtPrivErr
	}

	// Voting
	_, vtPrivErr := votingCont.SetPrivileged(GetTxOpts(authOwner, nil, big.NewInt(ONE_GWEI*2), 150000),
		paramAddr, dataAddr, listingAddr, investAddr)

	if vtPrivErr != nil {
		return nil, vtPrivErr
	}

	_, dtPrivErr := dataCont.SetPrivileged(GetTxOpts(authOwner, nil, big.NewInt(ONE_GWEI*2), 150000), listingAddr)

	if dtPrivErr != nil {
		return nil, dtPrivErr
	}

	blockchain.Commit()

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
