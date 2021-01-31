import qualified Data.IntMap as M
import Data.List

-- import Data.Maybe (isJust)

splitWhen :: (a -> Bool) -> [a] -> [[a]]
splitWhen _ [] = []
splitWhen c l = x : splitWhen c (drop (length x + 1) l)
  where
    x = takeWhile (not . c) l

groupElements :: Ord a => [a] -> [(a, Int)]
groupElements = map (\xs@(x : _) -> (x, length xs)) . group . sort

hashSide :: [Char] -> Int
hashSide ('#' : xs) = (2 ^ length xs) + hashSide xs
hashSide ('.' : xs) = hashSide xs
hashSide [] = 0

topSide :: [String] -> String
topSide = head

bottomSide :: [String] -> String
bottomSide = last

leftSide :: [String] -> String
leftSide = map head

rihgtSide :: [String] -> String
rihgtSide = map last

findRotations xs = [(id .), (reverse .)] <*> [topSide, bottomSide, leftSide, rihgtSide] <*> pure xs

updateMap n x m = if M.notMember x m then M.insert x n m else M.delete x m

insertSides :: [String] -> M.IntMap Int -> M.IntMap Int
-- insertSides (x : xs) m = foldr (\x acc -> M.insert x n acc) m sideHashes
insertSides (x : xs) m = foldr (updateMap n) m sideHashes
  where
    n = read . init . last $ words x
    sideHashes = map hashSide $ findRotations xs

-- NOTE: filtering for (== 4) because corners should have 2 edges * 2 for reflections
partOne = product . map fst . filter ((== 4) . snd) . groupElements . map snd . M.toList . foldr insertSides M.empty

main = do
  inp <- splitWhen (== "") . lines <$> readFile "input.txt"
  print $ partOne inp
