package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type myWriter interface {
	// write int to storage.  Undefined if nothing to read
	WriteInt(int)
	// mark that no more output will be passed to WriteInt
	Close()
}

type myReader interface {
	// read and return int from storage, return true if more info to be read
	ReadInt() (int, bool)
}

type myReadWriter interface {
	myReader
	myWriter
}

type myChan chan int

func (ch myChan) ReadInt() (int, bool) {
	a, b := <-ch
	return a, b
}

func (ch myChan) WriteInt(n int) {
	ch <- n
}

func (ch myChan) Close() {
	close(ch)
}

// stdin/stdout struct
type myStdin struct{}

func (r myStdin) ReadInt() (int, bool) {
	fmt.Printf("> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if len(text) > 0 {
		n, _ := strconv.Atoi(text[0 : len(text)-1])
		return n, true
	}
	return 0, true
}

type myStdout struct{}

func (r myStdout) WriteInt(n int) {
	fmt.Println(n)
}

func (r myStdout) Close() {
}

// single int input
type myInt int

func (ii myInt) ReadInt() (int, bool) {
	return int(ii), true
}

func (ii myInt) Close() {
}

// array of ints
type myArr struct {
	a  []int
	ii int
}

func (a myArr) ReadInt() (int, bool) {
	num := 0
	if a.ii < len(a.a) {
		num = a.a[a.ii]
		a.ii++
	} else {
		num = a.a[len(a.a)-1]
	}
	return num, true
}

func (a myArr) Close() {
}
