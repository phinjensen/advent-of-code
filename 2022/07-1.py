import fileinput

stack = ['/']
dirs = {'/': 0}
is_ls = False

for line in fileinput.input():
    line = line.rstrip();
    if line[0] == "$":
        is_ls = False
        command = line[2:4]
        if command == 'cd':
            dest = line[5:]
            if dest == '/':
                stack = ['/']
            elif dest == '..':
                prev = stack.pop()
                dirs[stack[-1]] += dirs[prev]
            else:
                stack.append(dest)
        elif command == 'ls':
            is_ls = True
    else:
        current_dir = stack[-1]
        if line.split(' ')[0].isnumeric():
            dirs.setdefault(current_dir, 0)
            dirs[current_dir] += int(line.split(' ')[0])

print(sum([dirs[i] for i in dirs if dirs[i] < 100000]))
