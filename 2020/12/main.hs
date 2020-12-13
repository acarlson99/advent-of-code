turnBoat turn dir = (dir + (turn `div` 90)) `mod` 4

-- 0=east
-- 1=south
-- 2=west
-- 3=north
boatForward n ((x, y), dir)
  | dir == 0 = (x + n, y)
  | dir == 1 = (x, y - n)
  | dir == 2 = (x - n, y)
  | dir == 3 = (x, y + n)

navigateBoat (cmd, n) ((x, y), dir)
  | cmd == 'N' = ((x, y + n), dir)
  | cmd == 'S' = ((x, y - n), dir)
  | cmd == 'E' = ((x + n, y), dir)
  | cmd == 'W' = ((x - n, y), dir)
  | cmd == 'R' = ((x, y), turnBoat n dir)
  | cmd == 'L' = ((x, y), turnBoat (360 - n) dir)
  | cmd == 'F' = (boatForward n ((x, y), dir), dir)

partOne = (\(a, b) -> abs a + abs b) . fst . foldl (flip navigateBoat) ((0, 0), 0)

rotateClockwise (x, y) = (y, - x)

rotatePoint 0 (px, py) = (px, py)
rotatePoint n (px, py) = rotatePoint (n -1) $ rotateClockwise (px, py)

navigateBoat' :: (Char, Int) -> ((Int, Int), (Int, Int)) -> ((Int, Int), (Int, Int))
navigateBoat' (cmd, n) ((bx, by), (wpx, wpy))
  | cmd == 'N' = ((bx, by), (wpx, wpy + n))
  | cmd == 'S' = ((bx, by), (wpx, wpy - n))
  | cmd == 'E' = ((bx, by), (wpx + n, wpy))
  | cmd == 'W' = ((bx, by), (wpx - n, wpy))
  | cmd == 'R' = ((bx, by), rotatePoint (n `div` 90) (wpx, wpy))
  | cmd == 'L' = ((bx, by), rotatePoint (360 - (n `div` 90)) (wpx, wpy))
  | cmd == 'F' = ((bx + wpx * n, by + wpy * n), (wpx, wpy))

partTwo = (\(a, b) -> abs a + abs b) . fst . foldl (flip navigateBoat') ((0, 0), (10, 1))

main = do
  inp <- map (\a -> (head a, read $ tail a)) . lines <$> readFile "input.txt" :: IO [(Char, Int)]
  print $ partOne inp
  print $ partTwo inp
