import           Data.List

vowels = "aeiou"

checkVowel = (>= 3) . length . (`intersect` vowels)

checkDouble (a : b : xs) = if a == b then True else checkDouble (b : xs)
checkDouble _            = False

badPairs = ["ab", "cd", "pq", "xy"]

checkBadPairs :: String -> Bool
checkBadPairs (a : b : xs) =
  if [a, b] `elem` badPairs then False else checkBadPairs (b : xs)
checkBadPairs _ = True

checks = [checkVowel, checkDouble, checkBadPairs]

checkGoodPair (a : b : xs) = (isInfixOf [a, b] xs) || (checkGoodPair (b : xs))
checkGoodPair _            = False

checkSeparatedPair (a : b : c : xs) =
  if a == c then True else checkSeparatedPair (b : c : xs)
checkSeparatedPair _ = False

check1 s = foldr1 (&&) (map ($ s) checks)

check2 s = (checkGoodPair s) && (checkSeparatedPair s)

countTrue = length . filter (== True)

solve f inp = print $ countTrue $ map f inp

main = do
  inp <- lines <$> readFile "input.txt"
  solve check1 inp
  solve check2 inp
