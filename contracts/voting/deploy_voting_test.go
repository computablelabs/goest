package voting

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

func TestDeployVoting(t *testing.T) {
	if deployedError != nil {
		t.Fatalf("Failed to deploy the Voting contract or Market token: %v", deployedError)
	}

	if len(deployed.VotingAddress.Bytes()) == 0 {
		t.Error("Expected a valid voting deployment address to be returned from deploy, got empty byte array instead")
	}
}

// NOTE: actual setting done in Main, as it must be, but tested here similar to deploy
func TestSetPrivileged(t *testing.T) {
	factory, p11r, data, list, invest, _ := deployed.VotingContract.GetPrivileged(nil)

	if factory != context.AuthFactory.From {
		t.Fatalf("Expected factory address of %v but got %v", context.AuthFactory.From, factory)
	}

	if list != context.AuthListing.From {
		t.Fatalf("Expected Listing address of %v but got %v", context.AuthListing.From, list)
	}

	if invest != context.AuthInvesting.From {
		t.Fatalf("Expected Investing address of %v but got %v", context.AuthInvesting.From, invest)
	}

	if p11r != context.AuthParameterizer.From {
		t.Fatalf("Expected p11r address of %v but got %v", context.AuthParameterizer.From, p11r)
	}

	if data != context.AuthDatatrust.From {
		t.Fatalf("Expected datatrust address of %v but got %v", context.AuthDatatrust.From, data)
	}
}

// we'll locate our testmain in these deploy_foo_test files as a pattern.
// NOTE the test main is run once per package, therefore
// the ctx and dep vars will be avail to the other tests in the package
func TestMain(m *testing.M) {
	// see ./helpers#context
	context = SetupBlockchain(big.NewInt(ONE_WEI))
	// see ./helpers#deployed
	deployed, deployedError = Deploy(context)

	// the voting contract must have its privileged addresses set or shit won't work, NOTE this is done, IRL, by the factory
	_, err := deployed.VotingContract.SetPrivileged(&bind.TransactOpts{
		From:     context.AuthFactory.From,
		Signer:   context.AuthFactory.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, context.AuthParameterizer.From, context.AuthDatatrust.From, context.AuthListing.From, context.AuthInvesting.From)

	if err != nil {
		// no T pointer here...
		log.Fatalf("Error setting privileged contract addresses: %v", err)
	}

	context.Blockchain.Commit()
	code := m.Run()
	os.Exit(code)
}
