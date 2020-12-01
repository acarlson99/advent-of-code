import           Data.Char
import           Data.List.Extra
import           Safe

checkOne :: String -> Bool
checkOne (a : b : c : xs) =
  if (((nextLetter a) == b) && ((nextLetter b) == c) && a < b && b < c)
    then True
    else checkOne $ b : c : xs
checkOne _ = False

checkTwo :: String -> Bool
checkTwo s = foldr ((&&) . not . (`elem` s)) True "iol"

countPairs (a : b : xs) | a == b    = 1 + countPairs xs
                        | otherwise = countPairs $ b : xs
countPairs _ = 0

checkThree :: String -> Bool
checkThree = (>= 2) . countPairs

checkAll = and . ([checkOne, checkTwo, checkThree] <*>) . pure

nextLetter :: Char -> Char
nextLetter c | o `elem` (map (subtract 1 . ord) "iol") = chr $ (+ 2) o
             | o >= (ord 'z')                          = 'a'
             | otherwise                               = chr $ (+ 1) o
  where o = ord c

nextString :: String -> String
nextString s = if newS < s
  then snoc (nextString $ init s) (last newS)
  else newS
  where newS = snoc (init s) $ nextLetter $ last s

nextValidString :: String -> String
nextValidString s =
  findJust (checkAll) $ tail $ scanl (const . nextString) s [1 ..]

inp = "vzbxkghb"

main = do
  let ns = nextValidString inp
  putStrLn ns
  putStrLn $ nextValidString ns
