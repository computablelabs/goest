package datatrusttest

import (
	"fmt"
	"github.com/computablelabs/goest/tests/test"
	"math/big"
	"testing"
)

func TestExternalRegister(t *testing.T) {
	// Test the registration of an external datatrust which gets funds from support
	// member 3 will need to deposit funds into the ether token
	_, depErr := deployed.EtherTokenContract.Deposit(test.GetTxOpts(context.AuthUser3,
		big.NewInt(test.ONE_ETH*2), big.NewInt(test.ONE_GWEI*2), 1000000)) // deposit 2 tokens
	test.IfNotNil(t, depErr, fmt.Sprintf("Error depositing to ethertoken: %v", depErr))
	context.Blockchain.Commit()

	// allow the ether token to spend on user3's behalf
	_, approveErr := deployed.EtherTokenContract.Approve(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), deployed.ReserveAddress, big.NewInt(test.ONE_ETH*2)) // up to 2 tokens
	test.IfNotNil(t, approveErr, fmt.Sprintf("Error approving market contract to spend: %v", approveErr))
	context.Blockchain.Commit()

	// Buy 10 gwei worth of CMT
	price, _ := deployed.ReserveContract.GetSupportPrice(nil)

	_, sErr := deployed.ReserveContract.Support(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 1000000), price.Mul(price, big.NewInt(10))) // 1e9 * 10
	test.IfNotNil(t, sErr, fmt.Sprintf("Error supporting market: %v", sErr))
	context.Blockchain.Commit()

	// user3 will need to have approved the voting contract to spend at least the stake
	incErr := test.MaybeIncreaseMarketTokenAllowance(context, deployed, context.AuthUser3, deployed.VotingAddress,
		big.NewInt(2*test.ONE_GWEI))
	test.IfNotNil(t, incErr, "Error maybe transferring market token approval")

	// Register the datatrust
	_, regErr := deployed.DatatrustContract.Register(test.GetTxOpts(context.AuthUser3, nil,
		big.NewInt(test.ONE_GWEI*2), 500000), "http://www.icanhazexternalbackend.com")
	test.IfNotNil(t, regErr, fmt.Sprintf("Error registering for backend status: %v", regErr))
	context.Blockchain.Commit()

	// we should have the candidate
	hash, _ := deployed.DatatrustContract.GetHash(nil, "http://www.icanhazexternalbackend.com")
	isReg, _ := deployed.VotingContract.CandidateIs(nil, hash, test.REGISTRATION)
	if isReg != true {
		t.Errorf("Expected is registered to be true, got: %v", isReg)
	}
}
