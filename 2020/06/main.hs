import AOC
import Data.List

partOne = sum . map (uniq . concat)

partTwo = sum . map (length . foldr intersect ['a' .. 'z'])

main = do
  inp <- splitWhen (== "") . splitWhen (\a -> (a == ' ') || (a == '\n')) <$> readFile "input.txt"
  print $ partOne inp
  print $ partTwo inp
