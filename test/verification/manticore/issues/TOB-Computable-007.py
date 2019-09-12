from manticore.ethereum import ManticoreEVM
from manticore.utils import config

consts = config.get_group("evm")
consts.oog = "ignore"

ONE_ETH = 1000000000000000000
ONE_FINNEY = 1000000000000000
ONE_SZABO = 1000000000000
ONE_GWEI = 1000000000
ONE_MWEI = 1000000
ONE_KWEI = 1000

################ Script #######################

m = ManticoreEVM()

market_owner_account = m.create_account(balance=1000)
print("[+] Creating a market owner account", hex(market_owner_account.address))

#market_owner_balance = m.make_symbolic_value(name="market_owner_balance")
market_owner_balance = 2 ** 32

mt_contract_account = m.solidity_create_contract(open("contracts/MarketToken.vy","r"), crytic_compile_args=dict(), owner=market_owner_account, args=[market_owner_account, market_owner_balance])
print("[+] Creating a MarketToken contract account", hex(mt_contract_account.address))

ether_owner_account = m.create_account(balance=1000)
print("[+] Creating a ether owner account", hex(ether_owner_account.address))

ether_owner_balance = 100 * ONE_ETH

et_contract_account = m.solidity_create_contract(open("contracts/EtherToken.vy","r"), crytic_compile_args=dict(), owner=ether_owner_account, args=[ether_owner_account, ether_owner_balance])
print("[+] Creating a EtherToken contract account", hex(et_contract_account.address))

voting_owner_account = m.create_account(balance=1000)
print("[+] Creating a voting owner account", hex(voting_owner_account.address))

voting_contract_account = m.solidity_create_contract(open("contracts/Voting.vy","r"), crytic_compile_args=dict(), owner=voting_owner_account, args=[mt_contract_account.address])
print("[+] Creating a Voting contract account", hex(voting_contract_account.address))

param_owner_account = m.create_account(balance=1000)
print("[+] Creating a param owner account", hex(param_owner_account.address))

#arg_1 = m.make_symbolic_value(name="pr_fl")
#arg_2 = m.make_symbolic_value(name="spd")
#arg_3 = m.make_symbolic_value(name="list_re")
#arg_4 = m.make_symbolic_value(name="stk")
#arg_5 = m.make_symbolic_value(name="vote_by_d")
#arg_6 = m.make_symbolic_value(name="pl")
#arg_7 = m.make_symbolic_value(name="back_p")
#arg_8 = m.make_symbolic_value(name="maker_p")
#arg_9 = m.make_symbolic_value(name="cost")

arg_1 = ONE_GWEI 
arg_2 = 110
arg_3 = ONE_ETH
arg_4 = 1 #m.make_symbolic_value(name="stk")
arg_5 = 1
arg_6 = 1
arg_7 = 25
arg_8 = 50
arg_9 = ONE_KWEI * 6

args = [voting_contract_account.address, arg_1, arg_2, arg_3, arg_4, arg_5, arg_6, arg_7, arg_8, arg_9]
p_contract_account = m.solidity_create_contract(open("contracts/Parameterizer.vy","r"), crytic_compile_args=dict(), owner=param_owner_account, args=args)
print("[+] Creating a Parameterizer contract account", hex(p_contract_account.address))

reserve_owner_account = m.create_account(balance=1000)
print("[+] Creating a reserve owner account", hex(reserve_owner_account.address))

reserve_contract_account = m.solidity_create_contract(open("contracts/Reserve.vy","r"), crytic_compile_args=dict(), owner=reserve_owner_account, args=[et_contract_account, mt_contract_account, p_contract_account])
print("[+] Creating a Reserve contract account", hex(reserve_contract_account.address))

datatrust_owner_account = m.create_account(balance=1000)
print("[+] Creating a datatrust owner account", hex(datatrust_owner_account.address))

# (ether_token_addr: address, voting_addr: address, p11r_addr: address, res_addr: address
datatrust_contract_account = m.solidity_create_contract(open("contracts/Datatrust.vy","r"), crytic_compile_args=dict(), owner=datatrust_owner_account, args=[et_contract_account, voting_contract_account, p_contract_account, reserve_contract_account])
print("[+] Creating a Datatrust contract account", hex(voting_contract_account.address))

