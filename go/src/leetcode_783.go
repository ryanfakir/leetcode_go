var res int
func minDiffInBST(root *TreeNode) int {
    res = 1 << 15 -1
    dfs(root, -1 << 15, 1 << 15 -1)
    return res
}

func dfs(root *TreeNode, left, right int ) {
    if root == nil {return}
    if left != -1 << 15 {res = int(math.Min(float64(res), float64(root.Val - left)))}
    if right != 1 << 15 -1 {res = int(math.Min(float64(res), float64(right - root.Val)))}
    dfs(root.Left, left, root.Val)
    dfs(root.Right, root.Val, right)
}
