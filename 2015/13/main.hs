main = do
  s <- map words <$> lines <$> readFile "input.txt"
  print $ take 10 s
