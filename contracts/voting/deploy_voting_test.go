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
func TestSetPrivilegedContracts(t *testing.T) {
	market, p11r, _ := deployed.VotingContract.GetPrivilegedAddresses(nil)

	if market != context.AuthMarket.From {
		t.Fatalf("Expected market address of %v but got %v", context.AuthMarket.From, market)
	}

	if p11r != context.AuthParameterizer.From {
		t.Fatalf("Expected p11r address of %v but got %v", context.AuthParameterizer.From, p11r)
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
	_, err := deployed.VotingContract.SetPrivilegedContracts(&bind.TransactOpts{
		From:     context.AuthFactory.From,
		Signer:   context.AuthFactory.Signer,
		GasPrice: big.NewInt(ONE_GWEI * 2),
		GasLimit: 1000000,
	}, context.AuthMarket.From, context.AuthParameterizer.From)

	if err != nil {
		// no T pointer here...
		log.Fatalf("Error setting privileged contract addresses: %v", err)
	}

	context.Blockchain.Commit()
	code := m.Run()
	os.Exit(code)
}
