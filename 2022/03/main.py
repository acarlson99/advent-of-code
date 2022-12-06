#!/usr/bin/env python3

from itertools import accumulate
import fileinput

def score(c):
    sc = 0;
    if c>='A' and c<='Z':
        sc+=26;
    sc += 1+ord(c.lower())-ord('a')
    return sc

def p1(lines):
    total=0
    for line in lines:
        n=len(line)//2
        s1=set(line[:n])
        s2=set(line[n:])
        s=s1&s2
        for c in s:
            sc = score(c)
            total+=sc
    print(total)

def p2(lines):
    total=0
    working=[]
    for line in lines:
        working.append(line)
        if len(working)==3:
            badge = list(accumulate(map(set,working), lambda x,y: x&y))[-1].pop()
            total+=score(badge)
            working=[]
    print(total)

if __name__=='__main__':
    lines = []
    for line in fileinput.input():
        lines.append(line.strip())
    p1(lines)
    p2(lines)

