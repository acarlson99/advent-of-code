#!/usr/bin/env python3

import fileinput

def findReal(n):
    extraFuel = int(n / 3) - 2
    if extraFuel <= 0:
        return (n)
    return (n + findReal(extraFuel))

totalA = 0
totalB = 0
for line in fileinput.input():
    n = int(int(line) / 3) - 2
    totalA += n
    totalB += findReal(n)

print("Part one:", totalA)
print("Part two:", totalB)
