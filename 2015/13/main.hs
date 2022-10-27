import           Data.List
import qualified Data.Map                      as M

f :: [String] -> (String, Int, String)
f [x, "would", "gain", n, _, _, _, _, _, _, y] = (x, read n, y)
f [x, "would", "lose", n, _, _, _, _, _, _, y] = (x, -1 * read n, y)

dropLast = reverse . drop 1 . reverse

fo3 (a, _, _) = a
mo3 (_, a, _) = a
lo3 (_, _, a) = a

toMap :: (String, Int, String) -> M.Map String (M.Map String Int)
toMap (x, n, y) = M.insert x (M.insert y n M.empty) M.empty

type NameMap = M.Map String (M.Map String Int)

mUnion :: NameMap -> NameMap -> NameMap
mUnion = M.unionWith (\a b -> M.union a b)

scorePerm m l = sum $ map (\(a, b) -> (m M.! b M.! a) + (m M.! a M.! b)) l'
 where
  (a, b) = splitAt 1 l
  l'     = zip l $ b ++ a

main = do
  s <- map (f . words . dropLast) <$> lines <$> readFile "input.txt"
  let guests = nub $ (map lo3 s) ++ map fo3 s
  let m      = foldr (\x acc -> mUnion (toMap x) acc) M.empty s
  print . maximum . map (scorePerm m) $ permutations guests

  let m' = foldr
        (\x a -> mUnion a $ mUnion (toMap ("me", 0, x)) (toMap (x, 0, "me")))
        m
        guests
  print . maximum . map (scorePerm m') $ permutations (["me"] ++ guests)
