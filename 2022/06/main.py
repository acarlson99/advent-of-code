#!/usr/bin/env python3

import fileinput

if __name__=='__main__':
    lines = []
    for line in fileinput.input():
        lines.append(line.strip())

    s = lines[0]
    print(s)
    for i in range(4,len(s)):
        rl = s[i-4:i]
        if len(set(rl))==4:
            print(rl)
            print(set(rl))
            print(i)
            break
    for i in range(14,len(s)):
        rl = s[i-14:i]
        if len(set(rl))==14:
            print(rl)
            print(set(rl))
            print(i)
            break
