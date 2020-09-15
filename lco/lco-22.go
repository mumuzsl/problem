package lco


func getKthFromEnd(head *ListNode, k int) *ListNode {
	count := 0
	for cur := head; cur != nil; cur = cur.Next {
		count++
	}
	r := head
	for i := 0; i < count-k; i++ {
		r = r.Next
	}
	return r
}

func getKthFromEnd2(head *ListNode, k int) *ListNode {
	former, latter := head, head
	for i := 0; i < k; i++ {
		former = former.Next
		if former == nil {
			return head
		}
	}
	for former != nil {
		former, latter = former.Next, latter.Next
	}
	return latter
}
