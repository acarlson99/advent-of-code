arrMin = foldr1 min

arrMax = foldr1 max

split _ "" = []
split c s  = firstWord : (split c rest)
 where
  firstWord = takeWhile (/= c) s
  rest      = drop (length firstWord + 1) s

sideLengths a = [a !! 0 * a !! 1, a !! 0 * a !! 2, a !! 1 * a !! 2]

squareFeetBox :: [Int] -> Int
squareFeetBox a = arrMin b + foldr (\v acc -> acc + 2 * v) 0 b
  where b = sideLengths a

partOne :: [[Int]] -> Int
partOne = foldr (\v acc -> acc + squareFeetBox v) 0

wrapAround :: [Int] -> Int
wrapAround a = foldr (\v acc -> acc + v * 2) 0 a - (2 * arrMax a)

ribonRequired a = (wrapAround a) + (product a)

partTwo :: [[Int]] -> Int
partTwo = foldr (\v acc -> acc + ribonRequired v) 0

main = do
  inp <- map (map read . split 'x') <$> lines <$> readFile "input.txt"
  print $ partOne inp
  print $ partTwo inp
