# ../intcode/intcode -e -f input.txt | awk -f solve.awk

BEGIN {
	total = 0
	lineNum = 0
}

{
	if (lineNum % 3 == 2 && $1 == "2") {
		total++
	}
	lineNum++
}

END {
	print "Part one: " total
}
