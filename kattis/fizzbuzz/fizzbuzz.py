arguments = input().split(" ")

x = int(arguments[0])
y = int(arguments[1])
n = int(arguments[2])

for i in range(1, n + 1):
    isFizz = (i % x == 0)
    isBuzz = (i % y == 0)
    if isFizz and isBuzz:
        print("FizzBuzz")
    elif isFizz:
        print("Fizz")
    elif isBuzz:
        print("Buzz")
    else:
        print(i)
