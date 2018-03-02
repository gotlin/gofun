package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func (t ListNode) String() string {
	res := ""
	head := &t
	for ; head != nil; {
		res = res + string(t.Val)
		head = head.Next
	}
	return res
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	res := &ListNode{0, nil}
	head := res
	jinwei := 0

	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 == nil {
			l1 = &ListNode{0, nil}
		}

		if l2 == nil {
			l2 = &ListNode{0, nil}
		}

		shiwei := (l1.Val + l2.Val + jinwei) / 10
		gewei := (l1.Val + l2.Val + jinwei) % 10
		l1 = l1.Next
		l2 = l2.Next
		jinwei = shiwei
		head.Next = &ListNode{gewei, nil}
		head = head.Next
	}

	if jinwei != 0 {
		head.Next = &ListNode{jinwei, nil}
	}

	return res.Next

}

func multipleOfIndex(ints []int) []int {

	res := make([]int, 1)
	for i, v := range ints[1:] {
		if (v != 0 || v%i == 0) {
			res = append(res, v)
		}
	}

	return res
}
