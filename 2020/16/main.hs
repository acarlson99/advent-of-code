splitWhen :: (a -> Bool) -> [a] -> [[a]]
splitWhen _ [] = []
splitWhen c l = x : splitWhen c (drop (length x + 1) l)
  where
    x = takeWhile (not . c) l

readDelim :: (Read a) => [Char] -> [Char] -> [a]
readDelim delims = map read . splitWhen (`elem` delims)

readIntCSV :: String -> [Int]
readIntCSV = readDelim ","

readRule :: [[Char]] -> [(Int, Int)]
readRule [a, _, b] = [readF a, readF b]
  where
    readF = (\[a, b] -> (a, b)) . readDelim "-"

parseInfile [a, b, c] = (map (readRule . words . last . splitWhen (== ':')) a, readIntCSV $ last b, map readIntCSV $ tail c)

constructRules xs = \a -> (or $ mapM f xs' a, a)
  where
    xs' = concat xs
    f (a, b) = \c -> c >= a && c <= b

partOne rs = sum . map snd . filter (not . fst) . concatMap (map rs')
  where
    rs' = constructRules rs

-- partTwo rs ts os = os'
--   where
--     rs' = constructRules rs
--     os' = filter (all (fst . rs')) os

main = do
  (rules, ticket, others) <- parseInfile . splitWhen (== "") . lines <$> readFile "input.txt"
  print $ partOne rules others

--   print $ partTwo rules ticket others
