package src

func hasPathSum(root *TreeNode, sum int) bool {
    if (root == nil) {
        return false
    }
    q := make([]*TreeNode, 0)
    value := make([]int, 0)
    q = append(q, root)
    value = append(value, root.Val)
    for len(q) != 0 {
        treeNode := q[0]
        q = q[1:]
        total := value[0]
        value = value[1:]
        if treeNode.Left == nil && treeNode.Right == nil {
            if sum == total {
                return true
            }
        }
        if treeNode.Left != nil {
            q = append(q, treeNode.Left)
            value = append(value, total + treeNode.Left.Val)
        }
        if treeNode.Right != nil {
            q = append(q, treeNode.Right)
            value = append(value, total+ treeNode.Right.Val)
        }
    }
    return false
}