#!/usr/bin/env python3

import fileinput

# file 20x20
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

def simplify(x,y):
    a,b = x,y
    while b:
        a,b = b, a % b
    return (x/a,y/a)

bestDict = []
best = 0
bestCoords = (-1,-1)
bestTotals = []
for x1, y1 in asteroids:
    # y2-y1 / x2-1
    localDict = {}
    totals = [set(),set(),set(),set()]
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

        if slope not in totals[idx]:
            totals[idx].add(slope)
            localDict[slope] = ((x2,y2))

    total = sum([len(set(t)) for t in totals])
    if total > best:
        bestTotals = totals
        best = total
        bestCoords = (x1,y1)
        bestDict = localDict

print(best, bestCoords)

# vapIdx = 0
# for lst in bestTotals:
#     if vapIdx + len(lst) < 200:
#         vapIdx += len(lst)
#         continue
#     slopes = [(y/x, (x,y)) for x,y in lst]
#     slopes.sort()
#     print(200-vapIdx)
#     print(slopes[200-vapIdx - 1])

# 302 < ans < 705
# 314
