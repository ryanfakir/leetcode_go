func buildTree(inorder []int, postorder []int) *TreeNode {
    return helper(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}


func helper(i, p []int, is, ie, ps, pe int) *TreeNode {
    if is > ie || ps > pe {return nil} 
    var j int
    for j = 0; j <= ie; j++ {
        if i[j] == p[pe] {break}
    }
    head := &TreeNode{Val:p[pe]}
    head.Left = helper(i, p, is, j-1, ps, ps + j -is-1)
    head.Right = helper(i, p, j+1, ie, ps+j -is, pe-1)
    return head
}
