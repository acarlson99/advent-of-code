import Control.Applicative
import Data.List

maximum' = foldr1 (\x y -> if x >= y then x else y)

findSeatRow ('B' : xs) = (2 ^ length xs) + findSeatRow xs
findSeatRow ('F' : xs) = findSeatRow xs
findSeatRow _ = 0

findSeatCol ('R' : xs) = (2 ^ length xs) + findSeatCol xs
findSeatCol ('L' : xs) = findSeatCol xs
findSeatCol _ = 0

findSeatID (a, b) = findSeatRow a * 8 + findSeatCol b

seatIDs = map findSeatID

partOne = maximum' . seatIDs

partTwo = f . sort . seatIDs
  where
    f (x : xs) = fst . head . dropWhile (uncurry (==)) . getZipList . (((,) <$> ZipList [x ..]) <*>) $ ZipList (x : xs)

main = do
  inp <- map (splitAt 7) . lines <$> readFile "input.txt"
  print $ partOne inp
  print $ partTwo inp
