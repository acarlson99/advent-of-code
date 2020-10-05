package main

import (
	"bufio"
	"math"
	"os"

	intcode "../libintcode"
)

func PowerSet(original []string) [][]string {
	powerSetSize := int(math.Pow(2, float64(len(original))))
	result := make([][]string, 0, powerSetSize)
	var index int
	for index < powerSetSize {
		var subSet []string
		for j, elem := range original {
			if index&(1<<uint(j)) > 0 {
				subSet = append(subSet, elem)
			}
		}
		result = append(result, subSet)
		index++
	}
	return result
}

func main() {
	solve := true

	getthere := ""
	if solve {
		items := []string{
			"easter egg",
			"mug",
			"sand",
			"weather machine",
			"festive hat",
			"shell",
			"whirled peas",
			"space heater",
		}
		pset := PowerSet(items)

		getthere = "west\nwest\nnorth\ntake heater\ntake space heater\nsouth\neast\nsouth\ntake festive hat\nsouth\ntake sand\nnorth\neast\ntake whirled peas\nwest\nnorth\neast\nsouth\ntake weather machine\ninv\nnorth\neast\ntake mug\neast\nsouth\nwest\nsouth\nwest\ntake shell\nsouth\nsouth\nsouth\nnorth\neast\nnorth\neast\neast\nnorth\nwest\neast\nsouth\nsouth\nnorth\nsouth\ntake easter egg\nnorth\ninv\nwest\nwest\nsouth\nwest\nsouth\nsouth\nnorth\ninv\nsouth\nnorth\neast\nnorth\neast\nwest\neast\nwest\nsouth\nwest\nsouth\nsouth\nnorth\neast\nnorth\neast\neast\nnorth\nwest\neast\nsouth\nsouth\nsouth\nnorth\nwest\nnorth\nwest\nwest\nsouth\nnorth\nwest\nwest\nnorth\nsouth\neast\neast\neast\neast\nsouth\nwest\nsouth\nwest\nsouth\nsouth\n"

		for _, s := range pset {
			for _, item := range s {
				getthere = getthere + "drop " + item + "\n"
			}
			getthere += "south\n"
			for _, item := range s {
				getthere = getthere + "take " + item + "\n"
			}
		}
	}

	inFile, err := os.Open("input.txt")
	if err != nil {
		panic(err) // TODO: address error
	}
	reader := bufio.NewReader(inFile)
	program := intcode.Read_program(reader)
	intcode.Exec_prog(intcode.Copy_arr(program), &intcode.INTCASCIIStdin{[]byte(getthere)}, &intcode.INTCASCIIStdout{})
}
