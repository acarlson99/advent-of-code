#!/usr/bin/dyalogscript

inp←⍎¨(⎕UCS 10)(≠⊆⊢)⊃⎕NGET '/tmp/input.txt'

p1←{+/(⌽1↓⌽⍵) < (1↓⍵)}

⎕←p1 inp

rot←{⍵↓⌽(2-⍵)↓⌽⍺}
p2←p1 {⊃+/(⍵ rot 0) (⍵ rot 1) (⍵ rot 2)}

⎕←p2 inp

