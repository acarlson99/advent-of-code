import qualified Data.Map.Strict               as M
import           Data.Bits
import           Data.Word
import           Data.Char
import           System.Unsafe

data Term = TConst Word16
          | TVar String

instance Show Term where
  show (TConst n) = show n
  show (TVar   s) = s

data Gate = GUnary (Word16 -> Word16) Term
          | GBinary (Word16 -> Word16 -> Word16) Term Term

instance Show Gate where
  show (GUnary _ t   ) = show t
  show (GBinary _ a b) = (show a) ++ " OP " ++ (show b)

type GateMap = M.Map String Gate

parseTerm s@(s1 : _) | isAlpha s1 = TVar s
                     | isDigit s1 = TConst $ read s
                     | otherwise  = error "TERM BAD"
parseTerm _ = error "TERM SHORT"

fixF f a b = f a $ coerce b

-- getOP :: String -> (Word16 -> Word16 -> Word16)
getOp "AND"    = (.&.)
getOp "OR"     = (.|.)
getOp "LSHIFT" = fixF shiftL
getOp "RSHIFT" = fixF shiftR
getOp _        = error "OP BAD"

createGate :: [String] -> Gate
createGate (x         : "->" : rhs) = GUnary id $ parseTerm x
createGate ("NOT" : x : "->" : rhs) = GUnary complement $ parseTerm x
createGate (x : op : y : "->" : rhs) =
  GBinary (getOp op) (parseTerm x) (parseTerm y)

insertExpr xs m = M.insert (last xs) (createGate xs) m

evalTerm :: GateMap -> Term -> (Word16, GateMap)
evalTerm m (TConst t) = (t, m)
evalTerm m (TVar   t) = evalMap m t

evalGate :: GateMap -> Gate -> (Word16, GateMap)
evalGate m (GUnary f a   ) = (f res, m') where (res, m') = evalTerm m a
evalGate m (GBinary f a b) = (f ra rb, m'')
 where
  (ra, m' ) = evalTerm m a
  (rb, m'') = evalTerm m' b

evalMap :: GateMap -> String -> (Word16, GateMap)
evalMap m s = case M.lookup s m of
  Just v -> (res, M.insert s (wrapVal res) m')
   where
    (res, m') = evalGate m v
    wrapVal   = (GUnary id) . TConst
  Nothing -> error $ "TERM NOT FOUND: " ++ s

showMap m = mapM_ print $ M.toList $ M.map (fst . evalGate m) m

partOne inp = do
  let m = foldr insertExpr M.empty inp
  fst $ evalMap m "a"

-- partTwo :: [String] -> Word16 -> Word16
partTwo inp v = fst $ evalMap (M.insert "b" (GUnary id (TConst v)) m) "a"
  where m = foldr insertExpr M.empty inp

main = do
  s <- map words <$> lines <$> readFile "input.txt"
  let pOne = partOne s
  print pOne
  print $ partTwo s pOne
