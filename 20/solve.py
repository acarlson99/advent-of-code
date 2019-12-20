#!/usr/bin/env python3


import fileinput
import networkx as nx

def check(x, y, inArr):
    print(x,y)
    if x < 0 or y < 0 or y >= len(inArr) or x >= len(inArr[y]) or inArr[y][x] != '.':
        return False
    return True

def isalpha(c):
    return ((c >= 'A' and c <= 'Z') or (c >= 'a' and c <= 'z'))

def getName2(x, y, inArr):
    c = inArr[y][x]
    if inArr[y][x] == '.':
        return str(x) + "," + str(y)
    c2 = ""
    if x > 0 and isalpha(inArr[y][x-1]):
        c2 = inArr[y][x-1]
    elif x < len(inArr[y]) and isalpha(inArr[y][x+1]):
        c2 = inArr[y][x+1]
    elif y > 0 and isalpha(inArr[y-1][x]):
        c2 = inArr[y-1][x]
    elif y < len(inArr) and isalpha(inArr[y+1][x]):
        c2 = inArr[y+1][x]
    names = [c, c2]
    names.sort()
    return names[0] + "," + names[1]

def getName(x, y, inArr):
    for x1,y1 in [(1,0), (-1,0), (0,1), (0,-1)]:
        x2 = x + x1
        y2 = y + y1
        if x2 < 0 or y2 < 0 or y2 >= len(inArr) or x2 >= len(inArr[y2]):
            continue
        if isalpha(inArr[y2][x2]):
            return getName2(x2, y2, inArr)
    return str(x) + "," + str(y)

portalMap = {}

inArr = []
for line in fileinput.input():
    inArr.append(list(line[0:len(line)-1]))

for line in inArr:
    print(line)

portals = {}

G = nx.Graph()
for y in range(len(inArr)):
    for x in range(len(inArr[y])):
        if check(x, y, inArr):
            n = getName(x, y, inArr)
            if n not in G:
                G.add_node(getName(x, y, inArr))

for y in range(len(inArr)):
    for x in range(len(inArr[y])):
        if check(x, y, inArr):
            n = getName(x, y, inArr)
            x2 = x + 1
            y2 = y + 1
            if check(x2, y, inArr):
                n2 = getName(x2, y, inArr)
                G.add_edge(n, n2)
            if check(x, y2, inArr):
                n2 = getName(x, y2, inArr)
                G.add_edge(n, n2)

print(len(nx.shortest_path(G, "A,A", "Z,Z")))
print(nx.algorithms.shortest_paths.unweighted.bidirectional_shortest_path(G, "A,A", "Z,Z"))
