package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	intcode "../libintcode"
)

func pw2(n float64) float64 {
	return n * n
}

func main() {
	var reader io.Reader
	reader = bufio.NewReader(os.Stdin)

	beamMap := [50][50]int{}

	count := 0
	program := intcode.Read_program(reader)
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			outArr := []int{}
			out := intcode.INTCArrWriter{&outArr}
			intcode.Exec_prog(intcode.Copy_arr(program), &intcode.INTCArr{[]int{x, y}, 0}, out)
			beamMap[y][x] = outArr[0]
			if outArr[0] == 1 {
				count++
			}
		}
	}
	for _, line := range beamMap {
		for _, c := range line {
			switch c {
			case 1:
				fmt.Printf("O")
			case 0:
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}

	fmt.Println("Part one:", count)
}
