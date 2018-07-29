var res int 
func maxDepth(root *TreeNode) int {
    res = 0
    max(root, 0)
    return res
}

func max(root *TreeNode, tmp int) {
    if root == nil {
        res = int(math.Max(float64(res), float64(tmp)))
        return
    }
    max(root.Left, tmp+1)
    max(root.Right, tmp + 1)
}
