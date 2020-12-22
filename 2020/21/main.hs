import AOC
import Data.Bifunctor
import Data.List

makeSets (is, ags) = (,) <$> ags <*> [is]

parseLine = (\[a, b] -> (words a, map init $ tail $ words b)) . splitWhen (== '(')

reduceAllergens :: [([String], [String])] -> [(String, [String])]
reduceAllergens = map (Data.Bifunctor.second (foldr1 intersect)) . map (\x -> (fst (head x), map snd x)) . groupBy (\a b -> fst a == fst b) . sort . concatMap makeSets

allergens :: [([String], [[Char]])] -> [String]
allergens = nub . concatMap (foldr1 intersect . snd) . map (\x -> (fst (head x), map snd x)) . groupBy (\a b -> fst a == fst b) . sort . concatMap makeSets

partOne :: [([String], [[Char]])] -> Int
partOne inp = length [x | x <- fs, x `notElem` allergens inp]
  where
    fs = concatMap fst inp

knownAllergens :: [(a1, [a2])] -> [(a1, [a2])]
knownAllergens = filter ((== 1) . length . snd)

unknownAllergens :: [(a1, [a2])] -> [(a1, [a2])]
unknownAllergens = filter ((/= 1) . length . snd)

runRound :: [(String, [String])] -> [(String, [String])]
runRound xs = dfs ++ map (\(a, b) -> (a, [b' | b' <- b, b' `notElem` dfs'])) uks
  where
    dfs = knownAllergens xs
    dfs' = concatMap snd dfs
    uks = unknownAllergens xs

runRounds :: [(String, [String])] -> [(String, [String])] -> [(String, [String])]
runRounds as l = if l == n then n else runRounds n as
  where
    n = runRound as

partTwo inp = intercalate "," . map (head . snd) $ sort knowns
  where
    ags = reduceAllergens inp
    knowns = runRounds ags []

main = do
  inp <- map parseLine . lines <$> readFile "input.txt"
  print $ partOne inp
  putStrLn $ partTwo inp
