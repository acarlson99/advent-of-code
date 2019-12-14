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

# backward chain FUEL
amounts = {}
print(backchain("FUEL", 1, amounts, reactions))
print(amounts)

amounts = {}
cycles = {}
oreAmount = 1000000000000
oreUsed = 0
done = False
# for ii in range(10000000):
fuelGenerated = 0
while not done:
    tmp = backchain("FUEL", 1, amounts, reactions)
    if oreUsed + tmp >= oreAmount:
        print(fuelGenerated)
        break
    oreUsed += tmp
    done = True
    for n in amounts:
        if amounts[n] != 0:
            print(n, amounts[n])
            done = False
            break
    fuelGenerated += 1

print(fuelGenerated)
print("USED", oreUsed)
print("LEFT", oreAmount - oreUsed)

cycleGen = oreAmount // oreUsed
print(cycleGen)
fuelGenerated = fuelGenerated * cycleGen
print(fuelGenerated)
oreUsed = oreUsed * cycleGen

done = False
while not done:
    tmp = backchain("FUEL", 1, amounts, reactions)
    # print(oreUsed)
    if oreUsed + tmp >= oreAmount:
        print(fuelGenerated)
        break
    oreUsed += tmp
    done = True
    for n in amounts:
        if amounts[n] != 0:
            done = False
            break
    fuelGenerated += 1

print(fuelGenerated)

# 63097
