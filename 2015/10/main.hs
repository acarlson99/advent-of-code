import           Data.List

runRound = foldr f "" . group
  where f = (\(x : xs) acc -> (show $ length (x : xs)) ++ [x] ++ acc)

inp = "1113122113"

main = do
  print $ length $ head $ scanr (\a b -> runRound b) inp [1 .. 40]
  print $ length $ head $ scanr (\a b -> runRound b) inp [1 .. 50]
