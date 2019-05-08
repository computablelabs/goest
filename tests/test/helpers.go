package test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

// by defining a simple interface which has a `fatal` method we can use our
// `IfNotNil` function with both testing.T (typical use case)  and log (TestMain blocks...)
type logish interface {
	Fatal(...interface{})
}

// used to generate auth objects
func GetAuthObject() *bind.TransactOpts {
	key, _ := crypto.GenerateKey()
	return bind.NewKeyedTransactor(key)
}

// IfNotNil is a convenience method to throw fatal errors
func IfNotNil(l logish, err error, msg string) {
	if err != nil {
		l.Fatal(msg)
	}
}

// GenBytes32 is a convenience function for some tests where a "GetHash" method is not available.
// Returns a type compatible with bytes32, given a string 32 chars or less.
func GenBytes32(str string) [32]byte {
	bytes := [32]byte{}
	copy(bytes[:], []byte(str))
	return bytes
}

// GetTxOpts is a function which allows us to more succintly place the transaction options into a "send"
// type transaction. We do this in place of wrapping our TX in a global session variable, as, we tend to
// use different values in various places. Returns the assembled Geth TransactOpts
func GetTxOpts(auth *bind.TransactOpts, val *big.Int, price *big.Int, limit uint64) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		Value:    val,
		GasPrice: price,
		GasLimit: limit,
	}
}

// MaybeApprove is a convenience method that inspects the Market Token Allowance of one address to another.
// If it is below the given threshold, increase the allowance until it meets said threshold
// Returns any error incurred.
func MaybeIncreaseMarketTokenApproval(c *Ctx, d *Dep, owner *bind.TransactOpts, spender common.Address, thresh *big.Int) error {
	allowed, _ := d.MarketTokenContract.Allowance(nil, owner.From, spender)
	if !(allowed.Cmp(thresh) >= 0) {
		// to keep the allowed from creeping, we'll set the approved to the threshold if less...
		delta := thresh.Sub(thresh, allowed)
		_, approveErr := d.MarketTokenContract.IncreaseApproval(GetTxOpts(owner, nil,
			big.NewInt(ONE_GWEI*2), 1000000), spender, delta)

		if approveErr != nil {
			return approveErr
		}

		c.Blockchain.Commit()
	}
	return nil
}

// MaybeTransferMarketToken is a convenience method that inspects the Market Token balance of a given address.
// If it is below the given threshold, increase the balance via transfer until it meets said threshold
// Returns any error incurred.
func MaybeTransferMarketToken(c *Ctx, d *Dep, from *bind.TransactOpts, to common.Address, thresh *big.Int) error {
	bal, _ := d.MarketTokenContract.BalanceOf(nil, to)
	if bal.Cmp(thresh) == -1 {
		delta := thresh.Sub(thresh, bal)
		// increase the bal by the delta so that it is exactly the threshold
		_, transErr := d.MarketTokenContract.Transfer(GetTxOpts(from, nil,
			big.NewInt(ONE_GWEI*2), 1000000), to, delta)

		if transErr != nil {
			return transErr
		}

		c.Blockchain.Commit()
	}
	return nil
}

// Commafy will take a big integer and return a string with commas so that logging
// big integers is human readable
func Commafy(n *big.Int) string {
	in := fmt.Sprintf("%d", n)
	out := make([]byte, len(in)+(len(in)-2+int(in[0]/'0'))/3)

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}
