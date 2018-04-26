from character_sequence_fetcher import *
from sequence_combiner import *

def generate_t9_sequence_from(message):
    if message == "":
        return ""
    else:
        t9_sequence = ""
        previous_sequence = None
        current_sequence = None
        for character in message:
            current_sequence = fetch_sequence_for_character(character)
            t9_sequence += combine_sequences(previous_sequence, current_sequence)
            previous_sequence = current_sequence
        return t9_sequence
