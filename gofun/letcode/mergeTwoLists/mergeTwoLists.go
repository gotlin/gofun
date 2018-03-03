package mergeTwoLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	res := &ListNode{0, nil}
	cur := res

	for ; l1 != nil && l2 != nil; {

		if l1.Val < l2.Val {
			cur.Next = &ListNode{l1.Val, nil}
			cur = cur.Next
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{l2.Val, nil}
			cur = cur.Next
			l2 = l2.Next
		}

	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}

	return res.Next

}
