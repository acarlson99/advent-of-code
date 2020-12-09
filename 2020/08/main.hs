readPosNeg ('+' : xs) = read xs
readPosNeg xs = read xs

-- return (PC, acc)
runOp :: [String] -> (Int, Int) -> (Int, Int)
runOp ["jmp", n] (pc, acc) = (pc + readPosNeg n, acc)
runOp ["nop", _] (pc, acc) = (pc + 1, acc)
runOp ["acc", n] (pc, acc) = (pc + 1, acc + readPosNeg n)

evalArr :: [[String]] -> (Int, Int) -> [Int] -> Int
evalArr is (pc, acc) visited = if pc `elem` visited then acc else evalArr is (pc', acc') (pc : visited)
  where
    op = is !! pc
    (pc', acc') = runOp op (pc, acc)

partOne inp = evalArr inp (0, 0) []

changeOp ["jmp", n] = ["nop", n]
changeOp ["nop", n] = ["jmp", n]
changeOp xs = xs

evalArr' :: [[String]] -> (Int, Int) -> [Int] -> Maybe Int
evalArr' is (pc, acc) visited
  | pc `elem` visited = Nothing
  | pc >= length is = Just acc
  | otherwise = evalArr' is (runOp (is !! pc) (pc, acc)) (pc : visited)

fromMaybe (Just a) = a
fromMaybe Nothing = error "bad"

partTwo inp = fromMaybe $ head $ dropWhile (== Nothing) $ map (\a -> evalArr' a (0, 0) []) rs'
  where
    rs = zip [0 ..] $ repeat inp
    rs' = map (\(n, xs) -> take n xs ++ changeOp (xs !! n) : drop (n + 1) xs) rs

main = do
  inp <- map words . lines <$> readFile "input.txt"
  print $ partOne inp
  print $ partTwo inp
