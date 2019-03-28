package parameterizer

import (
	"math/big"
	"testing"
)

func TestGetStake(t *testing.T) {
	stake, _ := deployed.ParameterizerContract.GetStake(nil)

	if stake.Cmp(big.NewInt(ONE_GWEI)) != 0 {
		t.Fatalf("Expected challengeStake to be 10**18, got: %v", stake)
	}
}

func TestGetConversionRate(t *testing.T) {
	rate, _ := deployed.ParameterizerContract.GetConversionRate(nil)

	if rate.Cmp(big.NewInt(ONE_GWEI)) != 0 {
		t.Fatalf("Expected conversionRate to be 10**9, got: %v", rate)
	}
}

func TestGetInvestDenominator(t *testing.T) {
	d, _ := deployed.ParameterizerContract.GetInvestDenominator(nil)

	if d.Cmp(big.NewInt(100)) != 0 {
		t.Fatalf("Expected Invest Denominator to be %v, got: %v", 100, d)
	}
}

func TestGetInvestNumerator(t *testing.T) {
	n, _ := deployed.ParameterizerContract.GetInvestNumerator(nil)

	if n.Cmp(big.NewInt(110)) != 0 {
		t.Fatalf("Expected Invest Numerator to be %v, got: %v", 110, n)
	}
}

func TestGetListReward(t *testing.T) {
	reward, _ := deployed.ParameterizerContract.GetListReward(nil)

	if reward.Cmp(big.NewInt(ONE_WEI)) != 0 {
		t.Fatalf("Expected listReward to be 10**18, got: %v", reward)
	}
}

func TestGetAccessReward(t *testing.T) {
	reward, _ := deployed.ParameterizerContract.GetAccessReward(nil)

	if reward.Cmp(big.NewInt(ONE_WEI/2)) != 0 {
		t.Fatalf("Expected listReward to be 10**9, got: %v", reward)
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
		t.Fatalf("Expected voteBy to be 20, got: %v", voteBy)
	}
}

func TestGetBackendPayment(t *testing.T) {
	pct, _ := deployed.ParameterizerContract.GetBackendPayment(nil)

	if pct.Cmp(big.NewInt(30)) != 0 {
		t.Fatalf("Expected backend pay pct to be 30, got: %v", pct)
	}
}

func TestGetMakerPayment(t *testing.T) {
	pct, _ := deployed.ParameterizerContract.GetMakerPayment(nil)

	if pct.Cmp(big.NewInt(50)) != 0 {
		t.Fatalf("Expected maker pay pct to be 50, got: %v", pct)
	}
}
