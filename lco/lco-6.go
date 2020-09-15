package lco

//Definition for singly-linked list.
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	temp := make([]int, 0)
	result := make([]int, 0)

	for p := head; p != nil; p = p.Next {
		temp = append(temp, p.Val)
	}

	tempLen := len(temp)

	for i := tempLen - 1; i >= 0; i-- {
		result = append(result, temp[i])
	}

	return result
}
