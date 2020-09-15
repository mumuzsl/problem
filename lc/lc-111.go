package lc

import "fmt"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	layerQueue := []int{1}
	r := 1

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		i := layerQueue[0]
		layerQueue = layerQueue[1:]
		fmt.Println(q)
		fmt.Println(i)
		if q.Left == nil && q.Right == nil {
			return i
		}
		if q.Left != nil {
			queue = append(queue, q.Left)
			layerQueue = append(layerQueue, i+1)
		}
		if q.Right != nil {
			queue = append(queue, q.Right)
			layerQueue = append(layerQueue, i+1)
		}
	}

	return r
}
