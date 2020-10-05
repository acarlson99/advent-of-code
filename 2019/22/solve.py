#!/usr/bin/env python3

import fileinput

def cut_n(n, cards):
    return cards[n:] + cards[0:n]

def deal_inc(inc, cards):
    new = cards.copy()
    ln = len(cards)
    for ii in range(ln):
        newIdx = 0
        new[(ii * inc) % ln] = cards[ii]
    return new

def shuffle_cards(cards, inArr):
    for line in inArr:
        if len(line) <= 0:
            continue
        if line[0] == 'c':
            line = line.split(" ")
            cards = cut_n(int(line[1]), cards)
        elif line[18] == 'k':
            cards = cards[::-1]
        elif line[0] == 'd':
            line = line.split(" ")
            cards = deal_inc(int(line[3]), cards)
    return cards

inArr = []
for line in fileinput.input():
    inArr.append(line.rstrip())

cards = [n for n in range(10007)]
cards = shuffle_cards(cards, inArr)

for i in range(len(cards)):
    if cards[i] == 2019:
        print("Part one:", i)
