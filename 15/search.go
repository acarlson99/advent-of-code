package main

import (
	"gopkg.in/karalabe/cookiejar.v2/collections/queue"
)

type Node struct {
	Pos  [2]int
	Path []int
}

func find_optimal_path(startPos, oxygenPos [2]int, space [][]byte) Node {
	qu := queue.New()
	qu.Push(Node{startPos, []int{}})
	for !qu.Empty() {
		tmpNode := qu.Pop()
		node := tmpNode.(Node)
		pos := node.Pos
		val := space[pos[1]][pos[0]]
		if pos[1] == oxygenPos[1] && pos[0] == oxygenPos[0] {
			return node
		}
		switch val {
		case WALL:
			continue
		case VISITED:
			// expand
			for _, n := range []int{1, 2, 3, 4} {
				newPos := move_drone(pos, n)
				qu.Push(Node{newPos, append(node.Path, n)})
			}
		default:
			continue
		}
		space[pos[1]][pos[0]] = WALL
	}
	return Node{}
}

func bfs_fill(qu *queue.Queue, space [][]byte) int {
	draw_map(space, [2]int{-1, -1})
	newQ := queue.New()
	for !qu.Empty() {
		tmpNode := qu.Pop()
		node := tmpNode.(Node)
		pos := node.Pos
		val := space[pos[1]][pos[0]]
		switch val {
		case WALL:
			continue
		case OXYGEN:
			continue
		case VISITED:
			// expand
			space[pos[1]][pos[0]] = OXYGEN
			for _, n := range []int{1, 2, 3, 4} {
				newPos := move_drone(pos, n)
				if space[newPos[1]][newPos[0]] == VISITED {
					newQ.Push(Node{newPos, append(node.Path, n)})
				}
			}
		}
	}
	if newQ.Empty() {
		return 0
	}
	return 1 + bfs_fill(newQ, space)
}
