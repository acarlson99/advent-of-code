inp←('1'∘=)⊃⎕NGET '/tmp/input.txt' 1
inp←(⊃⍴inp) (⊃⍴⊃inp) ⍴∊inp
d←{(÷∘2)⊃⍴⍵}

f ← {(>∘(d⍵))+/⍉⍵}
p1 ← {{(2⊥⍵) × (2⊥~⍵)}f⍵}
⎕←p1 inp

shouldRemove ← {=/1, ⍵/⍺}

f ← {(≥∘(d⍵))+/⍉⍵}
v←f inp

scrub ← {⍉{(shouldRemove∘v)/⍵}/⍉⍵}

