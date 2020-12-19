import Parsing

performOp "+" = (+)
performOp "*" = (*)

-- parse operation (1+1) left associative
parseOpLeftAssoc :: (Monad m, Num b, Alternative m) => m b -> m [Char] -> m b -> m b
parseOpLeftAssoc lhsf opf rhsf = do
  lhs <- lhsf
  addSuffix lhs lhsf opf rhsf
  where
    addSuffix lhs' lhsf' opf' rhsf' = do
      op <- opf'
      rhs <- rhsf'
      maybeAddSuffix (performOp op lhs' rhs) lhsf' opf' rhsf'
    maybeAddSuffix e lhsf' opf' rhsf' =
      addSuffix e lhsf' opf' rhsf' <|> return e

parseParenExpr = char '(' *> parseOperation <* char ')'

parseOperation = parseOpLeftAssoc (integer <|> parseParenExpr) (symbol "+" <|> symbol "*") (integer <|> parseParenExpr)

partOne = sum . map (sum . map fst . parse parseOperation)

parseParenExpr' = char '(' *> parseExpr <* char ')'

parseAdd = parseOpLeftAssoc (integer <|> parseParenExpr') (symbol "+") (integer <|> parseParenExpr')

parseMul = parseOpLeftAssoc (parseAdd <|> parseParenExpr' <|> integer) (symbol "*") (parseAdd <|> parseParenExpr' <|> integer)

parseExpr :: Parser Int
parseExpr =
  parseMul
    <|> parseAdd

partTwo = sum . map (sum . map fst . parse parseExpr)

main = do
  inp <- lines <$> readFile "input.txt"
  print $ partOne inp
  print $ partTwo inp
