import fileinput

curr = 0
maxi = 0

for line in fileinput.input():
    if line.rstrip() == '':
        if curr > maxi:
            maxi = curr;
        curr = 0
    else:
        curr += int(line.rstrip())

print(maxi)
