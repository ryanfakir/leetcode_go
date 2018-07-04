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



var res int = -1 << 16

func maxPathSum(root *TreeNode) int {
    if root == nil {return 0}
    res = root.Val
    helper(root)
    return res
}

func helper(root *TreeNode) int {
    if root == nil {return -1 << 16}
    lmax, rmax := helper(root.Left),helper(root.Right)
    l := int(math.Max(float64(lmax), 0.0))
    r := int(math.Max(float64(rmax), 0.0))
    res = int(math.Max(float64(res), float64(l + r + root.Val)))
    res = int(math.Max(float64(res), float64(lmax)))
    res = int(math.Max(float64(res), float64(rmax)))
    return int(math.Max(float64(l), float64(r))) + root.Val
}  
