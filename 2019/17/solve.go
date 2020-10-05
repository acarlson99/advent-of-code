package main

import (
	"bufio"
	"fmt"
	"os"

	intcode "../libintcode"
)

func check_xy(x, y int, arr [][]string) bool {
	if y < 0 || y >= len(arr) || x < 0 || x >= len(arr[y]) {
		return true
	}
	return arr[y][x] == "#"
}

func interpret_output(out []int) [][]string {
	whole := [][]string{}
	working := []string{}
	for _, n := range out {
		switch n {
		case 35:
			working = append(working, "#")
		case 46:
			working = append(working, ".")
		case 10:
			if len(working) > 0 {
				whole = append(whole, working)
				working = []string{}
			}
		default:
			working = append(working, string(byte(n)))
		}
	}
	return whole
}

func main() {
	// setup
	reader := bufio.NewReader(os.Stdin)
	program := intcode.Read_program(reader)

	a := []int{}
	arrOut := intcode.INTCArrWriter{&a}

	// Part one
	intcode.Exec_prog(intcode.Copy_arr(program), intcode.INTCChan(make(chan int)), arrOut)

	whole := interpret_output(a)

	total := 0
	for y, line := range whole {
		for x, _ := range line {
			if check_xy(x, y, whole) &&
				check_xy(x, y+1, whole) &&
				check_xy(x, y-1, whole) &&
				check_xy(x+1, y, whole) &&
				check_xy(x-1, y, whole) {
				total += x * y
			}
		}
	}

	for _, line := range whole {
		fmt.Println(line)
	}

	fmt.Println("Part one:", total)

	// Part two
	program[0] = 2
	inChan := make(chan int)
	outChan := make(chan int)

	go intcode.Exec_prog(intcode.Copy_arr(program), intcode.INTCChan(inChan), intcode.INTCChan(outChan))

	arrIdx := 0
	arrNIdx := 0
	inArrs := [][]int{
		[]int{'A', ',', 'B', ',', 'A', ',', 'C', ',', 'A', ',', 'B', ',', 'C', ',', 'B', ',', 'C', ',', 'A', '\n'},
		[]int{'L', ',', '1', '2', ',', 'R', ',', '4', ',', 'R', ',', '4', ',', 'L', ',', '6', '\n'},
		[]int{'L', ',', '1', '2', ',', 'R', ',', '4', ',', 'R', ',', '4', ',', 'R', ',', '1', '2', '\n'},
		[]int{'L', ',', '1', '0', ',', 'L', ',', '6', ',', 'R', ',', '4', '\n'},
		[]int{'n', '\n'},
		[]int{'0'},
	}

	output := 0
	for {
		select {
		case output2 := <-outChan:
			if output2 == 0 {
				goto end
			} else {
				output = output2
			}
		case inChan <- inArrs[arrIdx][arrNIdx]:
			if inArrs[arrIdx][arrNIdx] == '\n' {
				arrNIdx = 0
				arrIdx++
			} else {
				arrNIdx++
			}
		default:
			break
		}
	}
end:
	fmt.Println("Part two:", output)
}
