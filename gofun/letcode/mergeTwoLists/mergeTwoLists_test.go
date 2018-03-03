package mergeTwoLists

import "testing"
import "fmt"

func TestMergeTwoLists(t *testing.T) {
	l1:=&ListNode{1,&ListNode{2,&ListNode{3,nil}}}
	l2:=&ListNode{1,&ListNode{2,&ListNode{3,nil}}}


	res:=MergeTwoLists(l1,l2)

	for res!=nil {
		fmt.Println(res.Val)
		res = res.Next
	}

}
