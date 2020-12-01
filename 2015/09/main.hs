import           Data.List
import qualified Data.Map                      as M

-- insertConn m (x, "to", b, "=", n) = do
--   M.insert x ++ b v $ M.insert b ++ x v m
--   where v = read n

-- f :: [String] -> [(String, Int)] -> [(String, Int)]
f (x : "to" : y : "=" : n : []) acc = (x ++ y, v) : (y ++ x, v) : acc
  where v = read n

destNames :: [[String]] -> [String]
destNames = nub . foldr (\(x : _ : y : _) a -> x : y : a) []

possiblePaths = permutations . destNames

scorePath m (x : y : ys) = case M.lookup (x ++ y) m of
  Just n  -> n + scorePath m (y : ys)
  Nothing -> error "CITY NOT FOUND"
scorePath _ _ = 0

constructMap :: [[String]] -> M.Map String Int
constructMap = M.fromList . foldr f []

pathScores inp = map (scorePath m) pp
 where
  pp = possiblePaths inp
  m  = constructMap inp

main = do
  s <- map words <$> lines <$> readFile "input.txt"
  let scores = pathScores s
  print $ minimum scores
  print $ maximum scores
