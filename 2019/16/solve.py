#!/usr/bin/env python3

import fileinput
import numpy as np

def run_phase(inList, pattern, times):
    if times <= 0:
        return inList
    outList = []
    patternIdx = 0
    usePattern = pattern
    iterCnt = 1
    # for n in inList:
    for n in range(len(inList)):
        # print(n)
        patternIdx = 1
        num = 0
        for n2 in range(len(inList)):
            num += (inList[n2] * usePattern[patternIdx])
            patternIdx = (patternIdx + 1) % len(usePattern)
        outList.append(abs(num) % 10)
        usePattern = []
        iterCnt += 1
        for pn in pattern:
            for itn in range(iterCnt):
                usePattern.append(pn)
    return run_phase(outList, pattern, times-1)

line = fileinput.input().__next__()
inList = [int(n) for n in line if n != '\n']

l1 = run_phase(inList, [0,1,0,-1], 100)
print("Part one:", ''.join([str(n) for n in l1[0:8]]))

offset = int(''.join([str(n) for n in inList[0:7]]))
list2 = np.array([n for n in inList] * 10000)
list2 = list2[offset:]

# for i in range(100):
#     total = sum(list2)
#     for j in range(len(list2)):
#         t = total
#         total -= list2[j]
#         list2[j] = abs(t) % 10

for i in range(100):
    list2 = (list2[::-1].cumsum() % 10)[::-1]

print("Part two:", ''.join([str(n) for n in list2[0:8]]))
