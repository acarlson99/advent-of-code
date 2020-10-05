package main

import (
	"bufio"
	"fmt"
	"os"
)

type Body struct {
	P []int

	V []int
}

func applyGrav(ii int, bodies []Body) []int {
	v := append([]int(nil), bodies[ii].V...)
	p := append([]int(nil), bodies[ii].P...)

	for jj, body := range bodies {
		if ii == jj {
			continue
		}

		bp := body.P

		for kk := range p {
			if p[kk] < bp[kk] {
				v[kk] += 1
			} else if p[kk] > bp[kk] {
				v[kk] -= 1
			}
		}
	}
	return v
}

func applyVelocity(body Body) Body {
	p := body.P
	v := body.V
	return Body{[]int{p[0] + v[0], p[1] + v[1], p[2] + v[2]}, v}
}

func stepOne(bodies []Body) []Body {
	newBods := []Body{}
	for ii, body := range bodies {
		velocity := applyGrav(ii, bodies)
		b := applyVelocity(Body{body.P, velocity})
		newBods = append(newBods, b)
	}
	return newBods
}

func step(bodies []Body, n int) []Body {
	for ii := 0; ii < n; ii++ {
		bodies = stepOne(bodies)
	}
	return bodies
}

func energy(bodies []Body) int {
	total := 0
	for _, body := range bodies {
		pot := sumArr(body.V...)
		kin := sumArr(body.P...)
		total += pot * kin
	}
	return total
}

func findRep(bodies []Body) int {
	coords := []int{0, 0, 0}

	initial := bodies
	done := false
	for ii := 1; !done; ii++ {
		done = true
		bodies = stepOne(bodies)

		for jj := range coords {
			if coords[jj] == 0 {
				good := true
				for idx, body := range bodies {
					if body.V[jj] != initial[idx].V[jj] {
						good = false
						break
					}
				}
				if good {
					coords[jj] = ii * 2
				}
			}
		}

		for _, c := range coords {
			if c == 0 {
				done = false
				break
			}
		}
	}

	return lcm(coords...)
}

func main() {
	bodies := []Body{}

	reader := bufio.NewReader(os.Stdin)
	for text, _ := reader.ReadString('\n'); len(text) > 0; text, _ = reader.ReadString('\n') {
		newBody := Body{[]int{0, 0, 0}, []int{0, 0, 0}}
		fmt.Sscanf(text, "<x=%d, y=%d, z=%d>", &newBody.P[0], &newBody.P[1], &newBody.P[2])
		bodies = append(bodies, newBody)
	}

	fmt.Println("Part one:", energy(step(bodies, 1000)))
	fmt.Println("Part two:", findRep(bodies))
}
