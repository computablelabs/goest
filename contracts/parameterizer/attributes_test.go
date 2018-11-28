package parameterizer

import (
	"math/big"
	"testing"
)

func TestGetChallengeStake(t *testing.T) {
	stake, _ := deployed.ParameterizerContract.GetChallengeStake(nil)

	if stake.Cmp(big.NewInt(1000000000000000000)) != 0 {
		t.Fatalf("Expected challengeStake to be 10**18, got: %v", stake)
	}
}

func TestGetConversionRate(t *testing.T) {
	rate, _ := deployed.ParameterizerContract.GetConversionRate(nil)

	if rate.Cmp(big.NewInt(10000000000000000)) != 0 {
		t.Fatalf("Expected conversionRate to be 10**16, got: %v", rate)
	}
}

func TestGetConversionSlopeDenominator(t *testing.T) {
	d, _ := deployed.ParameterizerContract.GetConversionSlopeDenominator(nil)

	if d.Cmp(big.NewInt(101)) != 0 {
		t.Fatalf("Expected conversionSlopeDenominator to be 101, got: %v", d)
	}
}

func TestGetConversionSlopeNumerator(t *testing.T) {
	n, _ := deployed.ParameterizerContract.GetConversionSlopeNumerator(nil)

	if n.Cmp(big.NewInt(100)) != 0 {
		t.Fatalf("Expected conversionSlopeDenominator to be 100, got: %v", n)
	}
}

func TestGetListReward(t *testing.T) {
	reward, _ := deployed.ParameterizerContract.GetListReward(nil)

	if reward.Cmp(big.NewInt(1000000000000000000)) != 0 {
		t.Fatalf("Expected listReward to be 10**18, got: %v", reward)
	}
}

func TestGetQuorum(t *testing.T) {
	quorum, _ := deployed.ParameterizerContract.GetQuorum(nil)

	if quorum.Cmp(big.NewInt(50)) != 0 {
		t.Fatalf("Expected quorum to be 50, got: %v", quorum)
	}
}

func TestGetVoteBy(t *testing.T) {
	voteBy, _ := deployed.ParameterizerContract.GetVoteBy(nil)

	if voteBy.Cmp(big.NewInt(20)) != 0 {
		t.Fatalf("Expected voteBy to be 300, got: %v", voteBy)
	}
}
