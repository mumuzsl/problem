package lco

func mirrorTree(root *TreeNode) *TreeNode {
	queue := make([]*TreeNode, 1)
	queue[0] = root
	var curr *TreeNode
	for len(queue) > 0 {
		curr = queue[0]
		queue = queue[1:]
		if curr == nil {
			continue
		}
		curr.Left, curr.Right = curr.Right, curr.Left
		queue = append(queue, curr.Left)
		queue = append(queue, curr.Right)
	}

	return root
}
