from t9_sequence_generator import *

def generate_output(test_case, message):
    t9_sequence = generate_t9_sequence_from(message)
    return "Case #{}: {}".format(test_case, t9_sequence)
