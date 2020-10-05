package main

import (
	"bufio"
	"fmt"
	"os"

	intcode "../libintcode"
)

type io struct {
	in  chan int
	out chan int
}

type packet struct {
	dest int
	x    int
	y    int
}

var printed bool

func handle(chans []io, src, dest int) {
	x := <-chans[src].out
	y := <-chans[src].out
	if dest == 255 {
		if !printed {
			fmt.Println("Part one:", y)
			printed = true
		}
		nat.x = x
		nat.y = y
	} else {
		chans[dest].in <- x
		chans[dest].in <- y
	}
	idle = 0
}

var nat packet
var idle int

func main() {
	reader := bufio.NewReader(os.Stdin)
	program := intcode.Read_program(reader)

	idle = 0
	lastY := -1

	chans := []io{}
	for ii := 0; ii < 50; ii++ {
		inCh := make(chan int, 20)
		outCh := make(chan int, 20)
		// chans = append(chans, [2]chan int{inCh, outCh})
		tmp := io{inCh, outCh}
		chans = append(chans, tmp)
		go intcode.Exec_prog(intcode.Copy_arr(program), intcode.INTCChan(tmp.in), intcode.INTCChan(tmp.out))
		tmp.in <- ii
	}
	// intcode.Exec_prog(intcode.Copy_arr(program), intcode.INTCStdin{}, intcode.INTCStdout{})
	for {
		if idle >= 500 {
			chans[0].in <- nat.x
			chans[0].in <- nat.y
			if nat.y == lastY {
				fmt.Println("Part two:", lastY)
				goto end
			} else {
				lastY = nat.y
			}
			nat.x = 0
			nat.y = 0
			idle = 0
		}
		select {
		case dest := <-chans[0].out:
			handle(chans, 0, dest)
		case dest := <-chans[1].out:
			handle(chans, 1, dest)
		case dest := <-chans[2].out:
			handle(chans, 2, dest)
		case dest := <-chans[3].out:
			handle(chans, 3, dest)
		case dest := <-chans[4].out:
			handle(chans, 4, dest)
		case dest := <-chans[5].out:
			handle(chans, 5, dest)
		case dest := <-chans[6].out:
			handle(chans, 6, dest)
		case dest := <-chans[7].out:
			handle(chans, 7, dest)
		case dest := <-chans[8].out:
			handle(chans, 8, dest)
		case dest := <-chans[9].out:
			handle(chans, 9, dest)
		case dest := <-chans[10].out:
			handle(chans, 10, dest)
		case dest := <-chans[11].out:
			handle(chans, 11, dest)
		case dest := <-chans[12].out:
			handle(chans, 12, dest)
		case dest := <-chans[13].out:
			handle(chans, 13, dest)
		case dest := <-chans[14].out:
			handle(chans, 14, dest)
		case dest := <-chans[15].out:
			handle(chans, 15, dest)
		case dest := <-chans[16].out:
			handle(chans, 16, dest)
		case dest := <-chans[17].out:
			handle(chans, 17, dest)
		case dest := <-chans[18].out:
			handle(chans, 18, dest)
		case dest := <-chans[19].out:
			handle(chans, 19, dest)
		case dest := <-chans[20].out:
			handle(chans, 20, dest)
		case dest := <-chans[21].out:
			handle(chans, 21, dest)
		case dest := <-chans[22].out:
			handle(chans, 22, dest)
		case dest := <-chans[23].out:
			handle(chans, 23, dest)
		case dest := <-chans[24].out:
			handle(chans, 24, dest)
		case dest := <-chans[25].out:
			handle(chans, 25, dest)
		case dest := <-chans[26].out:
			handle(chans, 26, dest)
		case dest := <-chans[27].out:
			handle(chans, 27, dest)
		case dest := <-chans[28].out:
			handle(chans, 28, dest)
		case dest := <-chans[29].out:
			handle(chans, 29, dest)
		case dest := <-chans[30].out:
			handle(chans, 30, dest)
		case dest := <-chans[31].out:
			handle(chans, 31, dest)
		case dest := <-chans[32].out:
			handle(chans, 32, dest)
		case dest := <-chans[33].out:
			handle(chans, 33, dest)
		case dest := <-chans[34].out:
			handle(chans, 34, dest)
		case dest := <-chans[35].out:
			handle(chans, 35, dest)
		case dest := <-chans[36].out:
			handle(chans, 36, dest)
		case dest := <-chans[37].out:
			handle(chans, 37, dest)
		case dest := <-chans[38].out:
			handle(chans, 38, dest)
		case dest := <-chans[39].out:
			handle(chans, 39, dest)
		case dest := <-chans[40].out:
			handle(chans, 40, dest)
		case dest := <-chans[41].out:
			handle(chans, 41, dest)
		case dest := <-chans[42].out:
			handle(chans, 42, dest)
		case dest := <-chans[43].out:
			handle(chans, 43, dest)
		case dest := <-chans[44].out:
			handle(chans, 44, dest)
		case dest := <-chans[45].out:
			handle(chans, 45, dest)
		case dest := <-chans[46].out:
			handle(chans, 46, dest)
		case dest := <-chans[47].out:
			handle(chans, 47, dest)
		case dest := <-chans[48].out:
			handle(chans, 48, dest)
		case dest := <-chans[49].out:
			handle(chans, 49, dest)
		case chans[0].in <- -1:
			idle++
		case chans[1].in <- -1:
			idle++
		case chans[2].in <- -1:
			idle++
		case chans[3].in <- -1:
			idle++
		case chans[4].in <- -1:
			idle++
		case chans[5].in <- -1:
			idle++
		case chans[6].in <- -1:
			idle++
		case chans[7].in <- -1:
			idle++
		case chans[8].in <- -1:
			idle++
		case chans[9].in <- -1:
			idle++
		case chans[10].in <- -1:
			idle++
		case chans[11].in <- -1:
			idle++
		case chans[12].in <- -1:
			idle++
		case chans[13].in <- -1:
			idle++
		case chans[14].in <- -1:
			idle++
		case chans[15].in <- -1:
			idle++
		case chans[16].in <- -1:
			idle++
		case chans[17].in <- -1:
			idle++
		case chans[18].in <- -1:
			idle++
		case chans[19].in <- -1:
			idle++
		case chans[20].in <- -1:
			idle++
		case chans[21].in <- -1:
			idle++
		case chans[22].in <- -1:
			idle++
		case chans[23].in <- -1:
			idle++
		case chans[24].in <- -1:
			idle++
		case chans[25].in <- -1:
			idle++
		case chans[26].in <- -1:
			idle++
		case chans[27].in <- -1:
			idle++
		case chans[28].in <- -1:
			idle++
		case chans[29].in <- -1:
			idle++
		case chans[30].in <- -1:
			idle++
		case chans[31].in <- -1:
			idle++
		case chans[32].in <- -1:
			idle++
		case chans[33].in <- -1:
			idle++
		case chans[34].in <- -1:
			idle++
		case chans[35].in <- -1:
			idle++
		case chans[36].in <- -1:
			idle++
		case chans[37].in <- -1:
			idle++
		case chans[38].in <- -1:
			idle++
		case chans[39].in <- -1:
			idle++
		case chans[40].in <- -1:
			idle++
		case chans[41].in <- -1:
			idle++
		case chans[42].in <- -1:
			idle++
		case chans[43].in <- -1:
			idle++
		case chans[44].in <- -1:
			idle++
		case chans[45].in <- -1:
			idle++
		case chans[46].in <- -1:
			idle++
		case chans[47].in <- -1:
			idle++
		case chans[48].in <- -1:
			idle++
		case chans[49].in <- -1:
			idle++
		}
	}
end:
}
