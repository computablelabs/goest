from initialize_market import Market

def vote(m, market:Market, hash, voters, voted_by=200):
    '''

    :param market:
    :param hash:
    :param voters: list voter account
    :param voted_by:
    :return:
    '''

    for voter in voters:
        market.voting_contract.vote(hash, 1, caller=voter.address)
        print(f"[+] vote called")

    for state in m.ready_states:
        state.platform._timestamp += voted_by
    print(f"[+] timestamp increased")

def unstake_votes(market:Market, hash, voters):
    for voter in voters:
        market.voting_contract.unstake(hash, caller=voter.address)
        print(f"[+] unstake vote")