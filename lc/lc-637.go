package lc

func averageOfLevels(root *TreeNode) []float64 {
	queue := make([]*TreeNode, 1)
	queue[0] = root
	sum, count := 0, 1
	ans := make([]float64, 0)
	var cur *TreeNode
	var limit, nc int
	for len(queue) > 0 {
		limit = count
		count = 0
		sum = 0
		for i := 0; i < limit; i++ {
			cur = queue[0]
			queue = queue[1:]
			if cur != nil {
				sum += cur.Val
				queue = append(queue, cur.Left)
				queue = append(queue, cur.Right)
				count += 2
			} else {
				nc++
			}
		}
		if limit != nc {
			ans = append(ans, float64(sum)/float64(limit-nc))
		}
	}
	return ans
}
