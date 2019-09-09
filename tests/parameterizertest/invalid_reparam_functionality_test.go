package parameterizertest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
)

func TestInvalidSpread(t *testing.T) {
	// Spread must be >= 100
	_, err := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 200000), test.SPREAD, big.NewInt(50))
	test.IfNotNil(t, err, fmt.Sprintf("Error creating proposal: %v", err))

	context.Blockchain.Commit()

	// we should not see a candidate now
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(4), big.NewInt(50))
	isCan, _ := deployed.VotingContract.IsCandidate(nil, paramHash)

	if isCan {
		t.Errorf("Expected isCandidate to be false, got: %v", isCan)
	}
}

func TestInvalidStake(t *testing.T) {
	// Stake must be <= 1/3 of MarketToken supply
	_, err := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
		// Attempt to set the stake to 2 MarketToken, greater than 1/3 of
		// existing market cap of 5 MarketToken
		big.NewInt(test.ONE_GWEI*2), 200000), test.STAKE, big.NewInt(test.ONE_ETH*2))
	test.IfNotNil(t, err, fmt.Sprintf("Error creating proposal: %v", err))

	context.Blockchain.Commit()

	// we should not see a candidate now
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(1), big.NewInt(test.ONE_ETH*2))
	isCan, _ := deployed.VotingContract.IsCandidate(nil, paramHash)

	if isCan {
		t.Errorf("Expected isCandidate to be false, got: %v", isCan)
	}
}

func TestInvalidVoteByTooLong(t *testing.T) {
	// Vote_by must be <= 1209600. We'll try setting it to 1500000
	_, err := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 200000), test.VOTE_BY, big.NewInt(1500000))
	test.IfNotNil(t, err, fmt.Sprintf("Error creating proposal: %v", err))

	context.Blockchain.Commit()

	// we should not see a candidate now
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(1500000))
	isCan, _ := deployed.VotingContract.IsCandidate(nil, paramHash)

	if isCan {
		t.Errorf("Expected isCandidate to be false, got: %v", isCan)
	}
}

//func TestInvalidVoteByTooShort(t *testing.T) {
//	// Vote_by must be >= test.MIN_VOTE_BY. We'll try setting it to 5000
//	_, err := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
//		big.NewInt(test.ONE_GWEI*2), 200000), test.VOTE_BY, big.NewInt(5000))
//	test.IfNotNil(t, err, fmt.Sprintf("Error creating proposal: %v", err))
//
//	context.Blockchain.Commit()
//
//	// we should not see a candidate now
//	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(7), big.NewInt(5000))
//	isCan, _ := deployed.VotingContract.IsCandidate(nil, paramHash)
//
//	if isCan {
//		t.Errorf("Expected isCandidate to be false, got: %v", isCan)
//	}
//}

func TestInvalidPluraity(t *testing.T) {
	// Pluraity must be <= 100
	_, err := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 200000), test.PLURALITY, big.NewInt(150))
	test.IfNotNil(t, err, fmt.Sprintf("Error creating proposal: %v", err))

	context.Blockchain.Commit()

	// we should not see a candidate now
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(6), big.NewInt(150))
	isCan, _ := deployed.VotingContract.IsCandidate(nil, paramHash)

	if isCan {
		t.Errorf("Expected isCandidate to be false, got: %v", isCan)
	}
}

func TestInvalidMakerPayment(t *testing.T) {
	// maker_payment + backend_payment must be <= 100. backend_payment is 25.
	// Let's try setting maker_payment to 90
	_, err := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 200000), test.MAKER_PAYMENT, big.NewInt(90))
	test.IfNotNil(t, err, fmt.Sprintf("Error creating proposal: %v", err))

	context.Blockchain.Commit()

	// we should not see a candidate now
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(9), big.NewInt(90))
	isCan, _ := deployed.VotingContract.IsCandidate(nil, paramHash)

	if isCan {
		t.Errorf("Expected isCandidate to be false, got: %v", isCan)
	}
}

func TestInvalidBackendPayment(t *testing.T) {
	// maker_payment + backend_payment must be <= 100. maker is 50.
	// Let's try setting backend to 60
	_, err := deployed.ParameterizerContract.Reparameterize(test.GetTxOpts(context.AuthUser1, nil,
		big.NewInt(test.ONE_GWEI*2), 200000), test.BACKEND_PAYMENT, big.NewInt(60))
	test.IfNotNil(t, err, fmt.Sprintf("Error creating proposal: %v", err))

	context.Blockchain.Commit()

	// we should not see a candidate now
	paramHash, _ := deployed.ParameterizerContract.GetHash(nil, big.NewInt(8), big.NewInt(60))
	isCan, _ := deployed.VotingContract.IsCandidate(nil, paramHash)

	if isCan {
		t.Errorf("Expected isCandidate to be false, got: %v", isCan)
	}
}
