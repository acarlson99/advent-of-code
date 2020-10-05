package main

func sumArr(x ...int) int {
	total := 0
	for _, n := range x {
		total += abs(n)
	}
	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(x ...int) int {
	if len(x) == 0 {
		return 0
	} else if len(x) == 2 {
		for x[1] != 0 {
			x[0], x[1] = x[1], x[0]%x[1]
		}
	} else if len(x) > 2 {
		return gcd(x[0], gcd(x[1:]...))
	}
	return abs(x[0])
}

func lcm(x ...int) int {
	if len(x) > 2 {
		return lcm(x[0], lcm(x[1:]...))
	} else if x[0] == 0 && x[1] == 0 {
		return 0
	}
	return abs(x[0]*x[1]) / gcd(x[0], x[1])
}
