#!/usr/bin/env python3

import fileinput

def sumDict(d):
    if d['total']!=0:
        return d['total']
    for k in d.keys():
        if k=='total' or k=='..':
            continue
        if type(d[k]) == type(1):
            # int
            d['total']+=d[k]
        else:
            # subdir
            d['total']+=sumDict(d[k])
    return d['total']

def mkNewDir(wd):
    nd = {'total':0, '..':wd}
    return nd

def handleCD(n,base,wd):
    if n not in wd:
        wd[n]=mkNewDir(wd)
    return wd[n]

def getDirsSizeLessThan(d,n,t=0):
    if d['total']<n:
        t+=d['total']
    for k in d.keys():
        if k=='total' or k=='..' or type(d[k])==type(1):
            continue
        t+=getDirsSizeLessThan(d[k],n)
    return t

def getDirsSizeGreaterThan(d,n,ol=[]):
    if d['total']>=n:
        ol.append(d['total'])
    for k in d.keys():
        if k=='total' or k=='..' or type(d[k])==type(1):
            continue
        getDirsSizeGreaterThan(d[k],n,ol)
    return ol

if __name__=='__main__':
    lines = []
    for line in fileinput.input():
        lines.append(line.strip())
    d = {'total':0}
    wd = d
    for line in lines:
        spl = line.split(' ')
        if spl[1]=='cd':
            wd=handleCD(spl[2],d,wd)
        elif spl[0]=='dir':
            wd[spl[1]] = mkNewDir(wd)
        elif spl[0].isdigit():
            wd[spl[1]] = int(spl[0])
    sumDict(d)
    print(getDirsSizeLessThan(d,100000))

    # p2
    available = 70000000
    reqSpace = 30000000
    unused = available-d['total']
    spaceToFree = reqSpace-unused
    dirs = getDirsSizeGreaterThan(d,spaceToFree)
    print(min(dirs))
