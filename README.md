# Goest

## Golang testing and sandbox for the Computable Protocol.
[![Build Status](https://travis-ci.org/computablelabs/goest.svg?branch=master)](https://travis-ci.org/computablelabs/goest)

### Installation
- Python 3.6 is required. Hey, look at that next bullet...
- We strongly recommend you use a dedicated virtual environment for goest. Here `virtualenv` is used, but you could use `venv` as well:

    virtualenv -p python3.6 --no-site-packages <path-to-virtual-env-dir>
    source <path-to-virtual-env-dir>/bin/activate

- Make sure your Go environment is configured. See [go installation directions](https://golang.org/doc/install). Note that it is idiomatic go to have your
  `$GOPATH` and `$GOROOT` exported to your shell. Of particular import is `$GOPATH` which is used by some of the `invoke` tasks in this project. _This_ dev,
  for example, has:

    export GOROOT=~/.go
    export GOPATH=~/go

- The test suite uses `geth` utilities and libraries. You can install `geth` with `go get`:

    go get -d github.com/ethereum/go-ethereum

Alternatively, if you'd prefer, you can install Geth from source. See [instructions](https://github.com/ethereum/go-ethereum/wiki/Installing-Geth).
Once installed, ensure that `abigen` is available to your go environment by _at least_ building the `devtools`:

    cd $GOPATH/src/github.com/ethereum/go-ethereum
    make devtools (or make all)

### Run Tests
The test suites are written in Go, so you can run tests with the `invoke` task (from root directory):

    invoke test

You can also run individual tests by specifying a package. For example, here's how you would run the tests for `ethertoken`:

    go test -v ./tests/ethertokentest/

#### Developer Note: Making changes to contract
If you've changed a contract in this repo, you will need to regenerate the `.abi` and `.bin` files for that file. These files are used to regenerate the go API bindings that power the spec suite.
Finally, you should regenerate the external interface file for any changed contract. Here are the invoke tasks for performing these actions:

##### Precompilation
Vyper contracts can have their ABI and BIN files created by using the `invoke` task dedicated to "precompiling"

    invoke precompile

##### Compilation
Once ABI and BIN files are created we need to create Golang bindings so that the spec suite can run. There is a dedicated task for compilation of these:

    invoke compile (note this will also trigger precompilation)

##### Build
Aside from precompilation of Vyper files and  compilation of Golang bindings we generate external inferfaces for reference purposes. As a convenience there is a
single command to perforom all three:

    invoke build (runs precompilation, compilation and external interface construction)

#### Tasks file
You can see all the commands available to the invoke task runner in `tasks.py`
