-- stack.exe runhaskell .\main.hs
-- {-# LANGUAGE FlexibleContexts #-}

f :: Char -> Integer -> Integer
f v acc = if v == '(' then (acc + 1) else (acc - 1)

partOne inp = print $ foldr f 0 inp

f2 []       _   = 0
f2 (x : xs) acc = if acc < 0
  then 0
  else 1 + if x == '(' then f2 xs (acc + 1) else f2 xs (acc - 1)

partTwo inp = print $ f2 inp 0

main = do
  inp <- readFile "input.txt"
  partOne inp
  partTwo inp
