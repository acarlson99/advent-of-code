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

var ops []Opcode = []Opcode{
	Opcode{0, 0, "NULL"},
	Opcode{1, 3, "Add first two params and store in third"},
	Opcode{2, 3, "Mult first two params and store in third"},
	Opcode{3, 1, "Read from stdin and store int result in arg 1"},
	Opcode{4, 1, "Print var at arg 1"},
	Opcode{5, 2, "Jump instruction ptr to second param if first param non-zero"},
	Opcode{6, 2, "Jump instruction ptr to second param if first param zero"},
	Opcode{7, 3, "Store 1 in third param if first param less than second param, else 0"},
	Opcode{8, 3, "Store 1 in third param if first param equals second param, else 0"},
	Opcode{9, 1, "Change relative base to value of param"},
	// ...
	Opcode{99, 0, "Exit"},
}

// write to input chan, read from output chan, closes output on exit
// modifies arr
// I/O is blocking
func exec_prog(arr []int, input myReader, output myWriter) {
	ii := 0
	inputOpen := true
	rel := 0
	for {
		if arr[ii]%100 == 99 {
			goto end
		}

		ptrmod := false

		modes, op := parseOp(arr[ii])
		args := getArgs(arr, modes, ii, ops[op].args, rel)
		// args := []int{0, 0, 0, 0}
		// for jj := 0; jj < ops[op].args; jj++ {
		// 	args[jj] = getArg(arr, modes[jj], ii+jj+1)
		// }

		// fmt.Println(ii, arr)
		switch op {
		case 1:
			arr[args[2]] = arr[args[0]] + arr[args[1]]
		case 2:
			arr[args[2]] = arr[args[0]] * arr[args[1]]
		case 3:
			if inputOpen {
				arr[args[0]], inputOpen = input.ReadInt()
			} else {
				fmt.Println("Input closed")
				arr[args[0]] = 0
			}
		case 4:
			output.WriteInt(arr[args[0]])
		case 5:
			if arr[args[0]] != 0 {
				ii = arr[args[1]]
				ptrmod = true
			}
		case 6:
			if arr[args[0]] == 0 {
				ii = arr[args[1]]
				ptrmod = true
			}
		case 7:
			if arr[args[0]] < arr[args[1]] {
				arr[args[2]] = 1
			} else {
				arr[args[2]] = 0
			}
		case 8:
			if arr[args[0]] == arr[args[1]] {
				arr[args[2]] = 1
			} else {
				arr[args[2]] = 0
			}
		case 9:
			rel += arr[args[0]]
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
	output.Close()
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

	for ii := 0; ii < 10000; ii++ {
		a = append(a, 0)
	}

	// part one
	exec_prog(copy_arr(a), myInt(1), myStdin{})

	// part two
	exec_prog(copy_arr(a), myInt(2), myStdin{})
}
