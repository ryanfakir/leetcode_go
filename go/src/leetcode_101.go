package src

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

func helper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left != nil && right != nil {
		return left.Val == right.Val && helper(left.Left, right.Right) && helper(left.Right, right.Left)
	} else {
		return false
	}
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

func helper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if !(left != nil && right != nil) || left.Val != right.Val {
		return false
	}
	return helper(left.Right, right.Left) && helper(left.Left, right.Right)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var q []*TreeNode
	if (root.Left == nil && root.Right != nil) || (root.Left != nil && root.Right == nil) {
		return false
	}
	if root.Left != nil && root.Right != nil {
		q = append(q, root.Left)
		q = append(q, root.Right)
	}
	for len(q) > 0 {
		first := q[0]
		sec := q[1]
		q = q[2:]
		if first.Val != sec.Val {
			return false
		}
		if (first.Left == nil && sec.Right != nil) || (first.Left != nil && sec.Right == nil) {
			return false
		}
		if first.Left != nil && sec.Right != nil {
			q = append(q, first.Left)
			q = append(q, sec.Right)
		}
		if (first.Right == nil && sec.Left != nil) || (first.Right != nil && sec.Left == nil) {
			return false
		}
		if first.Right != nil && sec.Left != nil {
			q = append(q, first.Right)
			q = append(q, sec.Left)
		}

	}
	return true
}
