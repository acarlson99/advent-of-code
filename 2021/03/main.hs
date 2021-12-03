{-# LANGUAGE BinaryLiterals #-}

import AOC
import Data.List (transpose)

or10 :: [Bool] -> Bool
or10 s = num0 > num1
  where
    num0 = length $ filter (== False) s
    num1 = length $ filter (== True) s

f :: [[Bool]] -> Int
f xs = boolListToDec blst * boolListToDec olst
  where
    blst = map or10 (transpose xs)
    olst = map not blst

runRound :: Bool -> ([[Bool]], Int) -> ([[Bool]], Int)
runRound dorev (xs, bi) = (filter (\a -> fnm == (a !! bi)) xs, bi + 1)
  where
    fnm' = or10 $ map (!! bi) xs
    fnm = if dorev then not fnm' else fnm'

f' :: [[Bool]] -> Int
f' xs = oxNum * coNum
  where
    parse = ((boolListToDec . head . fst . head . dropWhile ((> 1) . length . fst)) .) . iterate
    oxNum = parse (runRound False) (xs, 0)
    coNum = parse (runRound True) (xs, 0)

main :: IO ()
main = do
  inp <- map (map cToB) . lines <$> readFile "input.txt"
  print $ f inp
  print $ f' inp
