#!/usr/bin/env python3

from sys import argv

f = open(argv[1], 'r')
contents = f.read().strip()
f.close()

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
        for _ in range(n):
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

def sign(n):
    if n<0:
        return -1
    if n>0:
        return 1
    return 0

def p2(inp):
    ss=set()
    knots = [(0,0) for _ in range(10)]
    for l in inp:
        d,n = l
        for _ in range(n):
            knots[0] = tadd(knots[0], mvtab[d])
            for i in range(len(knots)-1):
                head,tail = knots[i],knots[i+1]
                diff = tsub(head,tail)
                if 2 <= abs(diff[0]) or 2 <= abs(diff[1]):
                    knots[i+1] = tadd(tail, (sign(diff[0]), sign(diff[1])))
            ss.add(knots[-1])
    return len(ss)

print(p2(inp))
