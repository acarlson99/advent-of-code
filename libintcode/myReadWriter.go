package libintcode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type MyWriter interface {
	// write int to storage.  Undefined if nothing to read
	WriteInt(int)
	// mark that no more output will be passed to WriteInt
	Close()
}

type MyReader interface {
	// read and return int from storage, return true if more info to be read
	ReadInt() (int, bool)
}

type MyReadWriter interface {
	MyReader
	MyWriter
}

type MyChan chan int

func (ch MyChan) ReadInt() (int, bool) {
	a, b := <-ch
	return a, b
}

func (ch MyChan) WriteInt(n int) {
	ch <- n
}

func (ch MyChan) Close() {
	close(ch)
}

// stdin/stdout struct
type MyStdin struct{}

func (r MyStdin) ReadInt() (int, bool) {
	fmt.Printf("> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if len(text) > 0 {
		n, _ := strconv.Atoi(text[0 : len(text)-1])
		return n, true
	}
	return 0, true
}

type MyStdout struct{}

func (r MyStdout) WriteInt(n int) {
	fmt.Println(n)
}

func (r MyStdout) Close() {
}

// single int input
type MyInt int

func (ii MyInt) ReadInt() (int, bool) {
	return int(ii), true
}

func (ii MyInt) Close() {
}

// array of ints
type MyArr struct {
	a  []int
	ii int
}

func (a MyArr) ReadInt() (int, bool) {
	num := 0
	if a.ii < len(a.a) {
		num = a.a[a.ii]
		a.ii++
	} else {
		num = a.a[len(a.a)-1]
	}
	return num, true
}

func (a MyArr) Close() {
}
