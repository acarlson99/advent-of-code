# sed 's/[0-9]\{150\}/\0,/g' < input.txt | tr ',' '\n' | grep -v '^$' | awk -f solve.awk

BEGIN {
	ii = 0
}

{
	a[ii] = $1
	ii++
}

func count(arr, num) {
	split(arr, local_arr, "")
	cnt = 0
	for (ii in local_arr) {
		if (local_arr[ii] == num)
			cnt++
	}
	return cnt
}

func get_at_idx(arr, ii) {
	for (jj in arr) {
		split(arr[jj], local_arr, "")
		switch (local_arr[ii]) {
		case 0:
			return "░"
		case 1:
			return "█"
		}
	}
	return "X"
}

END {
	min = 1000000
	for (arr in a) {
		lmin = count(a[arr], 0)
		if (lmin < min) {
			min = lmin
			minarr = a[arr]
		}
	}
	print "Part one:", count(minarr, 2) * count(minarr, 1)

	for (ii = 0; ii <= length(a[0]); ii++) {
		msgarr[ii] = get_at_idx(a, ii)
	}

	print "Part two:"
	ORS = ""
	for (height = 0; height < 6; height++) {
		for (width = 1; width < 26; width++) {
			print msgarr[height * 25 + width]
		}
		print "\n"
	}
}
