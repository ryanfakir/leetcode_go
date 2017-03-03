package main

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var q []*TreeNode
	q = append(q, root)
	var res int
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			pop := q[0]
			q = q[1:]
			if pop.Left != nil && pop.Left.Left == nil && pop.Left.Right == nil {
				res += pop.Left.Val
			}
			if pop.Left != nil {
				q = append(q, pop.Left)
			}
			if pop.Right != nil {
				q = append(q, pop.Right)
			}
		}
	}
	return res
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	res = helper(root.Left, res, true)
	res = helper(root.Right, res, false)
	return res
}

func helper(node *TreeNode, res int, left bool) int {
	if node == nil {
		return res
	}
	if node.Left == nil && node.Right == nil && left {
		res += node.Val
	}
	res = helper(node.Left, res, true)
	res = helper(node.Right, res, false)
	return res
}
