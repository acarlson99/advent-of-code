#!/usr/bin/env python3

import fileinput

def parse_pair(string):
    output = string.split(" ")
    outQTY = int(output[0])
    # output[1] = output[1][0:len(output[1])-1]
    return (outQTY, output[1].rstrip())

# return fuel needed
def backchain(name, amountNeeded, amounts, graph):
    if name == "ORE":
        return amountNeeded
    needed = graph[name]
    total = 0
    for qty, name in needed[1]:
        res = 0
        if name == "ORE":
            total += qty
            continue
        if name not in amounts:
            amounts[name] = 0
        while amounts[name] < qty:
            res += backchain(name, qty, amounts, graph)
            amounts[name] += graph[name][0]
        amounts[name] -= qty
        total += res
    return total

# K = {output (qty, INPUTS)}
# INPUT = (qty, name)
reactions = {}
for line in fileinput.input():
    inputs, output = line.split(" => ")
    inputs = inputs.split(", ")
    output = parse_pair(output)
    inList = []
    for inp in inputs:
        inList.append(parse_pair(inp))
    reactions[output[1]] = (output[0], inList)
    # add rule to map of rules

# simplify graph to float values for each resource
def simplify_tree(root, graph, new):
    if root in new:
        return new[root]
    total = 0
    needed = graph[root]
    for qty, name in needed[1]:
        val = simplify_tree(name, graph, new)
        total += val * qty
    new[root] = total / needed[0]
    return new[root]

# backward chain FUEL
amounts = {}
print("Part one:", backchain("FUEL", 1, amounts, reactions))

# part two
new = {"ORE": 1}
simplify_tree("FUEL", reactions, new)
print("Part two:", int(1000000000000 // new["FUEL"]))

# 3 N => 1 F
# 1 J => 2 N

# F = N*3
# N = (1*J) / 2

# amount = (REQ * f(name)) / PRODUCED

# 63097
