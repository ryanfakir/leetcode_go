package main

import "math"

var res int = 1<<7 - 1
var pre int = -1

func getMinimumDifference(root *TreeNode) int {
	res = 1<<7 - 1
	helper(root)
	return res
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}
	helper(root.Left)
	if pre != -1 {
		res = int(math.Min(float64(res), float64(root.Val-pre)))
	}
	pre = root.Val
	helper(root.Right)
}
