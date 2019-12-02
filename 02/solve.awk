BEGIN {
	FS = ","
}

{
	for (ii = 1; ii <= NF; ii++) {
		base[ii - 1] = $ii
	}
}

func eval_program(base, noun, verb) {
	for (ii in base) {
		arr[ii] = base[ii]
	}
	arr[1] = noun
	arr[2] = verb
	ii = 0
	done = 0
	while (done != 1) {
		switch (arr[ii]) {
			case 1:
				a = arr[arr[ii + 1]]
				b = arr[arr[ii + 2]]
				arr[arr[ii + 3]] = a + b
				break
			case 2:
				a = arr[arr[ii + 1]]
				b = arr[arr[ii + 2]]
				arr[arr[ii + 3]] = a * b
				break
			case 99:
				done = 1
				break
		}
		ii += 4
	}
	return arr[0]
}

END {
	print "Part one:", eval_program(base, 12, 2)

	for (n = 0; n < 100; n++) {
		for (v = 0; v < 100; v++) {
			if (eval_program(base, n, v) == 19690720) {
				print "Part two:", 100 * n + v
				exit 0
			}
		}
	}
}
