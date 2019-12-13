# ../intcode/intcode -f input.txt | awk -f solve.awk

BEGIN {
	lineNum = 0
}

{
	switch (lineNum % 3) {
	case 0:
		x = $1
		break
	case 1:
		y = $1
		break
	case 2:
		a[x "," y] = $1
		break
	}
	lineNum++
}

END {
	total = 0
	for (ii in a) {
		if (a[ii] == "2") {
			total++
		}
	}
	print total
}
