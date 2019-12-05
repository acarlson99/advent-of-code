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
}

func getArg(arr []int, mode, ii int) int {
	if mode == 0 {
		return arr[arr[ii]]
	} else if mode == 1 {
		return arr[ii]
	}
	panic("Bad mode")
}

func getModes(op int) ([]int, int) {
	modes := []int{}
	modes = append(modes, (op/100)%10)
	modes = append(modes, (op/1000)%10)
	modes = append(modes, (op/10000)%10)
	opcode := op % 100
	return modes, opcode
}

func exec_prog(arr []int) int {
	ii := 0
	for {
		if arr[ii] == 99 {
			goto end
		}

		modes, op := getModes(arr[ii])
		// fmt.Println(modes, op)
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
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			text = text[0 : len(text)-1]
			arr[arr[ii+1]], _ = strconv.Atoi(text)
		case 4:
			fmt.Println(args[0])
		case 99:
			goto end
		default:
			fmt.Println(arr[ii])
			fmt.Println("YIKES")
			os.Exit(1)
		}
		ii += ops[op].args + 1
	}
end:
	return arr[0]
}

func main() {
	var a []int

	reader := bufio.NewReader(os.Stdin)
	// reader := strings.NewReader("1,1,1,0,99\n")
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

	fmt.Println("Part one:", exec_prog(a))
}
