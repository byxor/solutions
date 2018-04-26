def combine_sequences(previous, current):
    if previous is None:
        return current
    elif previous[-1] == current[0]:
        return " " + current
    else:
        return current
