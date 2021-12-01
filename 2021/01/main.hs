f a = length . filter id $ map (\c -> (uncurry (<) c)) xs
  where xs = zip a (tail a)

f' a = f xs
  where xs = zipWith3 (\a b c -> a+b+c) a (tail a) (tail $ tail a)

main = do
  inp <- map read . lines <$> readFile "input.txt" :: IO [Int]
  print $ f inp
  print $ f' inp
