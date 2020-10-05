package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type myReadWriter interface {
	// read and return int from storage, return true if more info to be read
	ReadInt() (int, bool)
	// write int to storage.  Undefined if nothing to read
	WriteInt(int)
	// mark that no more output will be passed to WriteInt
	Close()
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

type myStdin struct{}

func (r myStdin) ReadInt() (int, bool) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(text[0 : len(text)-1])
	return n, false
}

func (r myStdin) WriteInt(n int) {
	fmt.Println(n)
}

func (r myStdin) Close() {
}
