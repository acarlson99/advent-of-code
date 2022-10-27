#!/bin/sh

# 01
grep -oE '\-*[0-9]+' input.txt | paste -s -d + - | bc

# 02
jq < input.txt | ./script.p6
