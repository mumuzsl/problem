package lco

func reverseList(head *ListNode) *ListNode {
	var p1, p2, p3 *ListNode

	p1 = head
	p3 = nil
	for ; p1 != nil; p1 = p2 {
		p2, p1.Next, p3 = p1.Next, p3, p1
	}

	return p3
}
