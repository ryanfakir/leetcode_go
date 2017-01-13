package main

func hasPathSum(root *TreeNode, sum int) bool {
	return helper(0, root, sum)
}

func helper(val int, root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	root.Val = root.Val + val
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return true
		} else {
			return false
		}
	}
	if root.Left != nil {
		if helper(root.Val, root.Left, sum) {
			return true
		}
	}
	if root.Right != nil {
		if helper(root.Val, root.Right, sum) {
			return true
		}
	}
	return false
}
