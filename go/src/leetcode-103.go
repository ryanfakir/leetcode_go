func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {return nil}
    var q []*TreeNode = []*TreeNode{root}
    var n int
    var res [][]int
    for len(q) > 0 {
        size := len(q)
        var tmp []int
        var tmpQ []*TreeNode
        for i := 0; i < size; i++ {
            var j int
            j  = size -1 -i
            pop := q[j]
            tmp = append(tmp, pop.Val)
            if n %2 == 0 {
                //0
                if pop.Left != nil {
                    tmpQ = append(tmpQ, pop.Left)
                }
                if pop.Right != nil {
                    tmpQ = append(tmpQ, pop.Right)
                }
            } else {
                //1
                 if pop.Right != nil {
                    tmpQ = append(tmpQ, pop.Right)
                }
                if pop.Left != nil {
                    tmpQ = append(tmpQ, pop.Left)
                }
            }
        }
        q = tmpQ
        res = append(res, tmp)
        n++
    }
    return res
}
