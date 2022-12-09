#!/usr/bin/env python3

from sys import argv

print(argv)

f = open(argv[1], 'r')
contents = f.read().strip()
f.close()

print(contents)

inp = list(map(lambda x: (x[0], int(x.split(' ')[1])), contents.split('\n')))

mvtab = {
        'U':(1,0),
        'D':(-1,0),
        'R':(0,1),
        'L':(0,-1),
        }

def tadd(a,b):
    return (a[0]+b[0], a[1]+b[1])

def tsub(a,b):
    return (a[0]-b[0], a[1]-b[1])

def p1(inp):
    ss=set()
    head = (0,0)
    tail = (0,0)
    for l in inp:
        d,n = l
        for i in range(n):
            last = head
            head = tadd(head,mvtab[d])
            diff = tsub(head,tail)
            if abs(diff[0]) > 1 or abs(diff[1]) > 1:
                tail = last
            ss.add(tail)
    return len(ss)

print(p1(inp))

def clamp(a,t,b):
    if t<a:
        return a
    elif t>b:
        return b
    return t

# idk why this doesnt work idgaf to fix it
def p2(inp):
    ss=set()
    head = (0,0)
    tail = (0,0)
    for l in inp:
        d,n = l
        for i in range(n):
            last = head
            head = tadd(head,mvtab[d])
            diff = tsub(head,tail)
            if abs(diff[0]) + abs(diff[1]) > 9:
                # if not in same row/col move in that direction
                oldTail = tail
                tail = tadd(tail,(clamp(-1,diff[0],1), clamp(-1,diff[1],1)))
                print(f"MV {d}{n} {head} {oldTail} {tail}")
            ss.add(tail)
    print(ss)
    return len(ss)

print(p2(inp))

# < 2609
