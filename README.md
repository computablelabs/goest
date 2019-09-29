# Goest

## Golang testing and sandbox for the Computable Protocol.
[![Build Status](https://travis-ci.org/computablelabs/goest.svg?branch=master)](https://travis-ci.org/computablelabs/goest)

### Installation
- First start by installing vyper. See [vyper installation instructions](https://github.com/ethereum/vyper/blob/master/docs/installing-vyper.rst).
- Make sure your Go environment is configured. See [go installation directions](https://golang.org/doc/install)
- The test suite uses `geth` utilities and libraries. You can install `geth` with `go get`:

    go get -d github.com/ethereum/go-ethereum

Alternatively, if you'd prefer, you can install Geth from source. See [instructions](https://github.com/ethereum/go-ethereum/wiki/Installing-Geth).

#### Developer Note: Build ABIGEN
If you are editing and compiling the smart contracts you will need to ensure that `abigen` is available to your go environment.

    cd $GOPATH/src/github.com/ethereum/go-ethereum && make devtools (or make all)

### Run Tests
The test suites are written in Go, so you can run tests from the top level directory with

    invoke test

You can also run individual tests by specifying a subdirectory. For example, here's how you would run the tests for `ethertoken`:

    go test -v ./tests/ethertokentest/

#### Developer Note: Making changes to contract
If you've changed a contract in this repo, you will need to regenerate the `.abi` and `.bin` files tied to that file. These files are used to regenerate the go API bindings for the contract. You finally need to generate the external vyper files. For example, let's suppose that you've changed the `Voting.vy` contract. Then you would run these commands

##### Precompilation
Vyper contracts can have their ABI and BIN files created by using the `invoke` task dedicated to "precompiling"

    invoke precompile

##### Compilation
Once ABI and BIN files are created we need to create Golang bindings so that the spec suite can run. There is a dedicated
task for compilation of these:

    invoke compile (note this will also trigger precompilation)

##### Build
Aside from precompilation of Vyper files and  compilation of Golang bindings we generate external inferfaces for reference purposes. As a convenience there is a
single command to perforom all three:

    invoke build (runs precompilation, compilation and external interface construction)

#### Tasks file
You can see all the commands available to the invoke task runner in `tasks.py`
