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

func update_drone_pos(pos [2]int, direction int) [2]int {
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
	for _, n := range []int{1, 2, 3, 4} {
		newPos := update_drone_pos(dronePos, n)
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
		} else {
			visited[newPos[1]][newPos[0]] = 100
			path := find_path(newPos, dest, space, visited)
			if path != nil {
				fmt.Println(path, newPos, space[newPos[1]][newPos[0]])
				return append(path, n)
			}
		}
	}
	fmt.Println("BACKTRACK")
	return nil
}

// nav_to_node(dronePos, tmpNode, space, inChan, outChan)
func nav_to_node(dronePos, dest [2]int, space [][]byte, inChan, outChan chan int) int {
	if dronePos[0] == dest[0] && dronePos[1] == dest[1] {
		return -1
	}
	visited := make([][]byte, len(space))
	for n := range visited {
		visited[n] = make([]byte, len(space[n]))
	}
	path := find_path(dronePos, dest, space, visited)
	fmt.Println("DEST", dest, "POS", dronePos, path)
	output := -1
	test := [2]int{dronePos[0], dronePos[1]}
	for ii := len(path) - 1; ii >= 0; ii-- {
		if output == 0 {
			panic("WALL HIT")
		}
		test = update_drone_pos(test, path[ii])
		fmt.Println(space[test[1]][test[0]])
		fmt.Println(test)
		input := path[ii]
		inChan <- input
		output = <-outChan
		fmt.Println("IN:", input, "OUTPUT:", output)
	}
	return output
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

	qu := stack.New()
	qu.Push(dronePos)
	for !qu.Empty() {
		node := qu.Pop()
		tmpNode := node.([2]int)
		// check if node visited
		if space[tmpNode[1]][tmpNode[0]] != 0 {
			continue
		}
		// navigate to node
		output := nav_to_node(dronePos, tmpNode, space, inChan, outChan)
		// mark value, etc
		changed := false
		switch output {
		case -1:
			space[tmpNode[1]][tmpNode[0]] = VISITED
			changed = true
		case 0:
			fmt.Println("WALL", tmpNode)
			space[tmpNode[1]][tmpNode[0]] = WALL
			changed = false
		case 1:
			fmt.Println("VISITED", tmpNode)
			space[tmpNode[1]][tmpNode[0]] = VISITED
			dronePos = tmpNode
			changed = true
		case 2:
			fmt.Println("OXYGEN", tmpNode)
			space[tmpNode[1]][tmpNode[0]] = OXYGEN
			dronePos = tmpNode
			changed = true
		default:
			fmt.Println("AAAAAAAAAAAAAAAAAAAAAAA", output)
			panic("BAD")
		}

		// expand node
		if changed {
			for _, n := range []int{1, 2, 3, 4} {
				newNode := update_drone_pos(dronePos, n)
				if newNode[0] < 0 || newNode[1] < 0 || newNode[0] > spaceSize || newNode[1] > spaceSize || space[newNode[1]][newNode[0]] != 0 {
					continue
				}
				qu.Push(newNode)
			}
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
				default:
					c = '?'
				}
				fmt.Printf("%c", c)
			}
			fmt.Println()
		}
	}
}
