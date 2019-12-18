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
	str := fmt.Sprintf("%v%v%v", s.x, s.y, s.keys)
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
		r = append(r, newS)
	}
	return r
}

func (s *State) Check(maze []string) bool {
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

func partOne(targetKeys int, maze []string) int {
	// seen states
	seen := make(map[uint64]int)
	qu := queue.New()

	initState := State{}
	initState.dist = 0
	for y := range maze {
		for x := range maze[y] {
			if maze[y][x] == '@' {
				initState.x = x
				initState.y = y
			}
		}
	}

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
		if _, ok := seen[h]; ok {
			continue
		}
		seen[h] = node.dist

		for _, state := range node.Expand() {
			state.dist++
			newH := state.Hash()
			if _, ok := seen[newH]; ok {
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
	maze := []string{}
	for line, _, err := reader.ReadLine(); err == nil; line, _, err = reader.ReadLine() {
		maze = append(maze, string(line))
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

	fmt.Println("Part one:", partOne(numKeys, maze))
}
