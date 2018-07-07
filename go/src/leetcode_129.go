func sumNumbers(root *TreeNode) int {
    return helper(root, 0)
}

func helper(root *TreeNode, sum int) int {
    if root == nil {return 0}
    sum = sum * 10 + root.Val
    if root.Left == nil && root.Right == nil {
        return sum
    }
    left :=  helper(root.Left, sum)
    right :=  helper(root.Right, sum)
    return left + right
}
