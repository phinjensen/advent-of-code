import fileinput

totals = []
curr = 0

for line in fileinput.input():
    if line.rstrip() == '':
        totals.append(curr)
        curr = 0
    else:
        curr += int(line.rstrip())

print(sum(sorted(totals)[-3:]))
