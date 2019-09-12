def increase_timestamp(m, n):
    for state in m.all_states:
        state.platform._timestamp+=n

