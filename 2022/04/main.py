#!/usr/bin/env python3

import fileinput

if __name__=='__main__':
    lines = []
    for line in fileinput.input():
        lines.append(line.strip())
    t1 = 0
    t2 = 0
    for line in lines:
        l = line.split(',')
        r1 = l[0]
        a = r1.split('-')
        s1 = set(range(int(a[0]), 1+int(a[1])))
        r2 = l[1]
        b = r2.split('-')
        s2 = set(range(int(b[0]), 1+int(b[1])))
        print(s1,s2)
        if (s1|s2 == s1 or s1|s2 == s2):
            print('FOUND',s1,s2)
            t1 += 1
        if (s1&s2 != set()):
            t2 += 1
    print(t1)
    print(t2)
