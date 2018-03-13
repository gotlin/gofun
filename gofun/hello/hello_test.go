package hello

import "testing"
import (
	"fmt"
	"runtime"
	"os"
)

type Adder interface {
	add(x int) int
}

type Intt int

func (i Intt) add(x int) int {

	return int(i) + x
}

func DoAdd(ad *Adder, d int) int {
	return (*ad).add(d)
}

func TestHello(t *testing.T) {
	var adder Adder
	var in Intt = 1
	adder = in
	DoAdd(&adder, 1)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
