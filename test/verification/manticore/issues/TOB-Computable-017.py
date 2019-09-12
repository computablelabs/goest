from manticore.ethereum import ManticoreEVM
from manticore.utils import config
from manticore.ethereum.abi import ABI
#from manticore.core.smtlib import solver

from manticore.core.smtlib.solver import Z3Solver
solver = Z3Solver.instance()

from initialize_market import initialize
from constants import ONE_ETH
from backend_init import register_backend
from list import list

consts = config.get_group("evm")
consts.oog = "ignore"

################ Script #######################

m = ManticoreEVM()

market = initialize(m,
                    parameters={'cost':101,
                                'back_p':33,
                                'maker_p':33})

backend_user = register_backend(m, market, [market.market_token_owner])

user_account = m.create_account(balance= 1 * ONE_ETH)
print("[+] Creating a market owner account", hex(user_account.address))

market.ether_token_contract.deposit(value= 1 * ONE_ETH,
                                    caller=user_account)


print("[+] Deposit called")

market.ether_token_contract.approve(market.datatrust_contract.address,
                                    1 * ONE_ETH,
                                    caller=user_account)
print("[+] Approve called")

listing_hash = bytes.fromhex("41")
listing = m.create_account(balance= 1 * ONE_ETH)
print("[+] Creating a listing owner account", hex(listing.address))

list(m, market, listing_hash, listing, backend_user)

delivery_amount = m.make_symbolic_value(name="delivery_amount")
delivery_hash = bytes.fromhex("41424344")

market.ether_token_contract.balanceOf(market.datatrust_contract.address)

market.datatrust_contract.requestDelivery(delivery_hash,
                                          delivery_amount,
                                          caller=user_account)

market.ether_token_contract.balanceOf(market.datatrust_contract.address)

print("[+] Request delivery called")

market.datatrust_contract.listingAccessed(listing_hash,
                                          delivery_hash,
                                          delivery_amount,
                                          caller=backend_user)

print("[+] listingAccessed called")

market.listing_contract.claimBytesAccessed(listing_hash,
                                           caller=listing)

print("[+] claimBytesAccessed called")

delivery_url = bytes.fromhex("0001")
market.datatrust_contract.delivered(delivery_hash,
                                    delivery_url,
                                    caller=backend_user)

print("[+] delivered called")

market.ether_token_contract.balanceOf(market.datatrust_contract.address)

for state in m.ready_states:
    world = state.platform
    balance_of_tx = world.human_transactions[-1]

    last_return = ABI.deserialize("uint", balance_of_tx.return_data)

    m.generate_testcase(state, name="BugFound", only_if=last_return >= 1)

print(f"[+] Look for results in {m.workspace}")
