import           Data.Word

-- To continue, please consult the code grid in the manual.  Enter the code at row 2981, column 3075.

seed :: Word64
seed = 20151125
pMod = 33554393
pAdd = 252533

-- polygonalNum :: (RealFrac a, Integral b) => a -> b
polygonalNum n = n' * (n' + 1) `div` 2 + 1 where n' = n - 1

polygonalNumCoord (r, c) = polygonalNum (c + r - 1) + r - 1

-- generateCodes n = n : generateCodes (n * 252533 `rem` 33554393)
generateCodes n = scanl (\n _ -> n * pAdd `rem` pMod) n [1 ..]

-- val at (r, c) == f (r - c + 1) + c - 1
