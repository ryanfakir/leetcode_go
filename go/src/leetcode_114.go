package main

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	var left, right *TreeNode
	if root.Left != nil {
		left = root.Left
		root.Left = nil
	}
	flatten(root.Right)
	if left != nil {
		right = root.Right
		root.Right = left
		for left.Right != nil {
			left = left.Right
		}
		left.Right = right
	}
}
