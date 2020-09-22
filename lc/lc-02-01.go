package lc

func removeDuplicateNodes(head *ListNode) *ListNode {
	table := make(map[int]int)
	pre, cur := head, head
	for cur != nil {
		if table[cur.Val] == 0 {
			table[cur.Val] = 1
			pre = cur
		} else {
			pre.Next = cur.Next
		}
		cur = cur.Next
	}
	return head
}
