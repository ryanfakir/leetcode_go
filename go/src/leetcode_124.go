package main

import "math"

func maxPathSum(root *TreeNode) int {
	var max = math.MinInt8
	helper(root, &max)
	return max
}

func helper(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := int(math.Max(float64(helper(root.Left, max)), 0))
	right := int(math.Max(float64(helper(root.Right, max)), 0))
	*max = int(math.Max(float64(*max), float64(left+right+root.Val)))
	return int(math.Max(float64(left), float64(right))) + root.Val
}
