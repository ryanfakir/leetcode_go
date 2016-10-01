package src

 import "strconv"
func binaryTreePaths_recursive(root *TreeNode) []string {
    res := make([]string, 0)
    if root == nil {
        return res
    }
    if root.Left == nil && root.Right == nil {
        res = append(res, string(root.Val))
        return res
    }
    for _, el := range binaryTreePaths_recursive(root.Left) {
        el = strconv.Itoa(root.Val) + "->" + el
        res = append(res, el)
    }
    for _, el := range binaryTreePaths_recursive(root.Right) {
        el = strconv.Itoa(root.Val) + "->" + el
        res = append(res, el)
    }
    return res
}

func binaryTreePaths_BFS(root *TreeNode) []string {
    res := make([]string, 0)
    if root == nil {
        return res
    }
    q := make([]*TreeNode, 0)
    segment := make([]string, 0)
    q = append(q, root)
    segment = append(segment, strconv.Itoa(root.Val))
    for len(q) != 0 {
        treeNode := q[0]
        q = q[1:]
        seg := segment[0]
        segment = segment[1:]
        if treeNode.Left == nil && treeNode.Right == nil {
            res = append(res, seg)
        }
        if treeNode.Left != nil {
            q = append(q, treeNode.Left)
            segment = append(segment, seg + "->" + strconv.Itoa(treeNode.Left.Val))
        }
        if treeNode.Right != nil {
            q = append(q, treeNode.Right)
            segment = append(segment, seg + "->" + strconv.Itoa(treeNode.Right.Val))
        }
    }
    return res
}
