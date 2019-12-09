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

func exec_prog(arr []int, input int) int {
	ii := 0
	for {
		if arr[ii] == 99 {
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
			arr[arr[ii+1]] = input
		case 4:
			fmt.Println(args[0])
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
	return arr[0]
}

func copy_arr(base []int) []int {
	var a []int
	for _, num := range base {
		a = append(a, num)
	}
	return a
}

func main() {
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

	exec_prog(copy_arr(a), 1)
	exec_prog(copy_arr(a), 5)
}
