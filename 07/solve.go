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
	Opcode{7, 3, "Jump instruction ptr to third param if first param less than second param"},
	Opcode{8, 3, "Jump instruction ptr to third param if first param equals second param"},
	// ...
	Opcode{99, 0, "Exit"},
}

// write to input chan, read from output chan, closes output on exit
// modifies arr
// I/O is blocking
func exec_prog(arr []int, input, output myReadWriter) {
	ii := 0
	inputOpen := true
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
			if inputOpen {
				arr[arr[ii+1]], inputOpen = input.ReadInt()
			} else {
				fmt.Println("Input closed")
				arr[arr[ii+1]] = 0
			}
		case 4:
			output.WriteInt(args[0])
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
	output.Close()
}

func exec_perm(base []int, perm []int) int {
	sig := 0
	for _, num := range perm {
		a := copy_arr(base)
		var input, output myChan
		input = make(chan int)
		output = make(chan int)

		go exec_prog(a, input, output)
		input <- num
		input <- sig

		sig = (<-output)
	}
	return sig
}

func exec_perm2(base []int, perm []int) int {
	var tapes [][]int
	ins := [](myChan){make(chan int, 10), make(chan int, 10), make(chan int, 10), make(chan int, 10), make(chan int, 10)}
	outs := [](myChan){make(chan int), make(chan int), make(chan int), make(chan int), make(chan int)}

	for ii, num := range perm {
		tapes = append(tapes, copy_arr(base))
		go exec_prog(tapes[ii], ins[ii], outs[ii])
		ins[ii] <- num
	}

	sig_input := []int{0, 0, 0, 0, 0}
	closed := false
	sig := 0

	for !closed {
		sig = sig_input[0]
		for ii := range perm {
			lopen := true
			ins[ii] <- sig_input[ii]
			sig_input[(ii+1)%5], lopen = <-outs[ii]
			if !lopen {
				closed = true
			}
		}
	}
	return sig
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

	// part one
	answer := 0
	for _, perm := range permutations([]int{4, 3, 2, 1, 0}) {
		ret := exec_perm(a, perm)
		if ret > answer {
			answer = ret
		}
	}
	fmt.Println("Part one:", answer)

	// part two
	answer2 := 0
	for _, perm := range permutations([]int{9, 8, 7, 6, 5}) {
		ret := exec_perm2(a, perm)
		if ret > answer2 {
			answer2 = ret
		}
	}
	fmt.Println("Part two:", answer2)
}
