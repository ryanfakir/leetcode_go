package src

func minDepth(root *TreeNode) int {
    q := make([]*TreeNode, 0)
    if root == nil {
        return 0
    }
    q = append(q, root)
    lvl := 1
    for len(q) != 0 {
        size := len(q)
        for i :=0; i<size;i++ {
            el := q[0]
            q = q[1:]
            if el.Left == nil && el.Right == nil {
                return lvl
            }
            if el.Left != nil {
                q = append(q, el.Left)
            }
            if el.Right != nil {
                q = append(q, el.Right)
            }
        }
        lvl++;
    }
    return lvl
}
