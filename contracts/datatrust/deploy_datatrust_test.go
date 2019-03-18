package datatrust

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

func TestDeployDatatrust(t *testing.T) {
	if deployedError != nil {
		t.Fatalf("Failed to deploy the Datatrust contract or a dependency: %v", deployedError)
	}

	if len(deployed.VotingAddress.Bytes()) == 0 {
		t.Error("Expected a valid voting deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.ParameterizerAddress.Bytes()) == 0 {
		t.Error("Expected a valid parameterizer deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.DatatrustAddress.Bytes()) == 0 {
		t.Error("Expected a valid Datatrust deployment address to be returned from deploy, got empty byte array instead")
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

func TestVotingSetPrivileged(t *testing.T) {
	p11r, data, list, invest, _ := deployed.VotingContract.GetPrivileged(nil)

	if p11r != deployed.ParameterizerAddress {
		t.Fatalf("Expected p11r address of %v but got %v", deployed.ParameterizerAddress, p11r)
	}

	if data != deployed.DatatrustAddress {
		t.Fatalf("Expected Datatrust address of %v but got %v", deployed.DatatrustAddress, data)
	}

	if list != context.AuthListing.From {
		t.Fatalf("Expected listing address of %v but got %v", context.AuthListing.From, list)
	}

	if invest != context.InvestingAddress {
		t.Fatalf("Expected investing address of %v but got %v", context.InvestingAddress, invest)
	}
}

func TestDatatrustSetPrivileged(t *testing.T) {
	list, _ := deployed.DatatrustContract.GetPrivileged(nil)

	if list != context.AuthListing.From {
		t.Fatalf("Expected Listing address of %v but got %v", context.AuthListing.From, list)
	}
}

func TestMain(m *testing.M) {
	// see ./helpers#context
	context = SetupBlockchain(big.NewInt(ONE_WEI * 3)) // users have 3 ETH
	// see ./helpers#deployed
	deployed, deployedError = Deploy(big.NewInt(ONE_WEI*6), context) // 6 tokens in wei

	_, votingErr := deployed.VotingContract.SetPrivileged(&bind.TransactOpts{
		From:     context.AuthFactory.From,
		Signer:   context.AuthFactory.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, deployed.ParameterizerAddress, deployed.DatatrustAddress, context.AuthListing.From, context.InvestingAddress)

	if votingErr != nil {
		// no T pointer here...
		log.Fatalf("Error setting privileged contract addresses: %v", votingErr)
	}

	_, dataErr := deployed.DatatrustContract.SetPrivileged(&bind.TransactOpts{
		From:     context.AuthFactory.From,
		Signer:   context.AuthFactory.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, context.AuthListing.From)

	if dataErr != nil {
		log.Fatalf("Error setting privileged contract addresses: %v", dataErr)
	}

	context.Blockchain.Commit()
	code := m.Run()
	os.Exit(code)
}
