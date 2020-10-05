package main

// https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func copy_arr(base []int) []int {
	var a []int
	for _, num := range base {
		a = append(a, num)
	}
	return a
}

func getArg(arr []int, mode, ii int) int {
	if mode == 0 {
		return arr[arr[ii]]
	} else if mode == 1 {
		return arr[ii]
	}
	panic("Bad mode")
}

func parseOp(op int) ([]int, int) {
	modes := []int{}
	modes = append(modes, (op/100)%10)
	modes = append(modes, (op/1000)%10)
	modes = append(modes, (op/10000)%10)
	opcode := op % 100
	return modes, opcode
}
