from manticore.ethereum import ManticoreEVM
from manticore.utils import config

from utils import increase_timestamp
from initialize_market import initialize, DEFAULT_PARAMETERS
from constants import ONE_ETH

consts = config.get_group("evm")
consts.oog = "ignore"

################ Script #######################

m = ManticoreEVM()

market_parameters = DEFAULT_PARAMETERS
market_parameters['vote_by_d'] = m.make_symbolic_value(name="vote_by_d")
market_parameters['pl'] = m.make_symbolic_value(name="pl")
market_parameters['stk'] = 0

market = initialize(m, parameters=market_parameters)

user_balance = 100 * ONE_ETH
user_account = m.create_account(balance=user_balance)
print("[+] Creating a user account", hex(user_account.address))

user_list_hash = "my_list_hash"
market.listing_contract.list(user_list_hash, caller=user_account)

market.voting_contract.vote(user_list_hash, 1, caller=user_account)

print("[+] Increasing timestamp")
increase_timestamp(m, 32)

market.listing_contract.resolveApplication(user_list_hash, caller=user_account)

m.finalize()
print(f"[+] Look for results in {m.workspace}")
