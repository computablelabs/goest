package parameterizer

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/core"
	"log"
	"math/big"
	"os"
	"testing"
)

// variables decalred here have package scope
var context *ctx
var deployed *dep
var deployedError error

func TestDeployParameterizer(t *testing.T) {
	if deployedError != nil {
		t.Fatalf("Failed to deploy the Parameterizer contract or a dependency: %v", deployedError)
	}

	if len(deployed.VotingAddress.Bytes()) == 0 {
		t.Error("Expected a valid voting deployment address to be returned from deploy, got empty byte array instead")
	}

	if len(deployed.ParameterizerAddress.Bytes()) == 0 {
		t.Error("Expected a valid parameterizer deployment address to be returned from deploy, got empty byte array instead")
	}
}

func TestVotingSetPrivilegedContracts(t *testing.T) {
	p11r, data, list, invest, _ := deployed.VotingContract.GetPrivileged(nil)

	if p11r != deployed.ParameterizerAddress {
		t.Fatalf("Expected p11r address of %v but got %v", deployed.ParameterizerAddress, p11r)
	}

	if data != context.AuthDatatrust.From {
		t.Fatalf("Expected Datatrust address of %v but got %v", context.AuthDatatrust.From, list)
	}

	if list != context.AuthListing.From {
		t.Fatalf("Expected Listing address of %v but got %v", context.AuthListing.From, list)
	}

	if invest != context.AuthInvesting.From {
		t.Fatalf("Expected Investing address of %v but got %v", context.AuthInvesting.From, invest)
	}
}

// we'll locate our testmain in these deploy_foo_test files as a pattern.
// NOTE the test main is run once per package, therefore
// the ctx and dep vars will be avail to the other tests in the package
func TestMain(m *testing.M) {
	// see ./helpers#context
	context = SetupBlockchain(big.NewInt(ONE_WEI)) // 1 ETH in wei
	// see ./helpers#deployed
	deployed, deployedError = Deploy(big.NewInt(ONE_WEI*5), context)

	// after both of these steps, must we give the p11r some funds?
	// context.Alloc[deployed.ParameterizerAddress] = core.GenesisAccount{Balance: big.NewInt(1000000000000000000)}

	// the markettoken must have its privileges set
	_, marketErr := deployed.MarketTokenContract.SetPrivileged(&bind.TransactOpts{
		From:     context.AuthFactory.From,
		Signer:   context.AuthFactory.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, context.AuthListing.From, context.AuthInvesting.From)

	if marketErr != nil {
		log.Fatalf("Error setting privileged contract address: %v", marketErr)
	}

	// the voting contract must have its privileged addresk, NOTE this is done, IRL, by the factory
	_, err := deployed.VotingContract.SetPrivileged(&bind.TransactOpts{
		From:     context.AuthFactory.From,
		Signer:   context.AuthFactory.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, deployed.ParameterizerAddress, context.AuthDatatrust.From, context.AuthListing.From, context.AuthInvesting.From)

	if err != nil {
		// no T pointer here...
		log.Fatalf("Error setting privileged contract addresses: %v", err)
	}

	context.Blockchain.Commit()
	code := m.Run()
	os.Exit(code)
}
