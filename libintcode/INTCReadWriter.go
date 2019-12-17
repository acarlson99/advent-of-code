package libintcode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type INTCWriter interface {
	// write int to storage.  Undefined if nothing to read
	WriteInt(int)
	// mark that no more output will be passed to WriteInt
	Close()
}

type INTCReader interface {
	// read and return int from storage, return true if more info to be read
	ReadInt() (int, bool)
}

type INTCReadWriter interface {
	INTCReader
	INTCWriter
}

type INTCChan chan int

func (ch INTCChan) ReadInt() (int, bool) {
	a, b := <-ch
	return a, b
}

func (ch INTCChan) WriteInt(n int) {
	ch <- n
}

func (ch INTCChan) Close() {
	close(ch)
}

// stdin/stdout struct
type INTCStdin struct{}

func (r INTCStdin) ReadInt() (int, bool) {
	fmt.Printf("> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if len(text) > 0 {
		n, _ := strconv.Atoi(text[0 : len(text)-1])
		return n, true
	}
	return 0, true
}

type INTCStdout struct{}

func (r INTCStdout) WriteInt(n int) {
	fmt.Println(n)
}

func (r INTCStdout) Close() {
}

// single int input
type INTCInt int

func (ii INTCInt) ReadInt() (int, bool) {
	return int(ii), true
}

func (ii INTCInt) Close() {
}

// array of ints
type INTCArr struct {
	a  []int
	ii int
}

func (a INTCArr) ReadInt() (int, bool) {
	num := 0
	if a.ii < len(a.a) {
		num = a.a[a.ii]
		a.ii++
	} else {
		num = a.a[len(a.a)-1]
	}
	return num, true
}

func (a INTCArr) Close() {
}
