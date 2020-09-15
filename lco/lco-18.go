package lco


func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}

	for pre, curr := head, head.Next; curr != nil; pre, curr = curr, curr.Next {
		if curr.Val == val {
			pre.Next = curr.Next
			break
		}
	}

	return head
}

func deleteNode2(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}

	curr := head
	for ; curr.Next != nil && curr.Next.Val != val; curr = curr.Next {
	}

	if curr.Next != nil {
		curr.Next = curr.Next.Next
	}

	return head
}
