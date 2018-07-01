func isValidBST(root *TreeNode) bool {
    fmt.Println(-1<<16)
    return helper(root, -1<<32, 1<<32)
}

func helper(root *TreeNode, left, right int) bool {
    if root == nil {return true}
    if root.Val <= left || root.Val >= right {return false}
    if helper(root.Left, left, root.Val) && helper(root.Right, root.Val, right) {
        return true
    }
    return false
}
