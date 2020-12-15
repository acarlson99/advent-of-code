import qualified Data.Bifunctor as Bifunctor
import Data.Bits
import Data.Char
import Data.List
import qualified Data.Map as Map

splitOn _ [] = []
splitOn c xs = x : splitOn c (drop (length x) xs)
  where
    x = head xs : takeWhile (not . c) (drop 1 xs)

groupElements :: Ord a => [a] -> [(a, Int)]
groupElements = map (\xs@(x : _) -> (x, length xs)) . group . sort

onlyDigits = takeWhile isDigit . dropWhile (not . isDigit)

constructMask [] = id
constructMask (x : xs)
  | x == 'X' = constructMask xs
  | x == '1' = flip setBit (length xs) . constructMask xs
  | x == '0' = flip clearBit (length xs) . constructMask xs

maskBlock (x, xs) = map (Bifunctor.second f) xs
  where
    f = constructMask x

uniqueMemAddrs = Map.toList . foldl (flip $ uncurry Map.insert) Map.empty

partOne :: [([Char], [(Int, Int)])] -> Int
partOne = sum . map snd . uniqueMemAddrs . concatMap maskBlock

constructMaskHelper :: String -> [Int -> Int]
constructMaskHelper [] = [id]
constructMaskHelper (x : xs)
  | x == '0' = constructMaskHelper xs
  | x == '1' = (flip setBit (length xs) .) <$> constructMaskHelper xs
  | x == 'X' = [(flip setBit (length xs) .), (flip clearBit (length xs) .)] <*> constructMaskHelper xs

constructMask' :: String -> Int -> [Int]
constructMask' xs = (<*>) (constructMaskHelper xs) . pure

maskAssignment :: (Int -> [Int]) -> (Int, Int) -> [(Int, Int)]
maskAssignment mask (addr, v) = zip (mask addr) (repeat v)

maskBlock' :: (String, [(Int, Int)]) -> [(Int, Int)]
maskBlock' (x, xs) = concatMap (maskAssignment (constructMask' x)) xs

partTwo :: [(String, [(Int, Int)])] -> Int
partTwo = sum . map snd . uniqueMemAddrs . concatMap maskBlock'

parseAssignment :: [Char] -> (Int, Int)
parseAssignment xs = (read . onlyDigits $ head x, read $ last x)
  where
    x = words xs

parseInp (x : xs) = (last $ words x, map parseAssignment xs)

main = do
  inp <- map parseInp . splitOn (\a -> take 4 a == "mask") . lines <$> readFile "input.txt"
  print $ partOne inp
  print $ partTwo inp
