package main

import (
	"bufio"
	"fmt"
	"os"

	intcode "../libintcode"
)

func str_to_int_arr(str string) []int {
	r := []int{}
	for _, c := range str {
		r = append(r, int(c))
	}
	return r
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	program := intcode.Read_program(reader)

	{
		a := str_to_int_arr("OR A J\nAND B J\nAND C J\nNOT J J\nAND D J\nWALK\n")
		input := intcode.INTCArr{a, 0}
		b := []int{}
		output := intcode.INTCArrWriter{&b}
		intcode.Exec_prog(intcode.Copy_arr(program), &input, output)
		for _, c := range b {
			if c <= 'z' {
				// fmt.Printf("%c", c)
			} else {
				fmt.Println("Part one:", c)
			}
		}
	}
	{
		a := str_to_int_arr("OR A J\nAND B J\nAND C J\nNOT J J\nAND D J\nOR E T\nOR H T\nAND T J\nRUN\n")
		input := intcode.INTCArr{a, 0}
		b := []int{}
		output := intcode.INTCArrWriter{&b}
		intcode.Exec_prog(intcode.Copy_arr(program), &input, output)
		for _, c := range b {
			if c <= 'z' {
				// fmt.Printf("%c", c)
			} else {
				fmt.Println("Part two:", c)
			}
		}
	}
}
