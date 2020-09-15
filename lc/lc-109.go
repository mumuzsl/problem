package lc



func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	length := 0
	h := head
	p := head
	for p != nil {
		length++
		p = p.Next
	}

	var buildBST func(start, end int) *TreeNode
	buildBST = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := (start + end) >> 1
		left := buildBST(start, mid-1)
		root := new(TreeNode)
		root.Val = h.Val
		h = h.Next
		root.Left = left
		root.Right = buildBST(mid+1, end)
		return root
	}

	return buildBST(0, length-1)
}

func sortedListToBST2(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	var preSlow *ListNode

	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	root := new(TreeNode)
	root.Val = slow.Val

	if preSlow != nil {
		preSlow.Next = nil
		//root.Left = sortedListToBST(head)
	}

	if head == slow {
		return root
	}

	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(slow.Next)
	return root
}
