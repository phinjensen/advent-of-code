import fileinput

total = 0

def score(char):
    if char.islower():
        return ord(char) - ord('a') + 1
    if char.isupper():
        return ord(char) - ord('A') + 27

for line in fileinput.input():
    line = line.rstrip()
    l = len(line)//2
    left, right = set(line[:l]), set(line[l:])
    total += score((left & right).pop())

print(total)
