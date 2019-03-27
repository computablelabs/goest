package investing

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

func TestDeployInvesting(t *testing.T) {
	if deployedError != nil {
		t.Fatalf("Failed to deploy the Investing contract or a dependency: %v", deployedError)
	}

	if len(deployed.EtherTokenAddress.Bytes()) == 0 {
		t.Error("Expected a valid ether token deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.MarketTokenAddress.Bytes()) == 0 {
		t.Error("Expected a valid market token deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.ParameterizerAddress.Bytes()) == 0 {
		t.Error("Expected a valid parameterizer deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.InvestingAddress.Bytes()) == 0 {
		t.Error("Expected a valid Listing deployment address to be returned from deploy, got empty byte array instead")
	}

	// what does the hex string look like?
	// t.Log("********************************************************")
	// t.Log(deployed.MarketAddress.Hex())

	// ntr, _ := context.Blockchain.TransactionReceipt(nil, deployed.EtherTokenTransaction.Hash())
	// mtr, _ := context.Blockchain.TransactionReceipt(nil, deployed.MarketTokenTransaction.Hash())
	// pr, _ := context.Blockchain.TransactionReceipt(nil, deployed.ParameterizerTransaction.Hash())
	// vr, _ := context.Blockchain.TransactionReceipt(nil, deployed.VotingTransaction.Hash())
	// mr, _ := context.Blockchain.TransactionReceipt(nil, deployed.MarketTransaction.Hash())

	// t.Log("******************** Costs to deploy all-the-things *************************")
	// t.Logf("Gas used to deploy Ether Token %v", ntr.GasUsed)
	// t.Logf("Gas used to deploy Market Token %v", mtr.GasUsed)
	// t.Logf("Gas used to deploy Parameterizer %v", pr.GasUsed)
	// t.Logf("Gas used to deploy Voting %v", vr.GasUsed)
	// t.Logf("Gas used to deploy Market %v", mr.GasUsed)
}

func TestMarketTokenSetPrivilegedContracts(t *testing.T) {
	list, invest, _ := deployed.MarketTokenContract.GetPrivileged(nil)

	if list != context.ListingAddress {
		t.Fatalf("Expected listing address of %v but got %v", context.ListingAddress, list)
	}

	if invest != deployed.InvestingAddress {
		t.Fatalf("Expected investing address of %v but got %v", deployed.InvestingAddress, invest)
	}
}

// we'll locate our testmain in these deploy_foo_test files as a pattern.
// NOTE the test main is run once per package, therefore
// the ctx and dep vars will be avail to the other tests in the package
func TestMain(m *testing.M) {
	// see ./helpers#context
	context = SetupBlockchain(big.NewInt(ONE_WEI * 3)) // users have 3 ETH
	// see ./helpers#deployed
	deployed, deployedError = Deploy(big.NewInt(ONE_WEI*6), context) // 6 tokens in wei

	// the markettoken must have its privileges set
	_, marketErr := deployed.MarketTokenContract.SetPrivileged(&bind.TransactOpts{
		From:     context.AuthFactory.From,
		Signer:   context.AuthFactory.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, context.ListingAddress, deployed.InvestingAddress)

	if marketErr != nil {
		log.Fatalf("Error setting privileged contract address: %v", marketErr)
	}

	context.Blockchain.Commit()
	code := m.Run()
	os.Exit(code)
}
