# Verification and Testing artifacts for Computable

## Pre-requisites

1. Decompress this archive into the root of the computable repository. We tested this from commit [1502a7993d914047401fb3a7a3ee5d4dcfae2c7a](https://github.com/computablelabs/computable/tree/1502a7993d914047401fb3a7a3ee5d4dcfae2c7a)
2. Install the latest release of [Manticore](https://github.com/trailofbits/manticore).

## Manticore verification

The added manticore directory (`test/verification/manticore/issues`) contains the scripts to check for certain issues. In particular, we created scripts to verify issues:
* TOB-Computable-001
* TOB-Computable-002 
* TOB-Computable-004 
* TOB-Computable-005 
* TOB-Computable-007 
* TOB-Computable-008 
* TOB-Computable-012 
* TOB-Computable-013 
* TOB-Computable-017
* TOB-Computable-018 

### Instructions 

Just execute each python script from the root directory of this repository. For instance:

``` 
$ python3.6 test/verification/manticore/issues/TOB-Computable-017.py 
[+] Creating a market owner account 0x9adc5d1928deb703e5c481c5597542b9373d8079
[+] Creating a MarketToken contract account 0x2890bb02ae0c20b9642dce0000948eed88a9bf34
[+] Creating a ether owner account 0x2452ec3ae4c0e08173665b358b6a7be75a20651e
[+] Creating a EtherToken contract account 0xa18e0bb9250b1eabe78fdc9a310145dada0cda04
[+] Creating a voting owner account 0x36f69e0e499b798d1a8900016a1822fd6f7380c5
[+] Creating a Voting contract account 0xc777f8a832e5c2a4f7876ef7cd9b0055b3d93630
[+] Creating a param owner account 0x1aa7fef3c6c2ece7b8bbd8882836624cbc21c15f
[+] Creating a Parameterizer contract account 0xc72f052a397a64fd8ae4bcac007896178e434294
[+] Creating a reserve owner account 0x870b0c8393ca121563982f71f5879d1b5e9ee148
[+] Creating a Reserve contract account 0x5386bda65f40bc4b3e9b4eaea9e0b34619985da5
[+] Creating a datatrust owner account 0x1a4b771dd2a7089035a16dfce1d189149a4380ae
[+] Creating a Datatrust contract account 0xfef8a6a1790288e8f86924dc66728d6406f637d3
[+] Creating a listing owner account 0x51462d082e7db6b8dc9f5bd205940e75b30569d8
[+] Creating a Listing contract account 0xab6c37abea35cad7293da2923e91a050f5ddda4a
[+] setPrivileged called on Voting
[+] setPrivileged called on MarketToken
[+] setPrivileged called on Datatrust
[+] Approve called
[+] register called
[+] vote called
[+] timestamp increased
[+] registration resolved
[+] unstake vote
[+] Creating a market owner account 0x22261404ad311539658be4f7b6c7e92453d532e4
[+] Deposit called
[+] Approve called
[+] Creating a listing owner account 0x43a76a11a38f25af27e6481c3fa25b6ba5500188
[+] list called
[+] vote called
[+] timestamp increased
[+] setDataHash called
[+] resolveApplication called
[+] unstake vote
[+] Request delivery called
[+] listingAccessed called
[+] claimBytesAccessed called
[+] delivered called
[+] Look for results in /home/gustavo/projects/audit-computable/mcore_erii2wfc
```

After manticore finishes you can check the `mcore_*` directory with .tx files to examine the complete logs of the transactions discovered by Manticore to trigger each bug.
