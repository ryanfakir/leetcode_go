package src

 import "strconv"
func binaryTreePaths(root *TreeNode) []string {
    res := make([]string, 0)
    if root == nil {
        return res
    }
    if root.Left == nil && root.Right == nil {
        res = append(res, string(root.Val))
        return res
    }
    for _, el := range binaryTreePaths(root.Left) {
        el = strconv.Itoa(root.Val) + "->" + el
        res = append(res, el)
    }
    for _, el := range binaryTreePaths(root.Right) {
        el = strconv.Itoa(root.Val) + "->" + el
        res = append(res, el)
    }
    return res
}
