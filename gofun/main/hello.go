package main

import (
	"sync"
	"fmt"
)

var lock sync.Mutex
var m map[string]string
var s []string


func main() {
	s1 := []int{0, 1, 2, 3}
	s2 := []int{4, 5, 6, 7}

	s1 = append(s1, s2...)
	fmt.Println(s1)
}

func CalcPerimeter(a, b, c int) int {

	return a + b + c

}

func CalcArea(a, b, c int) int {

	return 0
}
