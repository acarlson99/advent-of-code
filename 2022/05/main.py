#!/usr/bin/env python3

import fileinput

def parseLine(l):
    s = l.split(' ')
    num = int(s[1])
    src = int(s[3])
    dest = int(s[5])
    return (num,src,dest)

if __name__=='__main__':
    lines = []
    for line in fileinput.input():
        if line[0]!='m':
            continue
        lines.append(line.strip())

    # p1
    d = {
        1 : list('WBGZRDCV')[::-1],
        2 : list('VTSBCFWG')[::-1],
        3 : list('WNSBC')[::-1],
        4 : list('PCVJNMGQ')[::-1],
        5 : list('BHDFLST')[::-1],
        6 : list('NMWTVJ')[::-1],
        7 : list('GTSCLFP')[::-1],
        8 : list('ZDB')[::-1],
        9 : list('WZNM')[::-1],
    }
    for line in lines:
        num,src,dest = parseLine(line)
        for _ in range(num):
            d[dest].append(d[src].pop())
    for i in sorted(d.keys()):
        print(d[i][-1],end='')
    print('')

    # p2
    d = {
        1 : list('WBGZRDCV')[::-1],
        2 : list('VTSBCFWG')[::-1],
        3 : list('WNSBC')[::-1],
        4 : list('PCVJNMGQ')[::-1],
        5 : list('BHDFLST')[::-1],
        6 : list('NMWTVJ')[::-1],
        7 : list('GTSCLFP')[::-1],
        8 : list('ZDB')[::-1],
        9 : list('WZNM')[::-1],
    }
    for line in lines:
        num,src,dest = parseLine(line)
        d[dest] += d[src][-num:]
        for _ in range(num):
            d[src].pop()
    for i in sorted(d.keys()):
        print(d[i][-1],end='')
    print('')
