from collections import namedtuple
from constants import ONE_GWEI, ONE_ETH, ONE_KWEI

# Default parameter for the Parameterizer contract
DEFAULT_PARAMETERS = {
    'pr_fl': ONE_GWEI,
    'spd': 110,
    'list_re': ONE_ETH,
    'stk': ONE_GWEI,
    'vote_by_d': 100,
    'pl': 50,
    'back_p': 25,
    'maker_p': 50,
    'cost': ONE_KWEI * 6
}

Market = namedtuple('Market', [
    'market_token_owner',
    'market_token_contract',
    'ether_token_owner',
    'ether_token_contract',
    'voting_owner',
    'voting_contract',
    'parameterizer_owner',
    'parameterizer_contract',
    'reserve_owner',
    'reserve_contract',
    'datatrust_owner',
    'datatrust_contract',
    'listing_owner',
    'listing_contract',
])

def initialize(m, user_initial_eth_balance=100, market_owner_token_balance=2**32, parameters=None, only_create=False):
    '''

    :param m:
    :param user_initial_eth_balance: initial wei balance of each user
    :param market_owner_token_balance: initial market tokens balance of the contract's owner
    :param parameters: dictionary holding the Parameterizer contract (see DEFAULT_PARAMETERS)
    :return:
    '''
    parameters = dict() if parameters is None else parameters
    parameters = {**DEFAULT_PARAMETERS, **parameters}

    market_owner_account = m.create_account(balance=user_initial_eth_balance)
    print("[+] Creating a market owner account", hex(market_owner_account.address))

    mt_contract_account = m.solidity_create_contract("contracts/MarketToken.vy",
                                                     owner=market_owner_account,
                                                     args=[market_owner_account,
                                                           market_owner_token_balance])
    print("[+] Creating a MarketToken contract account", hex(mt_contract_account.address))

    ether_owner_account = m.create_account(balance=user_initial_eth_balance)
    print("[+] Creating a ether owner account", hex(ether_owner_account.address))

    ether_owner_balance = 100 * ONE_ETH

    et_contract_account = m.solidity_create_contract("contracts/EtherToken.vy",
                                                     owner=ether_owner_account,
                                                     args=[ether_owner_account,
                                                           ether_owner_balance])


    print("[+] Creating a EtherToken contract account", hex(et_contract_account.address))

    voting_owner_account = m.create_account(balance=user_initial_eth_balance)
    print("[+] Creating a voting owner account", hex(voting_owner_account.address))

    voting_contract_account = m.solidity_create_contract("contracts/Voting.vy",
                                                         owner=voting_owner_account,
                                                         args=[mt_contract_account.address])

    print("[+] Creating a Voting contract account", hex(voting_contract_account.address))

    param_owner_account = m.create_account(balance=user_initial_eth_balance)
    print("[+] Creating a param owner account", hex(param_owner_account.address))

    v_addr = voting_contract_account.address
    pr_fl = parameters['pr_fl']
    spd = parameters['spd']
    list_re = parameters['list_re']
    stk = parameters['stk']
    vote_by_d = parameters['vote_by_d']
    pl = parameters['pl']
    back_p = parameters['back_p']
    maker_p = parameters['maker_p']
    cost = parameters['cost']

    args = [v_addr, pr_fl, spd, list_re, stk, vote_by_d, pl, back_p, maker_p, cost]
    p_contract_account = m.solidity_create_contract("contracts/Parameterizer.vy",
                                                    owner=param_owner_account,
                                                    args=args)

    print("[+] Creating a Parameterizer contract account", hex(p_contract_account.address))

    reserve_owner_account = m.create_account(balance=user_initial_eth_balance)
    print("[+] Creating a reserve owner account", hex(reserve_owner_account.address))

    reserve_contract_account = m.solidity_create_contract("contracts/Reserve.vy",
                                                          owner=reserve_owner_account,
                                                          args=[et_contract_account,
                                                                mt_contract_account,
                                                                p_contract_account])

    print("[+] Creating a Reserve contract account", hex(reserve_contract_account.address))

    datatrust_owner_account = m.create_account(balance=user_initial_eth_balance)
    print("[+] Creating a datatrust owner account", hex(datatrust_owner_account.address))

    # (ether_token_addr: address, voting_addr: address, p11r_addr: address, res_addr: address
    datatrust_contract_account = m.solidity_create_contract("contracts/Datatrust.vy",
                                                            owner=datatrust_owner_account,
                                                            args=[et_contract_account,
                                                                  voting_contract_account,
                                                                  p_contract_account,
                                                                  reserve_contract_account])

    print("[+] Creating a Datatrust contract account", hex(datatrust_contract_account.address))

    listing_owner_account = m.create_account(balance=user_initial_eth_balance)
    print("[+] Creating a listing owner account", hex(listing_owner_account.address))

    # market_token_addr: address, voting_addr: address, p11r_addr: address, res_addr: address, data_addr: addressi
    listing_contract_account = m.solidity_create_contract("contracts/Listing.vy",
                                                          owner=listing_owner_account,
                                                          args=[mt_contract_account,
                                                                voting_contract_account,
                                                                p_contract_account,
                                                                reserve_contract_account,
                                                                datatrust_contract_account])

    if not only_create:

        print("[+] Creating a Listing contract account", hex(listing_contract_account.address))

        voting_contract_account.setPrivileged(p_contract_account.address,
                                          reserve_contract_account.address,
                                          datatrust_contract_account.address,
                                          listing_contract_account.address,
                                          caller=voting_owner_account)

        print("[+] setPrivileged called on Voting")

        mt_contract_account.setPrivileged(reserve_contract_account.address,
                                      listing_contract_account.address,
                                      caller=market_owner_account)

        print("[+] setPrivileged called on MarketToken")

        datatrust_contract_account.setPrivileged(listing_contract_account.address,
                                             caller=datatrust_owner_account)

        print("[+] setPrivileged called on Datatrust")

        mt_contract_account.approve(voting_contract_account.address,
                                market_owner_token_balance,
                                caller=market_owner_account)
        print("[+] Approve called")

    market = Market(
        market_token_owner=market_owner_account,
        market_token_contract=mt_contract_account,
        ether_token_owner=ether_owner_account,
        ether_token_contract=et_contract_account,
        voting_owner=voting_owner_account,
        voting_contract=voting_contract_account,
        parameterizer_owner=param_owner_account,
        parameterizer_contract=p_contract_account,
        reserve_owner=reserve_owner_account,
        reserve_contract=reserve_contract_account,
        datatrust_owner=datatrust_owner_account,
        datatrust_contract=datatrust_contract_account,
        listing_owner=listing_owner_account,
        listing_contract=listing_contract_account
    )

    return market
