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

def get_view(x1,y1):
    pass

bestDict = []
best = 0
bestCoords = (-1,-1)
bestTotals = []
for x1, y1 in asteroids:
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

vapIdx = 0
print(len(bestTotals[0]))
print(len(bestTotals[1]))
print(len(bestTotals[2]))
print(len(bestTotals[3]))
for lst in bestTotals:
    if vapIdx + len(lst) < 10:
        print("SKIP")
        vapIdx += len(lst)
        continue

print("LEN",len(lst))
slopes = [((bestCoords[1]-y/bestCoords[0]-x), (x,y)) for x,y in lst]
print(len(slopes))
slopes.sort()
print(len(slopes))
print(10-vapIdx)
print("0", slopes[0])
print(len(slopes)-1, slopes[len(slopes)-1])
print(10 - vapIdx - len(slopes))
print(slopes[10-vapIdx])

# 302 < ans < 705
# 314
# 512
