package LongestCommonPrefix

import "testing"
import "fmt"

func TestLongestCommonPrefix(t *testing.T) {

	strs:=make([]string,1)
	strs = append(strs,"a")
	strs = append(strs,"ab")
	strs = append(strs,"abc")

	res:=LongestCommonPrefix(strs)
	fmt.Println(res)
}
