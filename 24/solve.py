#!/usr/bin/env python3

import numpy as np
import fileinput

def print_fmt(arr):
    for n in arr:
        for c in n:
            if c == 1:
                print("#", end='', flush=True)
            else:
                print(".", end='', flush=True)
        print("")

def check(arr, x, y):
    if y < 0 or y >= len(arr) or x < 0 or x >= len(arr[y]):
        return 0
    return arr[y][x]

def count_neighbors(arr, x, y):
    return check(arr, x+1, y) + check(arr, x-1, y) + check(arr, x, y+1) + check(arr, x, y-1)

def next_cycle(arr):
    new = np.array([[0 for x in range(len(arr[y]))] for y in range(len(arr))], int)
    for y in range(len(arr)):
        for x in range(len(arr[y])):
            n = count_neighbors(arr, x, y)
            if arr[y][x] == 1 and n == 1:
                new[y][x] = 1
            elif arr[y][x] == 0 and (n == 1 or n == 2):
                new[y][x] = 1
            else:
                new[y][x] = 0
    return new

def biodiv(arr):
    pos = 1
    total = 0
    for row in arr:
        for c in row:
            if c == 1:
                total += pos
            pos *= 2
    return total

inArr = []
for line in fileinput.input():
    inArr.append(line.rstrip())

# print(inArr)

height = len(inArr)
width = len(inArr[0])

workingArr = np.array([[0 for x in range(width)] for y in range(height)], int)

for y in range(len(workingArr)):
    for x in range(len(workingArr[y])):
        if inArr[y][x] == '#':
            workingArr[y][x] = 1
        else:
            workingArr[y][x] = 0

bdivs = {}
working = workingArr
while True:
    bd = biodiv(working)
    if bd in bdivs:
        print("Part one:", bd)
        break
    bdivs[bd] = 1
    working = next_cycle(working)
