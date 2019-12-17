package main

import (
	"bufio"
	"fmt"
	"os"
)

func check_xy(x, y int, arr [][]string) bool {
	if y < 0 || y >= len(arr) || x < 0 || x >= len(arr[y]) {
		return true
	}
	return arr[y][x] == "#"
}

func main() {
	// setup
	reader := bufio.NewReader(os.Stdin)
	program := read_program(reader)

	a := []int{}
	arrOut := myArrWriter{&a}

	exec_prog(copy_arr(program), myChan(make(chan int)), arrOut)

	whole := [][]string{}
	working := []string{}
	for _, n := range a {
		switch n {
		case 35:
			working = append(working, "#")
		case 46:
			working = append(working, ".")
		case 10:
			whole = append(whole, working)
			working = []string{}
		}
	}

	total := 0
	for y, line := range whole {
		for x, _ := range line {
			if check_xy(x, y, whole) &&
				check_xy(x, y+1, whole) &&
				check_xy(x, y-1, whole) &&
				check_xy(x+1, y, whole) &&
				check_xy(x-1, y, whole) {
				whole[y][x] = "0"
				total += x * y
			}
		}
	}

	fmt.Println("Part one:", total)
}
