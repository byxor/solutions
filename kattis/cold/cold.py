input()
arguments = input().split(" ")
numbers = [int(argument) for argument in arguments]
negatives = [number for number in numbers if number < 0]
print(len(negatives))
