BEGIN {
	cnt = 0
}

/byr/ {
	byr = 1
}

/iyr/ {
	iyr = 1
}

/eyr/ {
	eyr = 1
}

/hgt/ {
	hgt = 1
}

/hcl/ {
	hcl = 1
}

/ecl/ {
	ecl = 1
}

/pid/ {
	pid = 1
}

/cid/ {
	cid = 1
}

function chk() {
	return (byr == 1 && iyr == 1 && eyr == 1 && hgt == 1 && hcl == 1 && ecl == 1 && pid == 1)
}

/^$/ {

	if (chk()) {
		cnt += 1
	}

	byr = 0
	iyr = 0
	eyr = 0
	hgt = 0
	hcl = 0
	ecl = 0
	pid = 0
	cid = 0
}

END {
	if (chk()) {
		cnt += 1
	}
	print cnt
}
