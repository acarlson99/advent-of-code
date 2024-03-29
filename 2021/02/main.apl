#!/usr/bin/dyalogscript
inp←(⎕UCS 10)(≠⊆⊢)⊃⎕NGET '/tmp/input.txt'

forward←{(1⊃⍺) (2⊃⍺ + ⍵)}
up←{(1⊃⍺ - ⍵) (2⊃⍺)}
down←{⍺ up -⍵}
p1 ← {×/⊃{⍎⍕⍵⍺}/⍵, (0 0)}
⎕←p1 inp

forward←{(1⊃⍺+⍵×3⊃⍺) (2⊃⍺ + ⍵) (3⊃⍺)}
up←{(1⊃⍺) (2⊃⍺) (3⊃⍺ - ⍵)}
down←{⍺ up -⍵}
p2 ← {×/2⍴⊃{⍎⍕⍵⍺}/⌽(0 0 0), ⍵}
⎕←p2 inp

inp←'forward 5' 'down 5' 'forward 8' 'up 3' 'down 8' 'forward 2'
