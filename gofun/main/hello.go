package main

import (
	"sync"
	"fmt"
)

var lock sync.Mutex
var m map[string]string
var s []string

type T struct{
	a *[]string
}

func main() {

	t1:=T{}
	t2:=T{}

	fmt.Println(t1.a)
	fmt.Println(t1.a==t2.a)
}

func CalcPerimeter(a, b, c int) int {

	return a + b + c

}

func CalcArea(a, b, c int) int {

	return 0
}
