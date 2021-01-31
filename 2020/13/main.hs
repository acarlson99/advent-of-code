import Data.Bifunctor (first)

splitWhen :: (a -> Bool) -> [a] -> [[a]]
splitWhen _ [] = []
splitWhen c l = x : splitWhen c (drop (length x + 1) l)
  where
    x = takeWhile (not . c) l

readInt :: String -> Int
readInt = read

partOne (departTime, ids) = (busTime - departTime) * id
  where
    ids' = map readInt $ filter (/= "x") ids
    genSums = \n -> [n, n * 2 ..]
    smallestSumF xs = (head $ dropWhile (< departTime) xs, head xs)
    (busTime, id) = minimum $ map (smallestSumF . genSums) ids'

-- partTwo (t, ids) = fs
--   where
--     fs = map (Data.Bifunctor.first readInt) . filter ((/= "x") . fst) $ zip ids [0 ..]

main = do
  inp <- (\a -> (readInt $ head a, splitWhen (== ',') $ last a)) . lines <$> readFile "input.txt" :: IO (Int, [String])
  print $ partOne inp

--   print $ partTwo inp
