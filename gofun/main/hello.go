package main

import (
	"sync"
	"fmt"
)

var lock sync.Mutex
var m map[string]string
var s []string

type T struct {
	a string
	b string
}

var Wg sync.WaitGroup
func main() {


	A:=T{"a","b"}
	B:=A
	B.a = "b"

	fmt.Println(A)
	fmt.Println(B)

}

func CalcPerimeter(a, b, c int) int {

	return a + b + c

}

func CalcArea(a, b, c int) int {

	return 0
}
