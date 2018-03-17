package main

import (
	"sync"
	"fmt"
)

var lock sync.Mutex
var m map[string]string
var s []string


func main() {
	s1 := [][]int{{0, 1, 2, 3},{},{1}}

	fmt.Println(s1)
}

func CalcPerimeter(a, b, c int) int {

	return a + b + c

}

func CalcArea(a, b, c int) int {

	return 0
}
