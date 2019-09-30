from invoke import task

@task
def precompile_ether_token(c):
    print('Compiling Vyper EtherToken into ABI and binary.')
    c.run('vyper -f abi contracts/EtherToken.vy > contracts/ethertoken/ethertoken.abi && vyper -f bytecode contracts/EtherToken.vy > contracts/ethertoken/ethertoken.bin')

@task(precompile_ether_token)
def precompile_market_token(c):
    print('Compiling Vyper MarketToken into ABI and binary.')
    c.run('vyper -f abi contracts/MarketToken.vy > contracts/markettoken/markettoken.abi && vyper -f bytecode contracts/MarketToken.vy > contracts/markettoken/markettoken.bin')

@task(precompile_market_token)
def precompile_voting(c):
    print('Compiling Vyper Voting into ABI and binary.')
    c.run('vyper -f abi contracts/Voting.vy > contracts/voting/voting.abi && vyper -f bytecode contracts/Voting.vy > contracts/voting/voting.bin')

@task(precompile_voting)
def precompile_parameterizer(c):
    print('Compiling Vyper Parameterizer into ABI and binary.')
    c.run('vyper -f abi contracts/Parameterizer.vy > contracts/parameterizer/parameterizer.abi && vyper -f bytecode contracts/Parameterizer.vy > contracts/parameterizer/parameterizer.bin')

@task(precompile_parameterizer)
def precompile_reserve(c):
    print('Compiling Vyper Reserve into ABI and binary.')
    c.run('vyper -f abi contracts/Reserve.vy > contracts/reserve/reserve.abi && vyper -f bytecode contracts/Reserve.vy > contracts/reserve/reserve.bin')

@task(precompile_reserve)
def precompile_datatrust(c):
    print('Compiling Vyper Datatrust into ABI and binary.')
    c.run('vyper -f abi contracts/Datatrust.vy > contracts/datatrust/datatrust.abi && vyper -f bytecode contracts/Datatrust.vy > contracts/datatrust/datatrust.bin')

@task(precompile_datatrust)
def precompile_listing(c):
    print('Compiling Vyper Listing into ABI and binary.')
    c.run('vyper -f abi contracts/Listing.vy > contracts/listing/listing.abi && vyper -f bytecode contracts/Listing.vy > contracts/listing/listing.bin')

@task(precompile_listing)
def precompile(c):
    print('Smart contracts precompiled.')

@task
def compile_ether_token(c):
    print('Compiling EtherToken ABI and binary into .go')
    c.run('$GOPATH/bin/abigen -abi contracts/ethertoken/ethertoken.abi -pkg ethertoken -type EtherToken -out contracts/ethertoken/ethertoken.go -bin contracts/ethertoken/ethertoken.bin')

@task(compile_ether_token)
def compile_market_token(c):
    print('Compiling MarketToken ABI and binary into .go')
    c.run('$GOPATH/bin/abigen -abi contracts/markettoken/markettoken.abi -pkg markettoken -type MarketToken -out contracts/markettoken/markettoken.go -bin contracts/markettoken/markettoken.bin')

@task(compile_market_token)
def compile_voting(c):
    print('Compiling Voting ABI and binary into .go')
    c.run('$GOPATH/bin/abigen -abi contracts/voting/voting.abi -pkg voting -out contracts/voting/voting.go -bin contracts/voting/voting.bin')

@task(compile_voting)
def compile_parameterizer(c):
    print('Compiling Parameterizer ABI and binary into .go')
    c.run('$GOPATH/bin/abigen -abi contracts/parameterizer/parameterizer.abi -pkg parameterizer -out contracts/parameterizer/parameterizer.go -bin contracts/parameterizer/parameterizer.bin')

@task(compile_parameterizer)
def compile_reserve(c):
    print('Compiling Reserve ABI and binary into .go')
    c.run('$GOPATH/bin/abigen -abi contracts/reserve/reserve.abi -pkg reserve -out contracts/reserve/reserve.go -bin contracts/reserve/reserve.bin')

@task(compile_reserve)
def compile_datatrust(c):
    print('Compiling Datatrust ABI and binary into .go')
    c.run('$GOPATH/bin/abigen -abi contracts/datatrust/datatrust.abi -pkg datatrust -out contracts/datatrust/datatrust.go -bin contracts/datatrust/datatrust.bin')

@task(compile_datatrust)
def compile_listing(c):
    print('Compiling Listing ABI and binary into .go')
    c.run('$GOPATH/bin/abigen -abi contracts/listing/listing.abi -pkg listing -out contracts/listing/listing.go -bin contracts/listing/listing.bin')

@task(precompile, compile_listing)
def compile(c):
    print ('Smart contracts compiled')

@task
def compile_ether_token_external(c):
    print('Creating EtherToken external interface')
    c.run('vyper -f external_interface contracts/EtherToken.vy > contracts/ether_token_external.vy')

@task(compile_ether_token_external)
def compile_market_token_external(c):
    print('Creating MarketToken external interface')
    c.run('vyper -f external_interface contracts/MarketToken.vy > contracts/market_token_external.vy')

@task(compile_market_token_external)
def compile_voting_external(c):
    print('Creating Voting external interface')
    c.run('vyper -f external_interface contracts/Voting.vy > contracts/voting_external.vy')

@task(compile_voting_external)
def compile_parameterizer_external(c):
    print('Creating Parameterizer external interface')
    c.run('vyper -f external_interface contracts/Parameterizer.vy > contracts/parameterizer_external.vy')

@task(compile_parameterizer_external)
def compile_reserve_external(c):
    print('Creating Reserve external interface')
    c.run('vyper -f external_interface contracts/Reserve.vy > contracts/reserve_external.vy')

@task(compile_reserve_external)
def compile_datatrust_external(c):
    print('Creating Datatrust external interface')
    c.run('vyper -f external_interface contracts/Datatrust.vy > contracts/datatrust_external.vy')

@task(compile_datatrust_external)
def compile_listing_external(c):
    print('Creating Listing external interface')
    c.run('vyper -f external_interface contracts/Listing.vy > contracts/listing_external.vy')

@task(compile, compile_listing_external)
def build(c):
    print('Contracts compiled. External interfaces created')

@task
def test(c):
    c.run('go test ./tests/...')
