from constants import ONE_ETH
from initialize_market import Market
from vote import vote, unstake_votes

def register_backend(m, market:Market, voters):

    backend_user = m.create_account(balance = 1 * ONE_ETH)

    backend_url = "my_backend_url"
    backend_url_hash = bytes.fromhex("59cda2aae91169ada6b8881e546e9c97187a2e6506d09d17e1fc8c3697e89eca")
    market.datatrust_contract.register(backend_url, caller=backend_user)

    print(f"[+] register called")

    vote(m, market, backend_url_hash, voters)

    market.datatrust_contract.resolveRegistration(backend_url_hash, caller=backend_user)

    print(f"[+] registration resolved")

    unstake_votes(market, backend_url_hash, voters)

    return backend_user
