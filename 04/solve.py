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

r = re.compile('(.)\\1\\1+')
def check2(s):
    return check(s) and check(r.sub('', s))

num = 0
num2 = 0
for n in range(range_min, range_max):
    if check(str(n)):
        num += 1
    if check2(str(n)):
        num2 += 1

print("Part one:", num)
print("Part two:", num2)
