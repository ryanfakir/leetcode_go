package src

func levelOrder(root *TreeNode) [][]int {
    maxlvl := 0
    res := make([][]int, 0)
    if root == nil {
        return res
    }
    for {
        level := make([]int, 0)
        dfs(root, &level, 0, maxlvl)
        if len(level) == 0 {
            break
        }
        res = append(res, level)
        maxlvl++
    }
    return res
}

func dfs(root *TreeNode, level *[]int,curlvl int, maxlvl int) {
    if root == nil || curlvl > maxlvl{
        return
    }
    if curlvl == maxlvl {
        *level = append(*level, root.Val)
        return
    }
    dfs(root.Left, level, curlvl+1, maxlvl)
    dfs(root.Right, level, curlvl+1, maxlvl)
}