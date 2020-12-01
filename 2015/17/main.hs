import           Data.List

target :: Int
target = 150

findAll = filter (\a -> (sum a) == target) . subsequences

minLen inp = minimum $ map length inp

main = do
  inp <- map read <$> lines <$> readFile "input.txt"
  let a = findAll inp
  print $ length $ a
  print $ length $ filter (\x -> (minLen a) == (length x)) a
