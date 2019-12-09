package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// return entire file as string
func read_all(reader io.Reader) string {
	s := ""
	buf := make([]byte, 2048)
	for n, _ := reader.Read(buf); n > 0; n, _ = reader.Read(buf) {
		s += string(buf[0:n])
	}
	return s
}

// read from reader, return arr of nums
func read_program(reader io.Reader) []int {
	ints := []int{}
	text := read_all(reader)

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

	fmt.Println(ints[0:30])

	return ints
}

func main() {
	// setup
	var file, string, stdin bool
	flag.BoolVar(&file, "f", false, "Read from file")
	flag.BoolVar(&string, "s", false, "Read from string")
	flag.BoolVar(&stdin, "i", false, "Read from stdin")
	flag.Parse()

	args := flag.Args()

	var reader io.Reader
	if file {
		if len(args) != 1 {
			fmt.Println("usage: ./intcode -f filename")
			flag.Usage()
			os.Exit(1)
		}
		inFile, err := os.Open(args[0])
		if err != nil {
			panic(err) // TODO: address error
		}
		reader = bufio.NewReader(inFile)
	} else if string {
		if len(args) != 1 {
			fmt.Println("usage: ./intcode -s filename")
			flag.Usage()
			os.Exit(1)
		}
		reader = strings.NewReader(args[0])
	} else if stdin {
		reader = bufio.NewReader(os.Stdin)
	} else {
		flag.Usage()
		os.Exit(1)
	}
	program := read_program(reader)

	exec_prog(copy_arr(program), myStdin{}, myStdout{})
}
