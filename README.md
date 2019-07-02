# Goest

## Golang testing and sandbox for the Computable Protocol.
[![Build Status](https://travis-ci.org/computablelabs/goest.svg?branch=master)](https://travis-ci.org/computablelabs/goest)

### Installation

- First start by installing vyper. See [vyper installation instructions](https://github.com/ethereum/vyper/blob/master/docs/installing-vyper.rst).
- Make sure your Go environment is configured. See [go installation directions](https://golang.org/doc/install)
- The test suite uses `geth` utilities and libraries. You can install `geth` with `go get`:
```
go get -d github.com/ethereum/go-ethereum
```
Alternatively, if you'd prefer, you can install Geth from source. See [instructions](https://github.com/ethereum/go-ethereum/wiki/Installing-Geth).

### Run Tests

The test suites are written in Go, so you can run tests from the top level directory with
```
npm run test
```

You can also run individual tests by specifying a subdirectory. For example, here's how you would run the tests for `ethertoken`:

```
go test -v ./tests/ethertokentest/
```

### Making changes to contracts

If you've changed a contract in this repo, you will need to regenerate the `.abi` and `.bin` files tied to that file. These files are used to regenerate the go API bindings for the contract. You finally need to generate the external vyper files. For example, let's suppose that you've changed the `Voting.vy` contract. Then you would run these commands

```
npm run precompile:voting:vyper
npm run compile:voting:vyper
npm run compile:voting:vyper:external
```
You can see all the commands available by checking the [`package.json`](package.json) file.
