import fileinput

def score(char):
    if char.islower():
        return ord(char) - ord('a') + 1
    if char.isupper():
        return ord(char) - ord('A') + 27

group = set()
total = 0

for i, line in enumerate(fileinput.input()):
    if i % 3 == 0:
        group = set(line)
    line = line.rstrip()
    group = group & set(line)
    if i % 3 == 2:
        total += score(group.pop())

print(total)
