decode ('\\' : 'x' : _ : _ : xs) = (4 + a, 1 + b) where (a, b) = decode xs
decode ('"'                : xs) = (1 + a, 0 + b) where (a, b) = decode xs
decode ('\\' : '\\'        : xs) = (2 + a, 1 + b) where (a, b) = decode xs
decode ('\\' : '"'         : xs) = (2 + a, 1 + b) where (a, b) = decode xs
decode (x                  : xs) = (1 + a, 1 + b) where (a, b) = decode xs
decode []                        = (0, 0)

partOne inp = a - b
 where
  (a, b) = foldr helper (0, 0) inp
  helper s (a, b) = (a + a', b + b') where (a', b') = decode s

encodedLen' ('"'  : xs) = 2 + encodedLen' xs
encodedLen' ('\\' : xs) = 2 + encodedLen' xs
encodedLen' (x    : xs) = 1 + encodedLen' xs
encodedLen' []          = 0

encodedLen s = 2 + encodedLen' s

partTwo inp = foldr f 0 inp where f s acc = (encodedLen s) - (length s) + acc

-- encode' ('"'  : xs) = '\\' : '"' : encode' xs
-- encode' ('\\' : xs) = '\\' : '\\' : encode' xs
-- encode' (x    : xs) = x : encode' xs
-- encode' []          = ""

-- encode s = ('"' : encode' s) ++ "\""

-- partTwo inp = foldr f 0 inp
--   where f s acc = length (encode s) - (length s) + acc

main = do
  s <- lines <$> readFile "input.txt"
  print $ partOne s
  print $ partTwo s
