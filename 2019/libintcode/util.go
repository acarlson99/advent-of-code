package libintcode

import (
	"io"
	"strconv"
	"strings"
)

func Copy_arr(base []int) []int {
	var a []int
	for _, num := range base {
		a = append(a, num)
	}
	return a
}

func GetArgs(arr, mode []int, ii, nargs, rel int) []int {
	ret := []int{}
	for jj := 0; jj < nargs; jj++ {
		n := 0
		switch mode[jj] {
		case 0:
			n = arr[ii+jj+1]
		case 1:
			n = ii + jj + 1
		case 2:
			n = arr[ii+jj+1] + rel
		default:
			panic("Invalid mode")
		}
		ret = append(ret, n)
	}
	return ret
}

func ParseOp(op int) ([]int, int) {
	modes := []int{}
	pos := 100
	for ii := 0; ii < 3; ii++ {
		modes = append(modes, (op/pos)%10)
		pos *= 10
	}
	// modes = append(modes, (op/100)%10)
	// modes = append(modes, (op/1000)%10)
	// modes = append(modes, (op/10000)%10)
	opcode := op % 100
	return modes, opcode
}

// return entire file as string
func Read_all(reader io.Reader) string {
	s := ""
	buf := make([]byte, 2048)
	for n, _ := reader.Read(buf); n > 0; n, _ = reader.Read(buf) {
		s += string(buf[0:n])
	}
	return s
}

// read from reader, return arr of nums
func Read_program(reader io.Reader) []int {
	ints := []int{}
	text := Read_all(reader)

	for _, line := range strings.Split(text, "\n") {
		ss := strings.Split(line, ",")

		intlen := len(ss) - 1
		lastnum := ss[intlen]
		if len(lastnum) > 0 && lastnum[len(lastnum)-1] == '\n' {
			lastnum = lastnum[0 : len(lastnum)-1]
			ss[intlen] = lastnum
		}

		for _, num := range ss {
			n, err := strconv.Atoi(num)
			if err != nil {
				break
			}
			ints = append(ints, n)
		}
	}

	for ii := 0; ii < 10000; ii++ {
		ints = append(ints, 0)
	}

	return ints
}
