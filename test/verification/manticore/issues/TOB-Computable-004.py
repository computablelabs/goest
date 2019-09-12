from manticore.ethereum import ManticoreEVM
from manticore.utils import config
from manticore.core.smtlib.expression import BoolAnd, BoolNot

from initialize_market import initialize, DEFAULT_PARAMETERS

consts = config.get_group("evm")
consts.oog = "ignore"

################ Script #######################

m = ManticoreEVM()

market_parameters = DEFAULT_PARAMETERS

market = initialize(m, parameters=market_parameters, only_create=True)

arg_1 = m.make_symbolic_value(name="mt_arg_1")
arg_2 = m.make_symbolic_value(name="mt_arg_2")

# reserve: address, listing: address
market.market_token_contract.setPrivileged(arg_1, arg_2, caller=market.market_token_owner)

for state in m.ready_states: 
    m.generate_testcase(state, name="BugFound", only_if= BoolAnd(arg_1 == 0, BoolNot (arg_2 == 0)))
    m.generate_testcase(state, name="BugFound", only_if= BoolAnd(BoolNot(arg_1 == 0), arg_2 == 0))

print(f"[+] Look for results in {m.workspace}")
