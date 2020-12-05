import Control.Applicative
import Data.List

maximum' = foldr1 (\x y -> if x >= y then x else y)

findSeatID ('B' : xs) = (2 ^ length xs) + findSeatID xs
findSeatID ('F' : xs) = findSeatID xs
findSeatID ('R' : xs) = (2 ^ length xs) + findSeatID xs
findSeatID ('L' : xs) = findSeatID xs
findSeatID _ = 0

seatIDs = map findSeatID

partOne = maximum' . seatIDs

partTwo = f . sort . seatIDs
  where
    f (x : xs) = fst $ head $ dropWhile (uncurry (==)) $ zip [x ..] (x : xs)

main = do
  inp <- lines <$> readFile "input.txt"
  print $ partOne inp
  print $ partTwo inp
