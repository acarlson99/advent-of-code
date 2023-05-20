#!/usr/bin/env python3

import fileinput

col = lambda i,l: list(map(lambda x: x[i], l))
row = lambda i,l: l[i]

# count nums less than n, stopping at first number greater
def f(ls,n):
    for i in range(len(ls)):
        if ls[i] >= n:
            return i+1
    return len(ls)

if __name__=='__main__':
    lines = []
    for line in fileinput.input():
        lines.append(line.strip())
    lines = list(map(lambda x: list(map(int,x)), lines))
    t=0
    lss=[]
    for i in range(len(lines)):
        ls=[]
        for j in range(len(lines[i])):
            n = int(lines[i][j])
            c = col(j,lines)
            ca = [-1] + c[:i]
            cb = [-1] + c[i+1:]
            r = row(i,lines)
            ra = [-1] + r[:j]
            rb = [-1] + r[j+1:]
            # print(ra,n,rb)
            # print(ca,n,cb)
            if n > max(ca) or n > max(cb) or n > max(ra) or n > max(rb):
                t+=1
                ls.append(n)
            else:
                ls.append(0)
        lss.append(ls)
    # print()
    # list(map(print,lss))
    print(t)

    lss=[]
    for i in range(len(lines)):
        ls=[]
        for j in range(len(lines[i])):
            n = int(lines[i][j])
            c = col(j,lines)
            ca = c[:i][::-1] # reverse lefthand side for tree perspective
            cb = c[i+1:]
            r = row(i,lines)
            ra = r[:j][::-1]
            rb = r[j+1:]
            res = f(ca,n) * f(cb,n) * f(ra,n) * f(rb,n)
            # print('u',ca,'d',cb,'l',ra,'r',rb)
            # print(i,j,'=','u',f(ca,n),'d',f(cb,n),'l',f(ra,n),'r',f(rb,n))
            ls.append(res)
            # print(res)
        lss.append(ls)
    # print(lss)
    print(max(map(max,lss)))
