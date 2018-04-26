string = input()

num_whitespace = len([c for c in string if c == "_"])
num_lowercase = len([c for c in string if c.islower()])
num_uppercase = len([c for c in string if c.isupper()])
num_symbols = len([c for c in string if (not c.isalpha()) and (c != "_")])

numbers = [num_whitespace, num_lowercase, num_uppercase, num_symbols]

total = sum(numbers)

ratios = [n/total for n in numbers]

for ratio in ratios:
  print(ratio)
