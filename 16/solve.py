#!/usr/bin/env python3

import fileinput

def run_phase(inList, pattern, times):
    if times <= 0:
        return inList
    outList = []
    patternIdx = 0
    usePattern = pattern
    iterCnt = 1
    for n in inList:
        patternIdx = 1
        num = 0
        for n2 in inList:
            num += (n2 * usePattern[patternIdx])
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
# print(inList)
x = run_phase(inList, [0,1,0,-1], 100)
print("Part one:", ''.join([str(n) for n in x[0:8]]))

# offset = int(''.join([str(n) for n in inList[0:7]]))
# print(offset)
# print(''.join([str(n) for n in x[offset:offset+8]]))
