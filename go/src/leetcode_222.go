func countNodes(root *TreeNode) int {
    left, right := root, root
    var lh, rh int
    for left != nil {
        lh++
        left = left.Left
    }
    for right != nil {
        rh++
        right = right.Right
    }
    if rh == lh {return int(math.Pow(2, float64(rh)))- 1}
    return countNodes(root.Left) + countNodes(root.Right) + 1
}
