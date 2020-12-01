import           Data.Digest.Pure.MD5
import           Data.ByteString.Lazy.Internal
import           Data.List

inp = "ckczppom"

f numZ = isPrefixOf $ replicate numZ '0'

check n s = (f n) . show . md5 . packChars . (s ++) . show

findNum n start s = length $ takeWhile (== False) $ map (check n s) [start ..]

main = do
  s <- head <$> lines <$> readFile "input.txt"
  let one = findNum 5 0 s
      two = findNum 6 one s
  print one
  print two
