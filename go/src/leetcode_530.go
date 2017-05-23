package main

import "math"

var res int

func getMinimumDifference(root *TreeNode) int {
	helper(root, -1<<8)
	return res
}

func helper(root *TreeNode, pre int) {
	if root == nil {
		return
	}
	helper(root.Left)
	res = int(math.Min(float64(res), float64(root.Val-left)))
	helper(root.Right)
	res = int(math.Min(float64(res), float64(right-root.Val)))

}
