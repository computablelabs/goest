package market

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"math/big"
	"os"
	"testing"
)

// variables decalred here have package scope
var context *ctx
var deployed *dep
var deployedError error

func TestDeployMarket(t *testing.T) {
	if deployedError != nil {
		t.Fatalf("Failed to deploy the Market contract or a dependency: %v", deployedError)
	}

	if len(deployed.NetworkTokenAddress.Bytes()) == 0 {
		t.Error("Expected a valid network token deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.MarketTokenAddress.Bytes()) == 0 {
		t.Error("Expected a valid market token deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.VotingAddress.Bytes()) == 0 {
		t.Error("Expected a valid voting deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.ParameterizerAddress.Bytes()) == 0 {
		t.Error("Expected a valid parameterizer deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.MarketAddress.Bytes()) == 0 {
		t.Error("Expected a valid market deployment address to be returned from deploy, got empty byte array instead")
	}

	// what does the hex string look like?
	// t.Log("********************************************************")
	// t.Log(deployed.MarketAddress.Hex())

	// ntr, _ := context.Blockchain.TransactionReceipt(nil, deployed.NetworkTokenTransaction.Hash())
	// mtr, _ := context.Blockchain.TransactionReceipt(nil, deployed.MarketTokenTransaction.Hash())
	// pr, _ := context.Blockchain.TransactionReceipt(nil, deployed.ParameterizerTransaction.Hash())
	// vr, _ := context.Blockchain.TransactionReceipt(nil, deployed.VotingTransaction.Hash())
	// mr, _ := context.Blockchain.TransactionReceipt(nil, deployed.MarketTransaction.Hash())

	// t.Log("******************** Costs to deploy all-the-things *************************")
	// t.Logf("Gas used to deploy Network Token %v", ntr.GasUsed)
	// t.Logf("Gas used to deploy Market Token %v", mtr.GasUsed)
	// t.Logf("Gas used to deploy Parameterizer %v", pr.GasUsed)
	// t.Logf("Gas used to deploy Voting %v", vr.GasUsed)
	// t.Logf("Gas used to deploy Market %v", mr.GasUsed)
}

func TestMarketTokenSetPrivilegedContracts(t *testing.T) {
	market, _ := deployed.MarketTokenContract.GetPrivilegedAddresses(nil)

	if market != deployed.MarketAddress {
		t.Fatalf("Expected market address of %v but got %v", deployed.MarketAddress, market)
	}
}

func TestVotingSetPrivilegedContracts(t *testing.T) {
	market, p11r, _ := deployed.VotingContract.GetPrivilegedAddresses(nil)

	if market != deployed.MarketAddress {
		t.Fatalf("Expected market address of %v but got %v", deployed.MarketAddress, market)
	}

	if p11r != deployed.ParameterizerAddress {
		t.Fatalf("Expected p11r address of %v but got %v", deployed.ParameterizerAddress, p11r)
	}
}

// we'll locate our testmain in these deploy_foo_test files as a pattern.
// NOTE the test main is run once per package, therefore
// the ctx and dep vars will be avail to the other tests in the package
func TestMain(m *testing.M) {
	// see ./helpers#context
	context = SetupBlockchain(big.NewInt(ONE_WEI))
	// see ./helpers#deployed
	deployed, deployedError = Deploy(big.NewInt(ONE_WEI*6), context) // 6 tokens in wei

	// the markettoken must have its privileges set
	_, marketErr := deployed.MarketTokenContract.SetPrivilegedContracts(&bind.TransactOpts{
		From:     context.AuthFactory.From,
		Signer:   context.AuthFactory.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, deployed.MarketAddress)

	if marketErr != nil {
		log.Fatalf("Error setting privileged contract address: %v", marketErr)
	}

	_, votingErr := deployed.VotingContract.SetPrivilegedContracts(&bind.TransactOpts{
		From:     context.AuthFactory.From,
		Signer:   context.AuthFactory.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, deployed.MarketAddress, deployed.ParameterizerAddress)

	if votingErr != nil {
		// no T pointer here...
		log.Fatalf("Error setting privileged contract addresses: %v", votingErr)
	}

	context.Blockchain.Commit()
	code := m.Run()
	os.Exit(code)
}
