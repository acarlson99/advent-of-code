#!/usr/bin/env python3

import fileinput

if __name__=='__main__':
    lines = []
    for line in fileinput.input():
        lines.append(line.strip())
