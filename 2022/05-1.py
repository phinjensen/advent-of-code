import fileinput
import re

start = True
moves = False
stacks = []

for line in fileinput.input():
    line = line.rstrip("\n")
    if start:
        stacks = [[] for _col in range((len(line)+1)//4)]
        start = False
    if line == "":
        moves = True
        continue
    if moves:
        count, source, dest = map(int, re.findall(r"\d+", line))
        source -= 1
        dest -= 1
        for i in range(count):
            stacks[dest].append(stacks[source].pop())
    else:
        for i in range(len(stacks)):
            crate = line[i*4+1]
            if crate != ' ' and not crate.isnumeric():
                stacks[i].insert(0, crate)

print("".join([stack[len(stack)-1] for stack in stacks]))