listing_owner_account = m.create_account(balance=1000)
print("[+] Creating a datatrust owner account", hex(listing_owner_account.address))

# market_token_addr: address, voting_addr: address, p11r_addr: address, res_addr: address, data_addr: addressi
listing_contract_account = m.solidity_create_contract(open("contracts/Listing.vy","r"), crytic_compile_args=dict(), owner=listing_owner_account, args=[mt_contract_account, voting_contract_account, p_contract_account, reserve_contract_account, datatrust_contract_account])
print("[+] Creating a Listing contract account", hex(listing_contract_account.address))

# listing: address
datatrust_contract_account.setPrivileged(listing_contract_account, caller=datatrust_owner_account)

# reserve: address, listing: address
mt_contract_account.setPrivileged(reserve_contract_account, listing_contract_account, caller=market_owner_account)

# parameterizer: address, reserve: address, datatrust: address, listing: address
voting_contract_account.setPrivileged(p_contract_account, reserve_contract_account, datatrust_contract_account, listing_contract_account, caller=voting_owner_account)


# backend

backend_balance = 100 * ONE_ETH
backend_account = m.create_account(balance=backend_balance)
print("[+] Creating a backend account", hex(backend_account.address))

et_contract_account.deposit(value=backend_balance, caller=backend_account)

et_contract_account.approve(reserve_contract_account, backend_balance, caller=backend_account) 
mt_contract_account.approve(voting_contract_account, backend_balance, caller=backend_account) 

reserve_contract_account.support(100*ONE_GWEI, caller=backend_account)

backend_url = "my_backend_url"
backend_url_hash = bytes.fromhex("59cda2aae91169ada6b8881e546e9c97187a2e6506d09d17e1fc8c3697e89eca")

datatrust_contract_account.register(backend_url, caller=backend_account)
voting_contract_account.vote(backend_url_hash, 1, caller=backend_account)

for state in m.all_states:
    print("[+] Increasing timestamp")
    state.platform._timestamp+=32

datatrust_contract_account.resolveRegistration(backend_url_hash, caller=backend_account)
datatrust_contract_account.getBackendAddress()

# user

et_contract_account.approve(reserve_contract_account, ether_owner_balance, caller=ether_owner_account) 
mt_contract_account.approve(voting_contract_account, ether_owner_balance, caller=ether_owner_account) 

reserve_contract_account.support(100*ONE_GWEI, caller=ether_owner_account)

user_list_hash = "my_list_hash"
listing_contract_account.list(user_list_hash, caller=ether_owner_account)
voting_contract_account.vote(user_list_hash, 1, caller=ether_owner_account)

for state in m.all_states:
    print("[+] Increasing timestamp")
    state.platform._timestamp+=32

datatrust_contract_account.setDataHash(user_list_hash, backend_url_hash, caller=backend_account)

listing_contract_account.resolveApplication(user_list_hash, caller=ether_owner_account)
listing_contract_account.getListing(user_list_hash)

voting_contract_account.getStake(user_list_hash, ether_owner_account)
# attacker

attacker_balance = 100 * ONE_ETH
attacker_account = m.create_account(balance=attacker_balance)
print("[+] Creating a attacker account", hex(attacker_account.address))

et_contract_account.deposit(value=attacker_balance, caller=attacker_account)
et_contract_account.approve(reserve_contract_account, attacker_balance, caller=attacker_account) 
mt_contract_account.approve(voting_contract_account, attacker_balance, caller=attacker_account) 

#value = m.make_symbolic_value(name="attacker_to_buy")
reserve_contract_account.support(100*ONE_GWEI, caller=attacker_account)

listing_contract_account.challenge(user_list_hash, caller=attacker_account)
voting_contract_account.vote(user_list_hash, 1, caller=attacker_account)

for state in m.all_states:
    print("[+] Increasing timestamp")
    state.platform._timestamp+=32

listing_contract_account.resolveChallenge(user_list_hash, caller=attacker_account)
voting_contract_account.getStake(user_list_hash, ether_owner_account)

m.finalize()
print(f"[+] Look for results in {m.workspace}")
