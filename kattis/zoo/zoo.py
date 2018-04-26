from sys import stdin

l = 1
while True:
  n = int(input())
  if n == 0:
    break
  aminalz = {}
  for i in range(n):
    animal = stdin.readline().lower().strip()
    animal_split = animal.split(" ")
    real_animal = animal_split[-1]
    if real_animal in aminalz:
      aminalz[real_animal] = aminalz[real_animal] + 1
    else:
      aminalz[real_animal] = 1
  print("List " + str(l) + ":")
  for aminal in sorted(aminalz):
    print(aminal + " | " + str(aminalz[aminal]))
  l += 1
