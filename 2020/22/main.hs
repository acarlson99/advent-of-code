import AOC
import qualified Data.Map as M

runRound :: ([Int], [Int]) -> ([Int], [Int])
runRound (x : xs, y : ys) = if x > y then (xs ++ [x, y], ys) else (xs, ys ++ [y, x])

scoreDecks a b = sum $ zipWith (*) [1 ..] $ reverse deck
  where
    deck = if null a then b else a

partOne [a, b] = uncurry scoreDecks $ until (\(a, b) -> null a || null b) runRound (map read $ tail a, map read $ tail b)

runRound' :: ([Int], [Int], M.Map String Bool) -> ([Int], [Int], M.Map String Bool)
runRound' ([], ys, m) = ([], ys, m)
runRound' (xs, [], m) = (xs, [], m)
runRound' (x : xs, y : ys, m)
  | hs `M.member` m = (x : xs, y : ys, m)
  | (x <= length xs) && (y <= length ys) =
    let (xs', _, _) = runRound' (take x xs, take y ys, m')
     in if null xs' then runRound' (xs, ys ++ [y, x], m') else runRound' (xs ++ [x, y], ys, m')
  | x > y = runRound' (xs ++ [x, y], ys, m')
  | x < y = runRound' (xs, ys ++ [y, x], m')
  where
    hs = show (x : xs, y : ys)
    m' = M.insert (show (x : xs, y : ys)) True m

partTwo [a, b] = (\(a, b, _) -> scoreDecks a b) $ runRound' (map read $ tail a, map read $ tail b, M.empty)

main = do
  inp <- splitWhen (== "") . lines <$> readFile "input.txt"
  print $ partOne inp
  print $ partTwo inp
