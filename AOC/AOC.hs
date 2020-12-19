module AOC where

import Data.List

splitWhen :: (a -> Bool) -> [a] -> [[a]]
splitWhen _ [] = []
splitWhen c l = x : splitWhen c (drop (length x + 1) l)
  where
    x = takeWhile (not . c) l

-- like splitWhen, but keeps object split.  splitOn (<2) [1,3,1,4,1,5] = [[1,3],[1,4],[1,5]]
splitOn :: (a -> Bool) -> [a] -> [[a]]
splitOn _ [] = []
splitOn c xs = x : splitOn c (drop (length x) xs)
  where
    x = head xs : takeWhile (not . c) (drop 1 xs)

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

readDelim :: (Read a) => [Char] -> String -> [a]
readDelim delims = map read . splitWhen (`elem` delims)

readIntCSV :: String -> [Int]
readIntCSV = readDelim ","
