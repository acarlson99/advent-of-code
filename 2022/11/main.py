#!/usr/bin/env python3.9

import fileinput
import math

def lcm(a,b):
  return (a * b) // math.gcd(a,b)

if __name__=='__main__':
    lines = []
    w=[]
    wthrow=[-1,-1]
    for line in fileinput.input():
        l=line.strip().split(' ')
        if l[0]=='Monkey':
            w.append(eval(l[1].split(':')[0]))
        elif l[0]=='':
            w.append(wthrow)
            wthrow = [-1,-1]
            lines.append(w)
            w=[]
        elif l[0][0]=='S':
            w.append(eval('['+' '.join(l[2:])+']'))
        elif l[0][0]=='O':
            w.append(''.join(l[3:]))
        elif l[0][0]=='T':
            w.append(eval(l[-1]))
        elif l[1]=='true:':
            wthrow[0] = eval(l[-1])
        elif l[1]=='false:':
            wthrow[1] = eval(l[-1])

    w.append(wthrow)
    lines.append(w)

    # p1
    if True:
        mbusiness = [0 for _ in lines]
        for _ in range(20):
            for m in lines:
                newItems = list(map(lambda old: eval(m[2])//3, m[1]))
                mbusiness[m[0]] += len(newItems)
                for it in newItems:
                    if it%m[3]==0:
                        lines[m[4][0]][1].append(it)
                    else:
                        lines[m[4][1]][1].append(it)
                m[1]=[]
        b = sorted(mbusiness)[-2:]
        print(b[0]*b[1])
    # p2
    else:
        M = math.lcm(*[monk[3] for monk in lines])

        mbusiness = [0 for _ in lines]
        for _ in range(10000):
            for m in lines:
                newItems = list(map(lambda old: eval(m[2])%M, m[1]))
                mbusiness[m[0]] += len(newItems)
                for it in newItems:
                    if it%m[3]==0:
                        lines[m[4][0]][1].append(it)
                    else:
                        lines[m[4][1]][1].append(it)
                m[1]=[]

        list(map(print,lines))
        b = sorted(mbusiness)[-2:]
        print(b[0]*b[1])
