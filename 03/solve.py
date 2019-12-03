#!/usr/bin/env python3

import fileinput

def make_tuples(line):
    a = []
    x = 0
    y = 0
    for inst in line.split(","):
        if inst[0] == 'U':
            for n in range(int(inst[1::])):
                y -= 1
                a.append((x,y))
        if inst[0] == 'D':
            for n in range(int(inst[1::])):
                y += 1
                a.append((x,y))
        if inst[0] == 'L':
            for n in range(int(inst[1::])):
                x -= 1
                a.append((x,y))
        if inst[0] == 'R':
            for n in range(int(inst[1::])):
                x += 1
                a.append((x,y))
    return a

arrs = []
for line in fileinput.input():
    arrs.append(make_tuples(line))

seen = set()
repeated = set()
for l in arrs:
    for i in set(l):
        if i in seen:
            repeated.add(i)
        else:
            seen.add(i)

rep = map(lambda x: abs(x[0]) + abs(x[1]), repeated)

print("Part one:", min(rep))
