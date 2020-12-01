import           Data.Array.IO

type LightArray a = IOArray (Int, Int) a

size = 1000

parseCoords :: String -> (Int, Int)
parseCoords a = (read x, read y) where (x, y) = split a ','

lightOp arr op coord = do
  n <- readArray arr coord
  writeArray arr coord $ op n
  return ()

modArray :: LightArray a -> (a -> a) -> [String] -> IO ()
modArray arr op [a, "through", b] = do
  let (x1, y1) = parseCoords a
      (x2, y2) = parseCoords b
  mapM_
    (lightOp arr op)
    [ (x, y)
    | x <- [x1 `min` x2 .. x1 `max` x2]
    , y <- [y1 `min` y2 .. y1 `max` y2]
    ]

operate arr translate ("turn" : "on"  : xs) = modArray arr (translate On) xs
operate arr translate ("turn" : "off" : xs) = modArray arr (translate Off) xs
operate arr translate ("toggle"       : xs) = modArray arr (translate Toggle) xs
operate _   _         _                     = undefined

split s c = (s', s'')
 where
  s'  = takeWhile (/= c) s
  s'' = drop (1 + length s') s

testInp =
  "turn on 1,1 through 10,10\nturn off 3,3 through 5,5\ntoggle 2,2 through 6,6"

allCoords w h = [ (x, y) | x <- [0 .. w], y <- [0 .. h] ]

data Op = On | Off | Toggle

englishTranslate :: Op -> (Bool -> Bool)
englishTranslate On     = const True
englishTranslate Off    = const False
englishTranslate Toggle = not

elvishTranslate :: Op -> (Int -> Int)
elvishTranslate On     = (+ 1)
elvishTranslate Off    = (0 `max`) . subtract 1
elvishTranslate Toggle = (+ 2)

solve initVal translate summarize inp = do
  arr <- newArray ((0, 0), (size, size)) initVal
  mapM_ (operate arr translate) inp
  tf <- mapM (readArray arr) $ allCoords size size
  print $ summarize tf

partOne = solve False englishTranslate (length . filter (== True))

partTwo = solve (0 :: Int) elvishTranslate sum

main = do
  s <- map words <$> lines <$> readFile "input.txt"
  partOne s
  partTwo s
