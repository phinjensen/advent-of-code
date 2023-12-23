import sys

lines = [line.rstrip() for line in sys.stdin]
o2_options = lines[:]
co2_options = lines[:]
for bit in range(len(lines[0])):
    o2_count = sum([1 if i[bit] == '1' else -1 for i in o2_options])
    co2_count = sum([1 if i[bit] == '1' else -1 for i in co2_options])
    o2_common = '1' if o2_count >= 0 else '0'
    co2_common = '1' if co2_count >= 0 else '0'
    o2_options = [line for line in o2_options if line[bit] == o2_common]
    co2_options = [line for line in co2_options if line[bit] != co2_common]
    print("o2: ", o2_options)
    print("co2: ", co2_options)
