package lc

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		sum += root.Val
		root.Val = sum
		dfs(root.Left)
	}
	dfs(root)

	return root
}
