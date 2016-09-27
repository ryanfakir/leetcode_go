package src

import "math"

func isBalanced(root *TreeNode) bool {
	return getHeight(root) != -1
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := getHeight(root.Left)
	r := getHeight(root.Right)
	if l == -1 || r == -1 {
		return -1
	}
	if math.Abs(float64(l)-float64(r)) > 1 {
		return -1
	}
	return int(math.Max(float64(l), float64(r))) + 1
}
