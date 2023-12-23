import fileinput

total = 0

for line in fileinput.input():
    group1, group2 = [
            set(range(int(a), int(b)+1)) for a, b in [
                group.split('-') for group in line.split(',')
            ]]
    overlap = group1 & group2
    if len(overlap) > 0:
        total += 1

print(total)
