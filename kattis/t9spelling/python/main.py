from sys import stdin
from test_case import *

number_of_cases = int(input())

for test_case in range(1, number_of_cases + 1):
    message = stdin.readline()
    if message.endswith("\n"):
        message = message[:-1]
    output = generate_output(test_case, message)
    print(output)
