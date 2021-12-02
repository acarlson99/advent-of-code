f' (a,b) ("up", n) = (a-n,b)
f' (a,b) ("down", n) = (a+n,b)
f' (a,b) ("forward", n) = (a,b+n)

f :: [(String,Int)] -> Int
f = uncurry (*) . foldl f' (0,0)

g' (h,d,aim) ("up", n) = (h,d,aim-n)
g' (h,d,aim) ("down", n) = (h,d,aim+n)
g' (h,d,aim) ("forward", n) = (h+n, d+n*aim, aim)

g = tripToAns . foldl g' (0,0,0)
  where tripToAns (h,d,_) = h*d

main = do
  inp <- map ((\a -> (head a, read $ last a)) . words) . lines <$> readFile "input.txt" :: IO [(String, Int)]
  print $ f inp
  print $ g inp
