func generateTrees(n int) []*TreeNode {
    if n == 0 {return nil}
    return helper(1, n)
}

func helper(start, end int) []*TreeNode {
    if start > end {
        return []*TreeNode{nil}
    }
    var res []*TreeNode
    for i:= start; i<=end ; i++ {
        Left := helper(start, i-1)
        Right := helper(i+1, end)
        for _, l := range Left {
            for _, r := range Right {
                head := &TreeNode{Val:i}
                head.Left = l
                head.Right = r
                res = append(res, head)
            }
        }
    }
    return res
}
