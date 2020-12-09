import AOC

helper ps xs = if ts /= [] then helper (head xs : init ps) (tail xs) else head xs
  where
    ts = dropWhile (uncurry (==)) $ twoSum (head xs) ps

partOne inp = helper (reverse $ take 25 inp) (drop 25 inp)

trimArr sm ss x target =
  if sm + x <= target
    then (sm + x, x : ss)
    else
      let n = last ss
          ss' = init ss
       in trimArr (sm - n) ss' x target

rollingSum sm ss xs target
  | sm == target = ss
  | sm + head xs > target = let (sm', ss') = trimArr sm ss (head xs) target in rollingSum sm' ss' (tail xs) target
  | otherwise = rollingSum (sm + head xs) (head xs : ss) (tail xs) target

partTwo inp x = minimum arr + maximum arr
  where
    arr = rollingSum 0 [] inp x

main = do
  inp <- map read . lines <$> readFile "input.txt" :: IO [Int]
  let one = partOne inp
  print one
  print $ partTwo inp one
