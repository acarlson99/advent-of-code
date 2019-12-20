package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	intcode "../libintcode"
)

func check(x, y int, program []int) int {
	outArr := []int{}
	out := intcode.INTCArrWriter{&outArr}
	intcode.Exec_prog(intcode.Copy_arr(program), &intcode.INTCArr{[]int{x, y}, 0}, out)
	return outArr[0]
}

func main() {
	var reader io.Reader
	reader = bufio.NewReader(os.Stdin)

	beamMap := [50][50]int{}

	count := 0
	program := intcode.Read_program(reader)
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			beamMap[y][x] = check(x, y, program)
			if beamMap[y][x] == 1 {
				count++
			}
		}
	}

	// for _, line := range beamMap {
	// 	for _, c := range line {
	// 		switch c {
	// 		case 1:
	// 			fmt.Printf("O")
	// 		case 0:
	// 			fmt.Printf(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println("Part one:", count)

	x := 0
	y := 0
	for check(x+99, y, program) != 1 {
		y++
		for check(x, y+99, program) != 1 {
			x++
		}
	}
	fmt.Println("Part two:", x*10000+y)
}
