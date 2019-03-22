package listing

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"math/big"
	"os"
	"testing"
	"time"
)

// variables decalred here have package scope
var context *ctx
var deployed *dep
var deployedError error

func TestDeployListing(t *testing.T) {
	if deployedError != nil {
		t.Fatalf("Failed to deploy the Listing contract or a dependency: %v", deployedError)
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

	if len(deployed.ListingAddress.Bytes()) == 0 {
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

func TestMarketTokenSetPrivileged(t *testing.T) {
	list, invest, _ := deployed.MarketTokenContract.GetPrivileged(nil)

	if list != deployed.ListingAddress {
		t.Fatalf("Expected listing address of %v but got %v", deployed.ListingAddress, list)
	}

	if invest != context.AuthInvest.From {
		t.Fatalf("Expected investing address of %v but got %v", context.AuthInvest.From, invest)
	}
}

func TestVotingSetPrivileged(t *testing.T) {
	p11r, data, list, invest, _ := deployed.VotingContract.GetPrivileged(nil)

	if p11r != deployed.ParameterizerAddress {
		t.Fatalf("Expected p11r address of %v but got %v", deployed.ParameterizerAddress, p11r)
	}

	if data != deployed.DatatrustAddress {
		t.Fatalf("Expected datatrust address of %v but got %v", deployed.DatatrustAddress, data)
	}

	if list != deployed.ListingAddress {
		t.Fatalf("Expected listing address of %v but got %v", deployed.ListingAddress, list)
	}

	if invest != context.AuthInvest.From {
		t.Fatalf("Expected investing address of %v but got %v", context.AuthInvest.From, invest)
	}
}

func TestDatatrustSetPrivileged(t *testing.T) {
	list, _ := deployed.DatatrustContract.GetPrivileged(nil)

	if list != deployed.ListingAddress {
		t.Fatalf("Expected listing address of %v but got %v", deployed.ListingAddress, list)
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
	}, deployed.ListingAddress, context.AuthInvest.From)

	if marketErr != nil {
		log.Fatalf("Error setting privileged contract address: %v", marketErr)
	}

	_, votingErr := deployed.VotingContract.SetPrivileged(&bind.TransactOpts{
		From:     context.AuthFactory.From,
		Signer:   context.AuthFactory.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, deployed.ParameterizerAddress, deployed.DatatrustAddress, deployed.ListingAddress, context.AuthInvest.From)

	if votingErr != nil {
		// no T pointer here...
		log.Fatalf("Error setting privileged contract addresses: %v", votingErr)
	}

	// set datatrust priv...
	_, dataErr := deployed.DatatrustContract.SetPrivileged(&bind.TransactOpts{
		From:     context.AuthFactory.From,
		Signer:   context.AuthFactory.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, deployed.ListingAddress)

	if dataErr != nil {
		log.Fatalf("Error setting privileged contract addresses: %v", dataErr)
	}

	// setup the datatrust with a backend
	_, regErr := deployed.DatatrustContract.Register(&bind.TransactOpts{
		From:     context.AuthBackend.From,
		Signer:   context.AuthBackend.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 500000,
	}, "https://www.immabackend.biz")

	if regErr != nil {
		log.Fatalf("Error registering for backend status: %v", regErr)
	}

	context.Blockchain.Commit()

	// make member2 a council member (if not one). the invest contract can do this...
	isMember, _ := deployed.VotingContract.InCouncil(nil, context.AuthMember2.From)
	if isMember != true {
		_, councilErr := deployed.VotingContract.AddToCouncil(&bind.TransactOpts{
			From:     context.AuthInvest.From,
			Signer:   context.AuthInvest.Signer,
			GasPrice: big.NewInt(ONE_GWEI * 2),
			GasLimit: 100000,
		}, context.AuthMember2.From)

		if councilErr != nil {
			log.Fatal("Error adding member to council")
		}
	}

	context.Blockchain.Commit()

	// vote for the backend candidate
	hash, _ := deployed.DatatrustContract.GetHash(nil, "https://www.immabackend.biz")
	_, voteErr := deployed.VotingContract.Vote(&bind.TransactOpts{
		From:     context.AuthMember2.From,
		Signer:   context.AuthMember2.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 100000,
	}, hash, big.NewInt(1))

	if voteErr != nil {
		log.Fatalf("Error voting for candidate: %v", voteErr)
	}

	context.Blockchain.Commit()

	// move past the voteBy
	context.Blockchain.AdjustTime(100 * time.Second)
	context.Blockchain.Commit()

	// any council member can call for resolution
	_, resolveErr := deployed.DatatrustContract.ResolveRegistration(&bind.TransactOpts{
		From:     context.AuthMember2.From,
		Signer:   context.AuthMember2.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, hash)

	if resolveErr != nil {
		log.Fatalf("Error resolving application: %v", resolveErr)
	}

	context.Blockchain.Commit()
	code := m.Run()
	os.Exit(code)
}
