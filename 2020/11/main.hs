import qualified Data.Set as S

{-

L && neighbors == 0 = #
# && neighbors >= 4 = L
otherwise = id

-}

-- groupElements :: Ord a => [a] -> [(a, Int)]
-- groupElements = map (\xs@(x : _) -> (x, length xs)) . group . sort

------------------------------------------

surroundingCoords :: Int -> Int -> (Int, Int) -> [(Int, Int)]
surroundingCoords w h (x, y) = filter (\(a, b) -> a >= 0 && b >= 0 && a < w && b < h) $ [(x + x', y + y') | x' <- [-1 .. 1], y' <- [-1 .. 1], (x', y') /= (0, 0)]

countAdjacent w h fullT = sum . map (fromEnum . (`elem` fullT)) . surroundingCoords w h

stepCell w h fullT x coord
  | x == emptyC && adj == 0 = fullC
  | x == fullC && adj >= 4 = emptyC
  | otherwise = x
  where
    adj = countAdjacent w h fullT coord

step :: Int -> Int -> S.Set (Int, Int) -> S.Set (Int, Int) -> S.Set (Int, Int)
step w h floorT fullT =
  S.fromList $
    [ (x, y)
      | x <- [0 .. w -1],
        y <- [0 .. h -1],
        (x, y) `notElem` floorT,
        stepCell w h fullT (if (x, y) `elem` fullT then fullC else emptyC) (x, y) == fullC
    ]

emptyC = 'L'

fullC = '#'

floorC = '.'

parseInp :: [[Char]] -> (S.Set (Int, Int), S.Set (Int, Int))
parseInp inp = (S.fromList . map fst $ filter ((== floorC) . snd) xs, S.fromList . map fst $ filter ((== fullC) . snd) xs)
  where
    xs = concatMap (\(y, s) -> zipWith (\x c -> ((x, y), c)) [0 ..] s) $ zip [0 ..] inp

coordToChar floorT sittingT coord
  | coord `elem` floorT = floorC
  | coord `elem` sittingT = fullC
  | otherwise = emptyC

printMap w h floorT sittingT = mapM_ putStrLn $ [lineToS ([(x, y) | x <- [0 .. h -1]]) | y <- [0 .. w -1]]
  where
    lineToS = map (coordToChar floorT sittingT)

printMap' w h floorT sittingT = do
  printMap w h floorT sittingT
  putStrLn ""

conditionalIterate w h floorT sittingT = if sittingT == sT then length sT else conditionalIterate w h floorT sT
  where
    sT = step w h floorT sittingT

partOne inp = conditionalIterate w h floorT sittingT
  where
    (floorT, sittingT) = parseInp inp
    w = length $ head inp
    h = length inp

main = do
  inp <- lines <$> readFile "input.txt"
  print $ partOne inp
