#!/usr/bin/env python3

import fileinput
import re
import sys

line = fileinput.input().__next__()

# split by w*h into individual layers
def split_layers(line, width, height):
    ret = []
    for i in range(len(line)):
        if i//(width*height) >= len(ret):
            ret.append([])
        ret[i // (width * height)].append(int(line[i]))
    return ret

layers = split_layers(line[0:len(line)-1], 25, 6)

min_0 = 10000000
mlist = []
for l in layers:
    m = l.count(0)
    if m < min_0:
        min_0 = m
        mlist = l
print("Part one:", mlist.count(1) * mlist.count(2))

# compress layers into arr
arr = []
for i in range(len(layers[0])):
    for l in layers:
        if l[i] != 2:
            if l[i] == 1:
                # █
                arr.append(u"\u2588")
            else:
                arr.append(' ')
            break

print("Part two:")
i = 0
while i < len(arr):
    for c in arr[i:i+25]:
        print(c, end='')
    print('\n', end='')
    i += 25
