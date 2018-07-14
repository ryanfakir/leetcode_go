func rightSideView(root *TreeNode) []int {
    if root == nil {return nil}
    var q []*TreeNode = []*TreeNode{root}
    var res []int
    for len(q) >0 {
        size := len(q)
        res = append(res, q[len(q)-1].Val)
        for i := 0; i< size ;i++ {
            pop := q[0]
            q = q[1:]
            if pop.Left != nil {
                q = append(q, pop.Left)
            }
            if pop.Right != nil {
                q = append(q, pop.Right)
            }
        }
    }
    return res
}
