package main

import (
	"fmt"
	"gofun/letcode/LongestCommonPrefix"
 )

func main() {

	strs:=make([]string,1)
	strs = append(strs,"a")
	strs = append(strs,"ab")
	strs = append(strs,"abc")

	res:=LongestCommonPrefix.LongestCommonPrefix(strs)
	fmt.Println(res)


}