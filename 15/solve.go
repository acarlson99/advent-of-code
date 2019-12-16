package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	gc "github.com/rthornton128/goncurses"
	"gopkg.in/karalabe/cookiejar.v2/collections/queue"
	"gopkg.in/karalabe/cookiejar.v2/collections/stack"
)

const (
	WALL    = 1
	VISITED = 2
	OXYGEN  = 3
)

func move_drone(pos [2]int, direction int) [2]int {
	x := pos[0]
	y := pos[1]
	switch direction {
	case 1:
		y -= 1
	case 2:
		y += 1
	case 3:
		x -= 1
	case 4:
		x += 1
	default:
		panic(fmt.Sprintf("Invalid input: %d", direction))
	}
	return [2]int{x, y}
}

func find_path(dronePos, dest [2]int, space, visited [][]byte) []int {
	if dronePos[0] == dest[0] && dronePos[1] == dest[1] {
		return []int{}
	}
	visited[dronePos[1]][dronePos[0]] = 100
	for _, n := range []int{1, 2, 3, 4} {
		newPos := move_drone(dronePos, n)
		if newPos[0] == dest[0] && newPos[1] == dest[1] {
			return []int{n}
		}
		if newPos[0] < 0 || newPos[1] < 0 || newPos[0] > len(space) ||
			newPos[1] > len(space) || visited[newPos[1]][newPos[0]] != 0 ||
			space[newPos[1]][newPos[0]] == WALL ||
			space[newPos[1]][newPos[0]] != VISITED ||
			space[newPos[1]][newPos[0]] == 0 {
			continue
		}
		visited[newPos[1]][newPos[0]] = 100
		path := find_path(newPos, dest, space, visited)
		if path != nil {
			return append(path, n)
		}
	}
	return nil
}

// nav_to_node(dronePos, tmpNode, space, inChan, outChan)
func nav_to_node(dronePos, dest [2]int, space [][]byte, inChan, outChan chan int) (int, [2]int, [2]int) {
	if dronePos[0] == dest[0] && dronePos[1] == dest[1] {
		return -1, dronePos, dronePos
	}
	visited := make([][]byte, len(space))
	for n := range visited {
		visited[n] = make([]byte, len(space[n]))
	}
	path := find_path(dronePos, dest, space, visited)
	output := -1
	tmp := [2]int{dronePos[0], dronePos[1]}
	// newDrone := [2]int{-1, -1}
	for ii := len(path) - 1; ii >= 0; ii-- {
		if output == 0 {
			panic("WALL HIT")
		}
		input := path[ii]
		inChan <- input
		output = <-outChan
		draw_map(space, tmp)
		if output == 0 && ii == 0 {
			return output, tmp, move_drone(tmp, path[ii])
		}
		tmp = move_drone(tmp, path[ii])
	}
	return output, tmp, tmp
}

func draw_map(space [][]byte, bot [2]int) {
	if vis != true {
		return
	}
	for n := range space {
		for x, c := range space[n] {
			var color int16
			switch c {
			case VISITED:
				c = '.'
				color = 1
			case WALL:
				c = '#'
				color = 2
			case OXYGEN:
				c = 'O'
				color = 3
			case 'X':
				c = 'X'
				color = 4
			default:
				c = ' '
			}
			stdscr.ColorOn(color)
			stdscr.MovePrint(n, x, string(c))
			stdscr.ColorOff(color)
		}
	}
	if bot[0] != -1 {
		stdscr.ColorOn(3)
		stdscr.MovePrint(bot[1], bot[0], "Y")
		stdscr.ColorOff(3)
	}
	stdscr.Refresh()
	time.Sleep(10000000)
}

var stdscr *gc.Window
var vis bool

func main() {
	flag.Parse()
	if len(flag.Args()) > 0 {
		var err error
		stdscr, err = gc.Init()
		if err != nil {
			log.Fatal("init:", err)
		}

		if !gc.HasColors() {
			log.Fatal("Example requires a colour capable terminal")
		}
		if err := gc.StartColor(); err != nil {
			log.Fatal(err)
		}

		gc.InitPair(1, gc.C_RED, gc.C_RED)
		gc.InitPair(2, gc.C_BLACK, gc.C_CYAN)
		gc.InitPair(3, gc.C_RED, gc.C_GREEN)
		gc.InitPair(4, gc.C_YELLOW, gc.C_BLUE)
		gc.Cursor(0)
		stdscr.Keypad(true)
		vis = true
	}

	reader := bufio.NewReader(os.Stdin)
	program := read_program(reader)
	inChan := make(chan int)
	outChan := make(chan int)
	go exec_prog(program, myChan(inChan), myChan(outChan))

	spaceSize := 50
	space := make([][]byte, spaceSize)
	for n := range space {
		space[n] = make([]byte, spaceSize)
	}
	// x,y
	dronePos := [2]int{spaceSize/2 + 1, spaceSize/2 + 1}
	startPos := dronePos
	oxygenPos := dronePos
	// space[dronePos[1]][dronePos[0]] = VISITED

	stk := stack.New()
	stk.Push(dronePos)
	for !stk.Empty() {
		tmpNode := stk.Pop()
		node := tmpNode.([2]int)
		// check if node visited
		if space[node[1]][node[0]] != 0 {
			continue
		}
		// navigate to node
		output, tmpPos, markPos := nav_to_node(dronePos, node, space, inChan, outChan)
		dronePos = tmpPos
		// mark value, etc
		changed := false
		if space[markPos[1]][markPos[0]] != 0 {
			panic("REVISITING")
		}
		switch output {
		case -1:
			space[markPos[1]][markPos[0]] = VISITED
			changed = true
		case 0:
			space[markPos[1]][markPos[0]] = WALL
			changed = false
		case 1:
			space[markPos[1]][markPos[0]] = VISITED
			changed = true
		case 2:
			space[markPos[1]][markPos[0]] = OXYGEN
			oxygenPos = markPos
			changed = true
		default:
			panic("BAD")
		}
		draw_map(space, dronePos)
		// expand node
		if changed {
			for _, n := range []int{1, 2, 3, 4} {
				newNode := move_drone(dronePos, n)
				if newNode[0] < 0 || newNode[1] < 0 || newNode[0] > spaceSize || newNode[1] > spaceSize || space[newNode[1]][newNode[0]] != 0 {
					continue
				}
				stk.Push(newNode)
			}
		}
	}

	// see search.go
	endNode := find_optimal_path(startPos, oxygenPos, copy_byte_arr_arr(space))

	space[oxygenPos[1]][oxygenPos[0]] = VISITED
	qu := queue.New()
	qu.Push(Node{oxygenPos, []int{}})
	// see search.go
	minutes := bfs_fill(qu, space)
	if vis {
		gc.End()
	}
	fmt.Println("Part one:", len(endNode.Path))
	fmt.Println("Part two:", minutes)
}
