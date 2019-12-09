package main

func copy_arr(base []int) []int {
	var a []int
	for _, num := range base {
		a = append(a, num)
	}
	return a
}

func getArgs(arr, mode []int, ii, nargs, rel int) []int {
	ret := []int{}
	for jj := 0; jj < nargs; jj++ {
		n := 0
		switch mode[jj] {
		case 0:
			n = arr[ii+jj+1]
		case 1:
			n = ii + jj + 1
		case 2:
			n = arr[ii+jj+1] + rel
		default:
			panic("Invalid mode")
		}
		ret = append(ret, n)
	}
	return ret
}

func parseOp(op int) ([]int, int) {
	modes := []int{}
	pos := 100
	for ii := 0; ii < 3; ii++ {
		modes = append(modes, (op/pos)%10)
		pos *= 10
	}
	// modes = append(modes, (op/100)%10)
	// modes = append(modes, (op/1000)%10)
	// modes = append(modes, (op/10000)%10)
	opcode := op % 100
	return modes, opcode
}
