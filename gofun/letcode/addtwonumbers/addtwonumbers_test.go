package addtwonumbers

import (
	"testing"
	"fmt"
)

func Test_addTwoNumbers(t *testing.T) {

	l1 := &ListNode{2, &ListNode{9, &ListNode{4, nil}}}
	l2 := &ListNode{2, &ListNode{3, &ListNode{4, nil}}}

	res := addTwoNumbers(l1, l2)
	for ;res!=nil;{

		fmt.Println(res.Val)
		res=res.Next
	}

}
