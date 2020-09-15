package lco

func isSymmetric(root *TreeNode) bool {
	var recur func(L *TreeNode, R *TreeNode) bool
	recur = func(L *TreeNode, R *TreeNode) bool {
		if L == nil && R == nil {
			return true
		}
		if L == nil || R == nil || L.Val != R.Val {
			return false
		}
		return recur(L.Left, R.Right) && recur(L.Right, R.Left)
	}

	if root != nil {
		return recur(root.Left, root.Right)
	}
	return true
}
