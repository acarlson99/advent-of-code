package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"

	"gopkg.in/karalabe/cookiejar.v2/collections/queue"
)

type State struct {
	x    int
	y    int
	keys [26]bool
	dist int
}

func (s *State) Hash() uint64 {
	str := fmt.Sprintf("%v,%v,%v", s.x, s.y, s.keys)
	h := fnv.New64()
	h.Write([]byte(str))
	return h.Sum64()
}

func (s *State) Copy() State {
	keys := [26]bool{}
	for ii, k := range s.keys {
		keys[ii] = k
	}
	newS := State{s.x, s.y, keys, s.dist}
	return newS
}

func (s *State) Expand() []State {
	r := []State{}

	for _, val := range [][]int{
		[]int{-1, 0},
		[]int{1, 0},
		[]int{0, -1},
		[]int{0, 1},
	} {
		newS := s.Copy()
		newS.x += val[0]
		newS.y += val[1]
		newS.dist++
		r = append(r, newS)
	}
	return r
}

func (s *State) Check(maze [][]byte) bool {
	c := maze[s.y][s.x]
	if c == '#' {
		return false
	} else if c >= 'A' && c <= 'Z' { // door
		c += 32
		return s.keys[c-'a']
	} else if c >= 'a' && c <= 'z' { // key
		s.keys[c-'a'] = true
		return true
	} else if c == '@' {
		return true
	}
	return true
}

type State2 struct {
	positions [][2]int
	keys      [26]bool
	dist      int
}

func (s *State2) Hash() uint64 {
	str := fmt.Sprintf("%v:%v", s.positions, s.keys)
	h := fnv.New64()
	h.Write([]byte(str))
	return h.Sum64()
}

func (s *State2) Copy() State2 {
	newPos := [][2]int{}
	var newKeys [26]bool

	for _, v := range s.positions {
		newArr := [2]int{v[0], v[1]}
		newPos = append(newPos, newArr)
	}
	for ii, v := range s.keys {
		newKeys[ii] = v
	}
	return State2{newPos, newKeys, s.dist}
}

func (s *State2) Expand() []State2 {
	r := []State2{}

	for _, val := range [][]int{
		[]int{-1, 0},
		[]int{1, 0},
		[]int{0, -1},
		[]int{0, 1},
	} {
		for ii := range s.positions {
			newS := s.Copy()
			newS.positions[ii][0] += val[0]
			newS.positions[ii][1] += val[1]
			newS.dist++
			r = append(r, newS)
		}
	}
	return r
}

func (s *State2) Check(idx int, maze [][]byte) bool {
	for _, pos := range s.positions {
		c := maze[pos[1]][pos[0]]
		if c == '#' {
			return false
		} else if c >= 'A' && c <= 'Z' { // door
			c += 32
			if !s.keys[c-'a'] {
				return false
			}
			// return s.keys[c-'a']
		} else if c >= 'a' && c <= 'z' { // key
			s.keys[c-'a'] = true
			// return true
		} else if c == '@' {
			// return true
		}
		// return true
	}
	return true
}

func partOne(startX, startY, targetKeys int, maze [][]byte) int {
	// seen states
	seen := make(map[uint64]int)
	qu := queue.New()

	initState := State{}
	initState.dist = 0
	initState.x = startX
	initState.y = startY

	// enqueue init state
	qu.Push(initState)
	for !qu.Empty() {
		tmp := qu.Pop()
		node := tmp.(State)
		numKeys := 0
		for _, b := range node.keys {
			if b {
				numKeys++
			}
		}
		if numKeys >= targetKeys {
			return node.dist
		}
		h := node.Hash()
		// check if seen
		if dist, ok := seen[h]; ok && dist <= node.dist {
			continue
		}
		seen[h] = node.dist

		for _, state := range node.Expand() {
			newH := state.Hash()
			if dist, ok := seen[newH]; ok && dist <= state.dist {
				continue
			}
			if state.Check(maze) {
				qu.Push(state)
			}
		}
	}

	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	// maze := []string{}
	maze := [][]byte{}
	for line, _, err := reader.ReadLine(); err == nil; line, _, err = reader.ReadLine() {
		buf := make([]byte, len(line))
		copy(buf, line)
		maze = append(maze, buf)
	}

	numKeys := 0
	for _, line := range maze {
		for _, c := range line {
			fmt.Printf("%c", c)
			if c >= 'a' && c <= 'z' {
				numKeys++
			}
		}
		fmt.Println()
	}

	startX := 0
	startY := 0
	for y := range maze {
		for x := range maze[y] {
			if maze[y][x] == '@' {
				startX = x
				startY = y
			}
		}
	}

	fmt.Println("Part one:", partOne(startX, startY, numKeys, maze))

	// for y := -1; y <= 1; y++ {
	// 	for x := -1; x <= 1; x++ {
	// 		nx := startX + x
	// 		ny := startY + y
	// 		if x != 0 && y != 0 {
	// 			maze[ny][nx] = '@'
	// 		} else {
	// 			maze[ny][nx] = '#'
	// 		}
	// 	}
	// }

	// numKeys = 0
	// for _, line := range maze {
	// 	for _, c := range line {
	// 		fmt.Printf("%c", c)
	// 		if c >= 'a' && c <= 'z' {
	// 			numKeys++
	// 		}
	// 	}
	// 	fmt.Println()
	// }
}
