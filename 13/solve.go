package main

import (
	"bufio"
	"fmt"
	"os"

	intcode "../libintcode"
)

func selectVal(a, b int) (r int) {
	if a == b {
		return 0
	} else if a < b {
		return 1
	} else {
		return -1
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	program := intcode.Read_program(reader)

	program[0] = 2

	progIn := intcode.INTCChan(make(chan int))
	progOut := intcode.INTCChan(make(chan int))

	go intcode.Exec_prog(program, progIn, progOut)

	// x,y,val
	outs := []int{0, 0, 0}
	score := 0

	screenWH := 50
	screen := make([][]int, screenWH)
	for ii := range screen {
		screen[ii] = make([]int, screenWH)
	}

	paddleX := 0
	ballX := 0

	for {
		open := true
		select {
		case outs[0], open = <-progOut:
			outs[1], open = <-progOut
			outs[2], open = <-progOut
			if !open {
				goto end
			}
			if outs[0] == -1 && outs[1] == 0 {
				score = outs[2]
			} else {
				switch outs[2] {
				case 3:
					paddleX = outs[0]
					screen[outs[1]][outs[0]] = outs[2]
				case 4:
					ballX = outs[0]
					screen[outs[1]][outs[0]] = outs[2]
				default:
					screen[outs[1]][outs[0]] = outs[2]
				}
			}
		case progIn <- selectVal(paddleX, ballX):
		}

	}
end:
	fmt.Println("Part two:", score)
}
