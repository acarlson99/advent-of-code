#!/usr/bin/env python3

import re

range_min = 168630
range_max = 718098

def check(s):
    double = 0
    last = 0
    for c in s:
        n = int(c)
        if n < last:
            return False
        elif n == last:
            double += 1
        last = n
    return double >= 1

num = 0
for n in range(range_min, range_max):
    if check(str(n)):
        num += 1

print(num)

print(check("1123"))
print(check("11123"))
