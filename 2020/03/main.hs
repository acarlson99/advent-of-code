safe = '.'

tree = '#'

solve inp (y, x) (dy, dx)
  | y >= length inp = 0
  | otherwise = fromEnum (inp !! y !! (x `mod` length (inp !! y)) == tree) + solve inp (y + dy, x + dx) (dy, dx)

main = do
  inp <- lines <$> readFile "input.txt"
  print $ solve inp (0, 0) (1, 3)
  print $ foldr ((*) . solve inp (0, 0)) 1 [(1, 1), (1, 3), (1, 5), (1, 7), (2, 1)]
