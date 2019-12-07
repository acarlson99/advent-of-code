package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Opcode struct {
	code int
	args int
	desc string
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

var ops []Opcode = []Opcode{
	Opcode{0, 0, "NULL"},
	Opcode{1, 3, "Add first two params and store in third"},
	Opcode{2, 3, "Mult first two params and store in third"},
	Opcode{3, 1, "Read from stdin and store int result in arg 1"},
	Opcode{4, 1, "Print var at arg 1"},
	Opcode{5, 2, "Jump instruction ptr to second param if first param non-zero"},
	Opcode{6, 2, "Jump instruction ptr to second param if first param zero"},
	Opcode{7, 3, "Jump instruction ptr to third param if first param less than second param"},
	Opcode{8, 3, "Jump instruction ptr to third param if first param equals second param"},
	// ...
	Opcode{99, 0, "Exit"},
}

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

// return output, halted
func exec_prog(arr []int, inputs []int) (int, bool) {
	ii := 0
	inputnum := 0
	for {
		if arr[ii]%100 == 99 {
			goto end
		}

		ptrmod := false

		modes, op := parseOp(arr[ii])
		args := []int{0, 0, 0, 0}
		for jj := 0; jj < ops[op].args; jj++ {
			args[jj] = getArg(arr, modes[jj], ii+jj+1)
		}

		switch op {
		case 1:
			arr[arr[ii+3]] = args[0] + args[1]
		case 2:
			arr[arr[ii+3]] = args[0] * args[1]
		case 3:
			// reader := bufio.NewReader(os.Stdin)
			// text, _ := reader.ReadString('\n')
			// if len(text) < 1 {
			// 	fmt.Println("Correcting text to '0\\n'")
			// 	text = "0\n"
			// }
			// text = text[0 : len(text)-1]
			// arr[arr[ii+1]], _ = strconv.Atoi(text)
			arr[arr[ii+1]] = inputs[inputnum]
			inputnum++
		case 4:
			return args[0], false
			fmt.Println("OUTPUT:", args[0])
		case 5:
			if args[0] != 0 {
				ii = args[1]
				ptrmod = true
			}
		case 6:
			if args[0] == 0 {
				ii = args[1]
				ptrmod = true
			}
		case 7:
			if args[0] < args[1] {
				arr[arr[ii+3]] = 1
			} else {
				arr[arr[ii+3]] = 0
			}
		case 8:
			if args[0] == args[1] {
				arr[arr[ii+3]] = 1
			} else {
				arr[arr[ii+3]] = 0
			}
		case 99:
			goto end
		default:
			fmt.Println("YIKES unexpected opcode", arr[ii])
			os.Exit(1)
		}
		if !ptrmod {
			ii += ops[op].args + 1
		}
	}

end:
	fmt.Println("AAAAAAAAAAAAAAAAAAA")
	return arr[0], true
}

func exec_copy(base []int, inputs []int) (int, bool) {
	fmt.Println(inputs)
	var a []int
	for _, num := range base {
		a = append(a, num)
	}

	return exec_prog(a, inputs)
}

func exec_perm(base []int, perm []int) int {
	sig := 0
	for _, num := range perm {
		sig, _ = exec_copy(base, []int{num, sig})
	}
	return sig
}

func exec_perm2(base []int, perm []int) int {
	sig := 0
	halted := false
	esig := 0
	for ii := 0; ii < 10; ii++ {
		for _, num := range perm {
			sig, halted = exec_copy(base, []int{num, sig})
			fmt.Println(sig, halted, ii)
			if halted {
				goto end
			}
		}
		esig = sig
	}

end:
	return esig
}

func main() {
	// setup
	var a []int

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	ints := strings.Split(text, ",")

	intlen := len(ints) - 1
	lastnum := ints[intlen]
	lastnum = lastnum[0 : len(lastnum)-1]
	ints[intlen] = lastnum

	for _, num := range ints {
		n, err := strconv.Atoi(num)
		if err != nil {
			break
		}
		a = append(a, n)
	}

	// // part one
	// answer := 0
	// for _, perm := range permutations([]int{4, 3, 2, 1, 0}) {
	// 	ret := exec_perm(a, perm)
	// 	if ret > answer {
	// 		answer = ret
	// 	}
	// }
	// fmt.Println(answer)

	// part two
	fmt.Println(exec_perm2(a, []int{9, 8, 7, 6, 5}))
	// answer2 := 0
	// for _, perm := range permutations([]int{9, 8, 7, 6, 5}) {
	// 	ret := exec_perm2(a, perm)
	// 	if ret > answer2 {
	// 		answer2 = ret
	// 	}
	// }
	// fmt.Println(answer2)
}
