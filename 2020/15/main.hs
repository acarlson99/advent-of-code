import qualified Data.IntMap as M
import Data.List

inp :: [Int]
inp = [12, 20, 0, 6, 1, 17, 7]

testInp :: [Int]
testInp = [0, 3, 6]

runRound :: M.IntMap Int -> Int -> Int -> (M.IntMap Int, Int)
runRound mp lastNum roundNum = case mp M.!? lastNum of
  Nothing -> (M.insert lastNum (roundNum -1) mp, 0)
  Just x -> let x' = roundNum - 1 - x in (M.insert lastNum (roundNum -1) mp, x')

runRounds xs n = snd $ foldl' (\(mp, lastNum) roundNum -> runRound mp lastNum roundNum) (mp, last xs) [length xs + 1 .. n]
  where
    mp = M.fromList . init $ zip xs [1 ..]

main = do
  print $ runRounds inp 2020
  print $ runRounds inp 30000000
