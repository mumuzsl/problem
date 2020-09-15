package lco

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	table := make(map[int]int, 0)

	for i, v := range inorder {
		table[v] = i
	}

	root := new(TreeNode)
	root.Val = preorder[0]
	preorderLen := len(preorder)
	//queue := []*TreeNode{root}
	var curr *TreeNode

	for i := 1; i < preorderLen; i++ {
		curr = root
		newNode := new(TreeNode)
		newNode.Val = preorder[i]
		for {
			if table[newNode.Val] > table[curr.Val] {
				if curr.Right == nil {
					curr.Right = newNode
					break
				}
				curr = curr.Right
			} else {
				if curr.Left == nil {
					curr.Left = newNode
					break
				}
				curr = curr.Left
			}
		}
	}

	return root
}

func buildTree2(preorder []int, inorder []int) *TreeNode {
	table := make(map[int]int, 0)

	for i, v := range inorder {
		table[v] = i
	}

	var recur func(preRoot, inLeft, inRight int) *TreeNode
	recur = func(preRoot, inLeft, inRight int) *TreeNode {
		if inLeft > inRight {
			return nil
		}
		root := new(TreeNode)
		root.Val = preorder[preRoot]
		i := table[root.Val]
		root.Left = recur(preRoot+1, inLeft, i-1)
		root.Right = recur(i-inLeft+preRoot+1, i+1, inRight)
		return root
	}

	return recur(0, 0, len(inorder)-1)
}
