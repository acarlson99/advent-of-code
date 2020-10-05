package main

func copy_byte_arr_arr(base [][]byte) [][]byte {
	var a [][]byte
	for _, line := range base {
		newLine := []byte{}
		for _, num := range line {
			newLine = append(newLine, num)
		}
		a = append(a, newLine)
	}
	return a
}
