#!/usr/bin/env python3

import fileinput

def tadd(a,b):
    return (a[0]+b[0], a[1]+b[1])

def gidx(g,pos):
    return g[pos[1]][pos[0]]

def doDFS(g,start,expandF,isEndF):
    q=[]
    visited = set()
    q.append((start,1))
    while len(q)>0:
        pos,dist = q.pop(0)
        if isEndF(g,pos):
            return dist
        if pos in visited:
            continue
        visited.add(pos)
        for newPos in expandF(g,pos):
            if newPos not in visited:
                if isEndF(g,newPos):
                    return dist
                q.append((newPos,dist+1))
    return -1

def expandF(g,pos):
    c = gidx(g,pos)
    for d in [(0,1),(0,-1),(1,0),(-1,0)]:
        p = tadd(pos,d)
        if p[0]<0 or p[1]<0 or p[1]>=len(g) or p[0]>=len(g[p[1]]):
            continue
        nc = gidx(g,p)
        if nc=='E' and c=='z':
            yield p
        elif nc=='E':
            continue
        elif (c=='S' and nc=='a') or (ord(nc) - ord(c) <= 1):
            yield p

def expandF2(g,pos):
    c = gidx(g,pos)
    for d in [(0,1),(0,-1),(1,0),(-1,0)]:
        p = tadd(pos,d)
        if p[0]<0 or p[1]<0 or p[1]>=len(g) or p[0]>=len(g[p[1]]):
            continue
        nc = gidx(g,p)
        if ord(c) - ord(nc) <= 1:
            yield p

if __name__=='__main__':
    lines = []
    for line in fileinput.input():
        lines.append(line.strip())
    sPos = list(filter(lambda x: x[0]>=0 and x[1]>=0, zip(map(lambda s: s.find("S"), lines), range(len(lines)))))[0]
    print(doDFS(lines,sPos,expandF, (lambda g,pos: gidx(g,pos)=='E')))

    lines[sPos[1]] = lines[sPos[1]].replace('S', 'a')

    ePos = list(filter(lambda x: x[0]>=0 and x[1]>=0, zip(map(lambda s: s.find("E"), lines), range(len(lines)))))[0]
    lines[ePos[1]] = lines[ePos[1]].replace('E', 'z')
    print(doDFS(lines,ePos,expandF2, (lambda g,pos: gidx(g,pos)=='a')))
