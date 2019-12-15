package main

import (
	"bufio"
	"fmt"
	"os"

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
	// fmt.Println("NAV FROM", dronePos, "TO", dest)
	if dronePos[0] == dest[0] && dronePos[1] == dest[1] {
		return []int{}
	}
	visited[dronePos[1]][dronePos[0]] = 100
	for _, n := range []int{1, 2, 3, 4} {
		newPos := move_drone(dronePos, n)
		if newPos[0] == dest[0] && newPos[1] == dest[1] {
			// fmt.Println("FOUND", dronePos, dest)
			return []int{n}
		}
		// fmt.Println("SPACE", newPos, space[newPos[1]][newPos[0]])
		if newPos[0] < 0 || newPos[1] < 0 || newPos[0] > len(space) ||
			newPos[1] > len(space) || visited[newPos[1]][newPos[0]] != 0 ||
			space[newPos[1]][newPos[0]] == WALL ||
			space[newPos[1]][newPos[0]] != VISITED ||
			space[newPos[1]][newPos[0]] == 0 {
			continue
		}
		// fmt.Println("SPACE", space[newPos[1]][newPos[0]])
		visited[newPos[1]][newPos[0]] = 100
		path := find_path(newPos, dest, space, visited)
		if path != nil {
			// fmt.Println("PATH NOT NIL")
			// fmt.Println(path, newPos, space[newPos[1]][newPos[0]])
			return append(path, n)
		}
	}
	return nil
}

// nav_to_node(dronePos, tmpNode, space, inChan, outChan)
func nav_to_node(dronePos, dest [2]int, space [][]byte, inChan, outChan chan int) (int, [2]int, [2]int) {
	if dronePos[0] == dest[0] && dronePos[1] == dest[1] {
		// fmt.Println("No move")
		return -1, dronePos, dronePos
	}
	visited := make([][]byte, len(space))
	for n := range visited {
		visited[n] = make([]byte, len(space[n]))
	}
	path := find_path(dronePos, dest, space, visited)
	// fmt.Println("LEN", len(path))
	// fmt.Println("DEST", dest, "POS", dronePos, path)
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
		if output == 0 && ii == 0 {
			// fmt.Println("RETURNING", tmp, move_drone(tmp, path[ii]))
			return output, tmp, move_drone(tmp, path[ii])
		}
		tmp = move_drone(tmp, path[ii])
		// fmt.Println(space[tmp[1]][tmp[0]])
		// fmt.Println(tmp)
		// fmt.Println("IN:", input, "OUTPUT:", output)
	}
	// fmt.Println("A")
	return output, tmp, tmp
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	program := read_program(reader)
	inChan := make(chan int)
	outChan := make(chan int)
	go exec_prog(program, myChan(inChan), myChan(outChan))

	spaceSize := 64
	space := make([][]byte, spaceSize)
	for n := range space {
		space[n] = make([]byte, spaceSize)
	}
	// x,y
	dronePos := [2]int{spaceSize / 2, spaceSize / 2}
	// space[dronePos[1]][dronePos[0]] = VISITED

	qu := stack.New()
	qu.Push(dronePos)
	for !qu.Empty() {
		tmpNode := qu.Pop()
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
		// fmt.Println(dronePos, markPos)
		if space[markPos[1]][markPos[0]] != 0 {
			panic("REVISITING")
		}
		switch output {
		case -1:
			space[markPos[1]][markPos[0]] = VISITED
			changed = true
		case 0:
			// fmt.Println("WALL", markPos)
			space[markPos[1]][markPos[0]] = WALL
			changed = false
		case 1:
			// fmt.Println("VISITED", markPos)
			space[markPos[1]][markPos[0]] = VISITED
			// dronePos = markPos
			changed = true
		case 2:
			// fmt.Println("OXYGEN", markPos)
			space[markPos[1]][markPos[0]] = OXYGEN
			// dronePos = markPos
			changed = true
		default:
			// fmt.Println("AAAAAAAAAAAAAAAAAAAAAAA", output)
			panic("BAD")
		}

		// expand node
		if changed {
			for _, n := range []int{1, 2, 3, 4} {
				newNode := move_drone(dronePos, n)
				if newNode[0] < 0 || newNode[1] < 0 || newNode[0] > spaceSize || newNode[1] > spaceSize || space[newNode[1]][newNode[0]] != 0 {
					continue
				}
				qu.Push(newNode)
			}
		}
		// tmp := space[node[1]][node[0]]
		// space[node[1]][node[0]] = 'X'
	}
	for n := range space {
		for _, c := range space[n] {
			switch c {
			case VISITED:
				c = '.'
			case WALL:
				c = '#'
			case OXYGEN:
				c = 'O'
			case 'X':
				c = 'X'
			default:
				c = '?'
			}
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
	// space[node[1]][node[0]] = tmp
}
