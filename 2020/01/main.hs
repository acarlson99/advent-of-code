{-# LANGUAGE FlexibleContexts #-}

partOne :: [Int] -> Int
partOne xs = (snd . head . filter ((== 2020) . fst)) $ (\a b -> (a + b, a * b)) <$> xs <*> xs

partTwo :: [Int] -> Int
partTwo xs = (snd . head . filter ((== 2020) . fst)) $ (\a b c -> (a + b + c, a * b * c)) <$> xs <*> xs <*> xs

main :: IO ()
main = do
  inp <- map read . lines <$> readFile "input.txt" :: IO [Int]
  print $ partOne inp
  print $ partTwo inp
