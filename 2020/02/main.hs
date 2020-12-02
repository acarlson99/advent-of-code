countLetters str c = length $ filter (== c) str

runPart chk = foldr ((+) . f) 0
  where
    f (mn, mx, c, s) = if chk mn mx (head c) s (countLetters s (head c)) then 1 else 0

partOne mn mx _ _ numL = numL >= mn && numL <= mx

partTwo mn mx c s _ = (s !! mn' == c) /= (s !! mx' == c)
  where
    mn' = mn - 1
    mx' = mx -1

main = do
  inp <- map (\s -> (\[a, b, c, d] -> (read a, read b, c, d)) $ words [(if (c == '-') || (c == ':') then ' ' else c) | c <- s]) . lines <$> readFile "input.txt"
  print $ runPart partOne inp
  print $ runPart partTwo inp
