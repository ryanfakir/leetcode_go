package src

func levelOrderBottom(root *TreeNode) [][]int {
    res := make([][]int, 0)
    if root == nil {
        return res
    }
    q := make([]*TreeNode, 0)
    q = append(q, root)
    for len(q) != 0 {
        size := len(q)
        level := make([]int, 0)
        for i := 0; i < size; i++ {
            x := q[0]
            q = q[1:]
            level = append(level, x.Val)
            if x.Left != nil {
                q = append(q, x.Left)
            }
            if x.Right != nil {
                q = append(q, x.Right)
            }
        }
        res = append(res, level)
    }
    for left, right:= 0, len(res)-1; left < right; left,right = left+1, right-1 {
        res[left], res[right] = res[right], res[left]
    }
    return res
}