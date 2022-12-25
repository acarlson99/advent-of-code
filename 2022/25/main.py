#!/usr/bin/env python3

import fileinput

from functools import *

def reduceF(acc,c):
    acc *= 5
    if c=='=':
        acc -= 2
    elif c=='-':
        acc -= 1
    if c.isnumeric():
        acc += int(c)
    return acc

snafToInt = lambda x: reduce(reduceF,x,0)

def conv(n):
    if n==4:
        return '-'
    elif n==3:
        return '='
    else:
        return str(n)

def intToSnaf(n):
    chars = []
    carry = None

    while n>0:
        m = n % 5
        if carry is not None:
            m += snafToInt(carry)
            carry = None
        n = n // 5
        if m>=3:
            carry = '1'
            m = m%5
        chars.append(conv(m))

    if carry is not None:
        chars.append(carry)
    return ''.join(chars[::-1])

if __name__=='__main__':
    lines = []
    t = 0
    for line in fileinput.input():
        line=line.strip()
        lines.append(line)
        t += snafToInt(line)
        # print(line,snafToInt(line))
    print(intToSnaf(t))
