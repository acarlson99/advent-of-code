module AOC where

import Data.List

splitWhen :: (a -> Bool) -> [a] -> [[a]]
splitWhen _ [] = []
splitWhen c l = x : splitWhen c (drop (length x + 1) l)
  where
    x = takeWhile (not . c) l

groupElements :: Ord a => [a] -> [(a, Int)]
groupElements = map (\xs@(x : _) -> (x, length xs)) . group . sort

uniq :: Ord a => [a] -> Int
uniq = length . groupElements

twoSum :: Int -> [Int] -> [(Int, Int)]
twoSum n ns =
  let ixs = zip [0 ..] ns
   in [ (i, j)
        | (i, x) <- ixs,
          (j, y) <- drop (i + 1) ixs,
          (x + y) == n
      ]
