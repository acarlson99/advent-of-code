#!/usr/bin/env python3

import fileinput

# return list of (x,y) pairs
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

# array of arrays of (x,y) pairs
arrs = []
for line in fileinput.input():
    arrs.append(make_tuples(line))

# filter for intersections
seen = {}
repeated = set()
arrnum = 0
for l in arrs:
    for idx,i in enumerate(l):
        if i in seen:
            if seen[i][1] != arrnum:
                repeated.add((i, (idx, seen[i][0] + 1)))
        else:
            seen[i] = (idx + 1, arrnum)
    arrnum += 1

# find best intersection
d1 = 1000000000000
d2 = 1000000000000
for pos, dist in repeated:
    d1 = min(abs(pos[0]) + abs(pos[1]), d1)
    d2 = min(dist[0] + dist[1], d2)

print("Part one:", d1)
print("Part two:", d2)
