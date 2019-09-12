import sys
from collections import namedtuple
from manticore.ethereum import ManticoreEVM
from manticore.utils import config
from manticore.ethereum.abi import ABI
#from manticore.core.smtlib import solver

from manticore.core.smtlib.solver import Z3Solver
solver = Z3Solver.instance()

from initialize_market import initialize
from constants import ONE_ETH

consts = config.get_group("evm")
consts.oog = "ignore"

Account = namedtuple('Account', ['name', 'account', 'initial_supply'])

################ Helpers #######################

def init_user(m, market, name, initial_balance):
    user_account = m.create_account(balance=initial_balance)
    print(f"[+] Creating a {name} account {hex(user_account.address)}")

    market.ether_token_contract.deposit(value=initial_balance, caller=user_account)
    print(f"[+] {name} deposed its ether")
    market.ether_token_contract.approve(market.reserve_contract.address, initial_balance, caller=user_account)
    print(f"[+] {name} called approve")

    return Account(name=name,
                   account=user_account,
                   initial_supply=initial_balance)

def calls_support(market, calls):
    """

    :param market:
    :param calls: list(User, amount)
    :return:
    """
    for (user, amount) in calls:
        market.reserve_contract.support(amount, caller=user.account)
        print(f"[+] {user.name} called support")


def check_balance(m, market, user):
    market.ether_token_contract.balanceOf(user.account.address)
    for state in m.ready_states:
        world = state.platform
        balance_of_tx = world.human_transactions[-1]
        last_return = ABI.deserialize("uint", balance_of_tx.return_data)

        m.generate_testcase(state, name=f"BugFound_{user.name}", only_if=last_return >= user.initial_supply)


################ Scenarios #######################

def scenario1():
    m = ManticoreEVM()

    market = initialize(m,
                        parameters={'pr_fl': 10 ** 6},
                        market_owner_token_balance=10 ** 21)

    bob = init_user(m, market, 'bob', 100 * ONE_ETH)
    eve = init_user(m, market, 'eve', 100 * ONE_ETH)

    calls_support(market,
                  [(eve, 1 * ONE_ETH),
                   (bob, 10 * ONE_ETH)])

    market.reserve_contract.withdraw(caller=eve.account)
    print("[+] Eve called support")

    check_balance(m, market, eve)

    print(f"[+] Look for results in {m.workspace}")

def scenario2():
    m = ManticoreEVM()

    market = initialize(m,
                        parameters={'pr_fl': 10 ** 6},
                        market_owner_token_balance=10 ** 21)

    bob = init_user(m, market, 'bob', 200 * ONE_ETH)
    eve = init_user(m, market, 'eve', 100 * ONE_ETH)

    calls_support(market,
                  [(bob, 1 * ONE_ETH),
                   (eve, 1 * ONE_ETH),
                   (bob, 100 * ONE_ETH)])

    market.reserve_contract.withdraw(caller=eve.account)
    print("[+] Eve called support")

    check_balance(m, market, eve)

    print(f"[+] Look for results in {m.workspace}")

def scenario3():
    m = ManticoreEVM()

    market = initialize(m,
                        parameters={'pr_fl': 10 ** 6},
                        market_owner_token_balance=10 ** 21)

    alice = init_user(m, market, 'alice', 100 * ONE_ETH)
    bob = init_user(m, market, 'bob', 10000 * ONE_ETH)
    eve = init_user(m, market, 'eve', 100 * ONE_ETH)

    calls_support(market,
                  [(bob, 1 * ONE_ETH),
                   (alice, 1 * ONE_ETH),
                   (alice, 20 * ONE_ETH),
                   (bob, 50 * ONE_ETH),
                   (alice, 40 * ONE_ETH),
                   (eve, 100 * ONE_ETH),
                   (bob, 1000 * ONE_ETH)])

    market.reserve_contract.withdraw(caller=eve.account)
    print("[+] Eve called support")

    check_balance(m, market, eve)

    print(f"[+] Look for results in {m.workspace}")


SCENARIOS = {
    '1': scenario1,
    '2': scenario2,
    '3': scenario3,
}


################ Main #######################

def main():
    if len(sys.argv) != 2 or not sys.argv[1] in SCENARIOS:
        print('Wrong arguments. Usage ./script.py scenario_id')
        print(f'Valid ids: {list(SCENARIOS.keys())}')

    SCENARIOS[sys.argv[1]]()

if __name__ == "__main__":
    main()