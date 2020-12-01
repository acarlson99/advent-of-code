import           Data.List

applyMove (a, b) v | v == '^' = (a, b + 1)
                   | v == 'v' = (a, b - 1)
                   | v == '<' = (a - 1, b)
                   | v == '>' = (a + 1, b)

findCoords = scanl applyMove (0, 0)

smush = length . (map length) . group . sort

partOne = smush . findCoords

sepCoords (a : b : xs) = (a : as, b : bs) where (as, bs) = sepCoords xs
sepCoords []           = ([], [])
sepCoords xs           = (xs, [])

partTwo inp = smush $ (findCoords a) ++ (findCoords b)
  where (a, b) = sepCoords inp

main = do
  inp <- readFile "input.txt"
  print $ partOne inp
  print $ partTwo inp
