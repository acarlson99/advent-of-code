import AOC
import Data.Char
import Data.List

partOne = sum . map (fromEnum . (== 7) . length . filter (not . isPrefixOf "cid"))

chkHeight [a, b, c, 'c', 'm'] = v' >= 150 && v' <= 193
  where
    v' = read [a, b, c]
chkHeight [a, b, 'i', 'n'] = v' >= 56 && v' <= 76
  where
    v' = read [a, b]
chkHeight _ = False

validate :: (String, String) -> Bool
validate ("byr", v) = v' >= 1920 && v' <= 2002
  where
    v' = read v
validate ("iyr", v) = v' >= 2010 && v' <= 2020
  where
    v' = read v
validate ("eyr", v) = v' >= 2020 && v' <= 2030
  where
    v' = read v
validate ("hgt", v) = chkHeight v
validate ("hcl", '#' : xs) = foldr (\a acc -> acc && (isDigit a || (a >= 'a' && a <= 'f'))) (length xs == 6) xs
validate ("ecl", v) = v `elem` ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]
validate ("pid", v) = all isDigit v && length v == 9
validate _ = False

transform = (\(x : xs) -> (x, mconcat xs)) . splitWhen (== ':')

countValid = foldr ((+) . fromEnum . validate) 0

isValid :: [String] -> Bool
isValid inp = countValid dta == 7
  where
    dta = map transform inp

partTwo :: [[String]] -> Int
partTwo = foldr ((+) . fromEnum . isValid) 0

formatInput = map sort . splitWhen (== "") . splitWhen (\a -> (a == ' ') || (a == '\n'))

main =
  do
    inp <- formatInput <$> readFile "input.txt"
    print $ partOne inp
    print $ partTwo inp
