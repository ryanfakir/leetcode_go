func upsideDownBinaryTree(root *TreeNode) *TreeNode {
    var r, f *TreeNode
    for root != nil {
        left := root.Left
        right := root.Right
        root.Left = r
        root.Right = f
        f = root
        r = right
        root = left
    }
    return f
}
