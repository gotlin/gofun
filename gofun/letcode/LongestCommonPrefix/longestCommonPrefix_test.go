package LongestCommonPrefix

import "testing"
import "fmt"

func TestLongestCommonPrefix(t *testing.T) {

	strs:=make([]string,0)
	strs = append(strs,"ab")
	strs = append(strs,"abd")
	strs = append(strs,"abc")

	res:=LongestCommonPrefix(strs)
	fmt.Println(res)
}
