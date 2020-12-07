type BagMap = [(String, [(String, Int)])]

target = "shiny gold"

bagName :: String -> String -> String
bagName mod clr = mod ++ ' ' : clr

-- 3 muted magenta bags, 3 clear cyan bags
parseNames :: [String] -> [(String, Int)]
parseNames (n : mod : clr : _ : xs) = (bagName mod clr, read n) : parseNames xs
parseNames _ = []

-- shiny lime bags contain 3 muted magenta bags, 3 clear cyan bags.
parseBag :: [String] -> (String, [(String, Int)])
parseBag (mod : clr : "bags" : "contain" : "no" : "other" : "bags" : _) = (bagName mod clr, [])
parseBag (mod : clr : "bags" : "contain" : xs) = (bagName mod clr, parseNames xs)

containsGold :: String -> BagMap -> Bool
containsGold name bs = case lk of
  Just xs -> any (\(a, _) -> (a == target) || a `containsGold` bs) xs
  Nothing -> False
  where
    lk = lookup name bs

partOne inp = sum (map (fromEnum . flip containsGold bs . fst) bs)
  where
    bs = map parseBag inp

countContained name bs = case lk of
  Just xs -> sum $ map (\(s, n) -> n + n * countContained s bs) xs
  Nothing -> 0
  where
    lk = lookup name bs

partTwo inp = countContained target bs
  where
    bs = map parseBag inp

main = do
  inp <- map words . lines <$> readFile "input.txt"
  print $ partOne inp
  print $ partTwo inp
