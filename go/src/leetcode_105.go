func buildTree(preorder []int, inorder []int) *TreeNode {
    return helper(preorder, inorder, 0 , len(preorder)-1, 0, len(inorder)-1)
}


func helper(p, i []int, ps, pe, is, ie int) *TreeNode {
    if ps > pe || is > ie {return nil}
    var j int
    for j = is; j <= ie; j++ {
        if p[ps] == i[j] {break}
    }
    head := &TreeNode{Val:p[ps]}
    head.Left = helper(p, i, ps+1, ps + j - is, is, j-1)
    head.Right = helper(p, i, ps+j -is +1, pe, j+1, ie)
    return head
}
