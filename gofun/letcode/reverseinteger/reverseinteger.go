package reverseinteger

import (
	"strconv"
)
var MIN int = -0x80000000
var MAX int = 0x7FFFFFFF


func reverse(x int) int {

	var sr string
	if x>0 {
		sr = ReverseString(strconv.Itoa(x))
	}else{
		sr = "-"+ReverseString(strconv.Itoa(-x))
	}

	res,err:=strconv.Atoi(sr)

	if err == nil && res<MAX && res>MIN{
		return res
	}

	return 0
}



func ReverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}