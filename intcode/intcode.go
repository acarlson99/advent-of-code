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
	var stdin bool
	var file, stringIn string
	flag.StringVar(&file, "f", "", "Read from file")
	flag.StringVar(&stringIn, "s", "", "Read from string")
	flag.BoolVar(&stdin, "i", false, "Read from stdin")

	var sep, end string
	flag.StringVar(&end, "end", "\n", "End of input string")
	flag.StringVar(&sep, "sep", "\n", "Separator string")

	var assemble, disassemble, evaluate bool
	flag.BoolVar(&assemble, "a", false, "Assemble input")
	flag.BoolVar(&disassemble, "d", false, "Disassemble input")
	flag.BoolVar(&evaluate, "e", false, "Evaluate input")

	flag.Parse()

	var reader io.Reader
	if file != "" {
		inFile, err := os.Open(file)
		if err != nil {
			panic(err) // TODO: address error
		}
		reader = bufio.NewReader(inFile)
	} else if stringIn != "" {
		reader = strings.NewReader(stringIn)
	} else if stdin {
		reader = bufio.NewReader(os.Stdin)
	} else {
		fmt.Println("Specify input method (stdin, string, file)")
		flag.Usage()
		os.Exit(1)
	}

	if evaluate {
		program := intcode.Read_program(reader)
		intcode.Exec_prog(intcode.Copy_arr(program), &intcode.INTCStdin{}, intcode.NewStdoutSep(sep, end))
	} else if assemble {
		fmt.Println("Assembly not yet implemented")
	} else if disassemble {
		fmt.Println("Disassembly not yet implemented")
	} else {
		fmt.Println("Specify action (asm, disasm, eval)")
		flag.Usage()
		os.Exit(1)
	}
}
