package links

import (
	"testing"
	"fmt"
)

func TestExtract(t *testing.T) {
	res,err:=Extract("http://gopl.io")
	if nil == err {
		fmt.Println(res)
	}else{
		fmt.Println(err)
	}
}
