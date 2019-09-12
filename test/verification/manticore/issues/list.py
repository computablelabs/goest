from initialize_market import Market
from vote import vote, unstake_votes

def list(m, market:Market, listing_hash, listing_owner, backend_user):
    '''

    :param market:
    :param listing_hash:
    :param listing_owner:
    :param voters_yes: list of voter account
    :param voters_no:  list of voter account
    :param voted_by:
    :return:
    '''
    market.listing_contract.list(listing_hash, caller=listing_owner)
    print("[+] list called")
    vote(m, market, listing_hash, [market.market_token_owner])

    listing_data = bytes.fromhex("deadbeef")
    market.datatrust_contract.setDataHash(listing_hash, listing_data, caller=backend_user)
    print("[+] setDataHash called")

    market.listing_contract.resolveApplication(listing_hash, caller=market.listing_owner)

    print("[+] resolveApplication called")
    unstake_votes(market, listing_hash, [market.market_token_owner])

