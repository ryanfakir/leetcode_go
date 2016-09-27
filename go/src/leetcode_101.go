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