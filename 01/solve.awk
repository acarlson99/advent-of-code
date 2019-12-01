#!/usr/bin/awk

func f(n) {
	return int(n / 3) - 2
}

func findReal(n) {
	extra = f(n)
	if (extra <= 0) {
		return (n)
	} else {
		return (n + findReal(extra))
	}
}

{
	c += f($1)
	c2 += findReal(f($1))
}

END {
	print c
	print c2
}
