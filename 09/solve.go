package main

import (
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
func exec_prog(arr []int, input, output myReadWriter) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

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
			rel = arr[args[0]]
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

	// reader := bufio.NewReader(os.Stdin)
	// text, _ := reader.ReadString('\n')

	text := "3,225,1,225,6,6,1100,1,238,225,104,0,1002,114,46,224,1001,224,-736,224,4,224,1002,223,8,223,1001,224,3,224,1,223,224,223,1,166,195,224,1001,224,-137,224,4,224,102,8,223,223,101,5,224,224,1,223,224,223,1001,169,83,224,1001,224,-90,224,4,224,102,8,223,223,1001,224,2,224,1,224,223,223,101,44,117,224,101,-131,224,224,4,224,1002,223,8,223,101,5,224,224,1,224,223,223,1101,80,17,225,1101,56,51,225,1101,78,89,225,1102,48,16,225,1101,87,78,225,1102,34,33,224,101,-1122,224,224,4,224,1002,223,8,223,101,7,224,224,1,223,224,223,1101,66,53,224,101,-119,224,224,4,224,102,8,223,223,1001,224,5,224,1,223,224,223,1102,51,49,225,1101,7,15,225,2,110,106,224,1001,224,-4539,224,4,224,102,8,223,223,101,3,224,224,1,223,224,223,1102,88,78,225,102,78,101,224,101,-6240,224,224,4,224,1002,223,8,223,101,5,224,224,1,224,223,223,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,1107,226,677,224,102,2,223,223,1006,224,329,101,1,223,223,1108,226,677,224,1002,223,2,223,1005,224,344,101,1,223,223,8,226,677,224,102,2,223,223,1006,224,359,1001,223,1,223,1007,226,677,224,1002,223,2,223,1005,224,374,101,1,223,223,1008,677,677,224,1002,223,2,223,1005,224,389,1001,223,1,223,1108,677,226,224,1002,223,2,223,1006,224,404,1001,223,1,223,1007,226,226,224,1002,223,2,223,1005,224,419,1001,223,1,223,1107,677,226,224,1002,223,2,223,1006,224,434,101,1,223,223,108,677,677,224,1002,223,2,223,1005,224,449,1001,223,1,223,1107,677,677,224,102,2,223,223,1005,224,464,1001,223,1,223,108,226,226,224,1002,223,2,223,1006,224,479,1001,223,1,223,1008,226,226,224,102,2,223,223,1005,224,494,101,1,223,223,108,677,226,224,102,2,223,223,1005,224,509,1001,223,1,223,8,677,226,224,1002,223,2,223,1006,224,524,101,1,223,223,7,226,677,224,1002,223,2,223,1006,224,539,101,1,223,223,7,677,226,224,102,2,223,223,1006,224,554,1001,223,1,223,7,226,226,224,1002,223,2,223,1006,224,569,101,1,223,223,107,677,677,224,102,2,223,223,1006,224,584,101,1,223,223,1108,677,677,224,102,2,223,223,1006,224,599,1001,223,1,223,1008,677,226,224,1002,223,2,223,1005,224,614,1001,223,1,223,8,677,677,224,1002,223,2,223,1006,224,629,1001,223,1,223,107,226,677,224,1002,223,2,223,1006,224,644,101,1,223,223,1007,677,677,224,102,2,223,223,1006,224,659,101,1,223,223,107,226,226,224,1002,223,2,223,1006,224,674,1001,223,1,223,4,223,99,226\n"
	// text = "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99\n"

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

	// part one
	exec_prog(a, myStdin{}, myStdin{})
}
