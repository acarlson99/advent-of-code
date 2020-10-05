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

type INTCStdoutSep struct {
	Sep string
	End string
	ii  int
}

func NewStdoutSep(sep, end string) *INTCStdoutSep {
	return &INTCStdoutSep{sep, end, 0}
}

func (r *INTCStdoutSep) WriteInt(n int) {
	if r.ii != 0 {
		fmt.Printf(r.Sep)
	}
	r.ii++
	fmt.Printf("%d", n)
}

func (r *INTCStdoutSep) Close() {
	fmt.Printf(r.End)
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
	A  []int
	II int
}

func (a *INTCArr) ReadInt() (int, bool) {
	num := 0
	if a.II < len(a.A) {
		num = a.A[a.II]
		a.II++
	} else {
		num = a.A[len(a.A)-1]
	}
	return num, true
}

func (a *INTCArr) Close() {
}

type INTCArrWriter struct {
	A *[]int
}

func (a INTCArrWriter) WriteInt(n int) {
	*a.A = append(*a.A, n)
}

func (a INTCArrWriter) Close() {
}

type INTCASCIIStdout struct {
	Buf []byte
}

func (a *INTCASCIIStdout) WriteInt(n int) {
	if n == 10 {
		fmt.Println(string(a.Buf))
		a.Buf = []byte{}
	} else {
		a.Buf = append(a.Buf, byte(n))
	}
}

func (a *INTCASCIIStdout) Close() {}

type INTCASCIIStdin struct {
	Buf []byte
}

func (a *INTCASCIIStdin) ReadInt() (int, bool) {
	if len(a.Buf) == 0 {
		fmt.Printf("> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		a.Buf = []byte(text)
	}
	c := a.Buf[0]
	a.Buf = a.Buf[1:]
	return int(c), true
}
