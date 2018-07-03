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

func flatten(root *TreeNode)  {
    var q []*TreeNode
    for root != nil || len(q) > 0 {
        if root.Left != nil {
            if root.Right != nil {
                q = append(q, root.Right)
            }
            root.Right = root.Left
            root.Left = nil
        }
        if root.Right == nil  {
            if len(q) > 0 {
                pop := q[len(q)-1]
                q = q[:len(q)-1]
                root.Right = pop
            } else {
                return
            }
        } else {
            root = root.Right
        }
    }
}
