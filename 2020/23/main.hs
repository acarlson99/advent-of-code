import qualified Data.IntMap as M
import Data.List

inp :: [Int]
inp = [6, 1, 4, 7, 5, 2, 8, 3, 9]

testInp :: [Int]
testInp = [3, 8, 9, 1, 2, 5, 4, 6, 7]

inp2 n xs = take n $ xs ++ iterate (+ 1) (maximum inp + 1)

inpToMap xs = M.fromList $ zip (last xs : init xs) xs

findDest x m = filter (`M.member` m) $ reverse [0 .. x -1] ++ [length m, length m - 1, length m - 2]

mapToStr' m n xs = if n == 1 && not (null xs) then xs else mapToStr' m n' (n' : xs)
  where
    (Just n') = M.lookup n m

mapToStr m = concatMap show . init . reverse $ mapToStr' m 1 []

move' x m = M.insert ct dn $ M.insert d ch $ M.insert x cn m
  where
    (Just ch) = M.lookup x m
    (Just cm) = M.lookup ch m
    (Just ct) = M.lookup cm m
    (Just cn) = M.lookup ct m
    d = head . filter (\a -> a `notElem` [ch, cm, ct]) $ findDest x m
    (Just dn) = M.lookup d m

move (x, m) = (f $ M.lookup x m', m')
  where
    m' = move' x m
    f (Just a) = a

-- partOne inp = (!! 100) $ iterate move' (3, inpToMap inp)

partOne inp = mapToStr . snd . (!! 100) $ iterate move (head inp, inpToMap inp)

-- why does this segfault
partTwo inp = b * c
  where
    m = snd . (!! 10000000) $ iterate' move (head inp, inpToMap $ inp2 1000000 inp)
    (Just a) = M.lookup 1 m
    (Just b) = M.lookup a m
    (Just c) = M.lookup b m

-- partTwo = product . take 2 . tail . dropWhile (/= 1) . moveRounds 10000000 . inp2 1000000

main = do
  putStrLn $ partOne inp

--   print $ partTwo inp
