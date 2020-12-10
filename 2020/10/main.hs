import Data.List

diffList xs = zipWith (-) (tail xs) xs

partOne :: [Int] -> Int
partOne inp = length (filter (== 1) xs) * length (filter (== 3) xs)
  where
    xs = diffList $ 0 : inp ++ [last inp + 3]

{-

-- would work, but slow as shit
countBacktrack (x : xs) last target
  | last >= target = 1
  | x - last > 3 = 0
  | otherwise = curr + next
  where
    curr = countBacktrack xs x target
    next = countBacktrack xs last target
countBacktrack _ last target = fromEnum (last >= target)

partTwoSlow inp = countBacktrack inp 0 (last inp)

-}

-- [(value, paths)]
evalList xs [] = xs
evalList xs (i : is) = evalList ((i, c) : xs) is
  where
    c = sum . map snd . filter ((<= 3) . abs . subtract i . fst) $ take 3 xs

partTwo = snd . head . evalList [(0, 1)]

main = do
  inp <- sort . map read . lines <$> readFile "input.txt" :: IO [Int]
  print $ partOne inp
  print $ partTwo inp
