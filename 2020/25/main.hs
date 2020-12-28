import Data.List

nextLoop subNum n = (n * subNum) `mod` 20201227

genLoops sn = iterate (nextLoop sn) 1

main = do
  inp <- map read . lines <$> readFile "input.txt" :: IO [Int]
  let (Just a) = elemIndex (last inp) $ genLoops 7
  print $ flip (!!) a $ genLoops (head inp)

-- < 14643366
