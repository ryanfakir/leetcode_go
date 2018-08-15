var dict map[string]int
var res []*TreeNode
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    dict = make(map[string]int)
    res = nil 
    dfs(root)
    return res
}

func dfs(root *TreeNode) string { 
    if root == nil {return "*"}
    tmp := strconv.Itoa(root.Val) + "," + dfs(root.Left) + "," + dfs(root.Right)
    dict[tmp]++
    if dict[tmp] == 2 {
        res = append(res, root)
    }
    return tmp
}
