package main

import (
	"fmt"
	"math"
)

type Triple struct {
	x int
	y int
	z int
}

type Body struct {
	p Triple

	v Triple
}

func applyGrav(ii int, bodies []Body) Triple {
	v := bodies[ii].v
	p := bodies[ii].p

	for jj, body := range bodies {
		if ii == jj {
			continue
		}

		bp := body.p

		if p.x < bp.x {
			v.x += 1
		} else if p.x > bp.x {
			v.x -= 1
		}

		if p.y < bp.y {
			v.y += 1
		} else if p.y > bp.y {
			v.y -= 1
		}

		if p.z < bp.z {
			v.z += 1
		} else if p.z > bp.z {
			v.z -= 1
		}
	}
	return v
}

func applyVelocity(body Body) Body {
	p := body.p
	v := body.v
	return Body{Triple{p.x + v.x, p.y + v.y, p.z + v.z}, v}
}

func stepOne(bodies []Body) []Body {
	newBods := []Body{}
	for ii, body := range bodies {
		velocity := applyGrav(ii, bodies)
		b := applyVelocity(Body{body.p, velocity})
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

func sumTriple(t Triple) int {
	return int(math.Abs(float64(t.x)) + math.Abs(float64(t.y)) + math.Abs(float64(t.z)))
}

func energy(bodies []Body) int {
	total := 0
	for _, body := range bodies {
		pot := sumTriple(body.v)
		kin := sumTriple(body.p)
		total += pot * kin
	}
	return total
}

func main() {
	bodies := []Body{
		Body{Triple{5, 4, 4}, Triple{0, 0, 0}},
		Body{Triple{-11, -11, -3}, Triple{0, 0, 0}},
		Body{Triple{0, 7, 0}, Triple{0, 0, 0}},
		Body{Triple{-13, 2, 10}, Triple{0, 0, 0}}}

	new := step(bodies, 1000)
	fmt.Println("Part one:", energy(new))
}
