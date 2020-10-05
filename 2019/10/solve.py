#!/usr/bin/env python3

import fileinput

def simplify(x,y):
    a,b = x,y
    while b:
        a,b = b, a % b
    return (x/a,y/a)

def get_view(x1,y1):
    slopes = [set(),set(),set(),set()]
    for x2,y2 in asteroids:
        if x1 == x2 and y1 == y2:
            continue
        slope = (x2 - x1, y2 - y1)
        slope = simplify(slope[0], slope[1])

        if x2 >= x1 and y2 >= y1:
            idx = 0
        elif x2 < x1 and y2 >= y1:
            idx = 1
        elif x2 >= x1 and y2 < y1:
            idx = 2
        elif x2 < x1 and y2 < y1:
            idx = 3

        if slope not in slopes[idx]:
            slopes[idx].add(slope)
    return slopes

# read input
yy = 0
# list of x,y pairs
asteroids = []
lines = []
for line in fileinput.input():
    lines.append(line)
    for xx in range(len(line)):
        if line[xx] == '#':
            asteroids.append((xx,yy))
    yy += 1

# part one
best = 0
bestCoords = (-1,-1)
bestSlopes = []
for x1, y1 in asteroids:
    slopes = get_view(x1,y1)
    total = sum([len(set(t)) for t in slopes])
    if total > best:
        bestSlopes = slopes
        best = total
        bestCoords = (x1,y1)

print("Part one:", best)

# part two
vapIdx = 1
for lst in bestSlopes:
    if vapIdx + len(lst) < 200:
        vapIdx += len(lst)
        continue

    # (slope, (rise, run))
    slopes = [((y/x), (x,y)) for x,y in lst]
    slopes.sort()
    rise, run = slopes[200-vapIdx][1]
    x,y = bestCoords
    # fourth quadrant so going up and left
    x -= rise
    y -= run
    while lines[int(y)][int(x)] != '#':
        x -= rise
        y -= run
    print("Part two:", int(x * 100 + y))
