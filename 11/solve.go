package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// return num of pannels it paints
func repair_robot(program []int, startingColor int) (int, [][]int) {
	hullWidth := 5000
	hullHeight := 5000

	hull := make([][]int, hullHeight)
	for ii := range hull {
		hull[ii] = make([]int, hullWidth)
	}

	count := make([][]int, hullHeight)
	for ii := range count {
		count[ii] = make([]int, hullWidth)
	}

	botX := hullWidth / 2
	botY := hullHeight / 2

	// 0,1,2,3 = up,right,down,left
	botFacing := 0
	dirs := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	input := myChan(make(chan int))
	output := myChan(make(chan int))
	go exec_prog(program, input, output)

	hull[botY][botX] = startingColor

	for {
		go func() {
			input <- hull[botY][botX]
		}()

		// 0 black 1 white
		color, open := <-output
		// hangs here for some reason
		hull[botY][botX] = color
		count[botY][botX]++
		if !open {
			break
		}

		// 0 left 90 deg, 1 right 90 deg
		direction, open := <-output
		if !open {
			break
		}
		if direction == 0 {
			botFacing--
		} else {
			botFacing++
		}
		botFacing = (botFacing + 4) % 4

		d := dirs[botFacing]
		botX += d[0]
		botY += d[1]
	}
	total := 0
	for _, line := range count {
		for _, v := range line {
			if v != 0 {
				total++
			}
		}
	}
	return total, hull
}

func main() {
	flag.Parse()
	args := flag.Args()
	inFile, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(inFile)
	program := read_program(reader)
	total, _ := repair_robot(copy_arr(program), 0)
	fmt.Println("Part one:", total)

	fmt.Println("Part two:")
	total, hull := repair_robot(copy_arr(program), 1)

	// trim output to only what I need
	lines := [][]int{}
	m := 100000
	for _, line := range hull {
		for ii, v := range line {
			if v != 0 {
				lines = append(lines, line)
				if ii < m {
					m = ii
				}
				goto eol
			}
		}
	eol:
	}

	for _, line := range lines {
		for ii := m; ii < m+50; ii++ {
			c := ' '
			if line[ii] == 0 {
				c = '░'
			} else {
				c = '█'
			}
			fmt.Printf("%c", c)
		}
		fmt.Printf("\n")
	}
}
