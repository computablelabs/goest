# Goest

## Golang testing and sandbox for the Computable Protocol.
[![Build Status](https://travis-ci.org/computablelabs/goest.svg?branch=master)](https://travis-ci.org/computablelabs/goest)

### Installation

- First start by installing vyper (see [installation instructions](https://github.com/ethereum/vyper/blob/master/docs/installing-vyper.rst)).
- Make sure your Go environment is configured (see [directions](https://golang.org/doc/install))
- The test suite uses `Geth` utilities and libraries. [Install Geth from source](https://github.com/ethereum/go-ethereum/wiki/Installation-Instructions-for-Ubuntu) in the right place in your Go environment.

### Run Tests

The test suites are written in Go, so you can run tests from the top level directory with
```
go test -v ./tests/...
```

You can also run individual tests by specifying a subdirectory.

### Making changes to contracts

If you've changed a contract in this repo, you will need to regenerate the `.abi` and `.bin` files tied to that file. These files are used to regenerate the go API bindings for the contract. You finally need to generate the external vyper files. For example, let's suppose that you've changed the `Voting.by` contract. Then you would run these commands

```
npm run precompile:voting:vyper
npm run compile:voting:vyper
npm run compile:voting:vyper:external
```


