package lc

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	root := new(ListNode)
	cur := root
	var a, b, s int
	p1, p2 := l1, l2
	for {
		a, b = 0, 0
		if p1 != nil {
			a = p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			b = p2.Val
			p2 = p2.Next
		}
		s, carry = a+b+carry, 0
		if s >= 10 {
			carry = 1
		}
		cur.Val = s % 10
		if p1 == nil && p2 == nil && carry == 0 {
			break
		}
		cur.Next = new(ListNode)
		cur = cur.Next
	}

	return root
}
