package test

import (
	"math/big"
)

const ONE_ETH = 1000000000000000000
const ONE_FINNEY = 1000000000000000
const ONE_SZABO = 1000000000000
const ONE_GWEI = 1000000000
const ONE_MWEI = 1000000
const ONE_KWEI = 1000

// const ONE_SZABO = 1000000
// const ONE_FINNEY = 1000

var UNDEFINED = big.NewInt(0)
var APPLICATION = big.NewInt(1)
var CHALLENGE = big.NewInt(2)
var REPARAM = big.NewInt(3)
var REGISTRATION = big.NewInt(4)

var STAKE = big.NewInt(1)
var PRICE_FLOOR = big.NewInt(2)
var LIST_REWARD = big.NewInt(5)
var PLURALITY = big.NewInt(6)
var VOTE_BY = big.NewInt(7)
