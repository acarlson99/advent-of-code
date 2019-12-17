package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	intcode "../libintcode"
)

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
	program := intcode.Read_program(reader)

	intcode.Exec_prog(intcode.Copy_arr(program), intcode.INTCStdin{}, intcode.INTCStdout{})
}
