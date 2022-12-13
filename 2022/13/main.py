#!/usr/bin/env python3

import fileinput
import functools

eq=0
gt=1
lt=-1

def cmp(a,b):
    # print(a,'<=',b)
    al = type(a)==type([1])
    bl = type(b)==type([1])
    if al and bl:
        for an,bn in zip(a,b):
            c = cmp(an,bn)
            if c != eq:
                return c
        return lt if len(a)<len(b) else gt if len(a)>len(b) else eq
    elif not al and not bl:
        return lt if a<b else gt if a>b else eq
    else:
        if not al:
            a = [a]
        else:
            b = [b]
        return cmp(a,b)

if __name__=='__main__':
    lines = []
    w=[]
    for line in fileinput.input():
        l = line.strip()
        if l=='':
            lines.append(w)
            w=[]
        else:
            w.append(eval(l))
    if w!=[]:
        lines.append(w)

    # p1
    total=0
    for i,ls in enumerate(lines):
        res = cmp(ls[0],ls[1])
        if res!=gt:
            # print(i+1)
            total+=i+1
    print(total)

    # p2
    l2=[]
    for ls in lines:
        l2.append(ls[0])
        l2.append(ls[1])
    l2.append([[6]])
    l2.append([[2]])
    arr = sorted(l2, key=functools.cmp_to_key(cmp))
    print((1+arr.index([[2]])) * (1+arr.index([[6]])))
