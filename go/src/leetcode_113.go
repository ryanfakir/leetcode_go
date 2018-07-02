func pathSum(root *TreeNode, sum int) [][]int {
    return helper(root, nil, sum, nil)
}

func helper(root *TreeNode, tmp []int, sum int, res [][]int) [][]int {
    if root == nil {return res}
    if root.Val - sum == 0 &&  root.Left == nil && root.Right == nil {
            tmp = append(tmp, root.Val)
            nTmp := make([]int, len(tmp))
            copy(nTmp, tmp)
            res = append(res, nTmp)
            return res
    } else {
        tmp = append(tmp, root.Val)
        res = helper(root.Left, tmp, sum-root.Val, res)
        res = helper(root.Right, tmp, sum-root.Val, res)
        tmp = tmp[:len(tmp)-1]
        return res
    }
} 
