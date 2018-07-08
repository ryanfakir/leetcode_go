func postorderTraversal(root *TreeNode) []int {
    if root == nil {return nil}
    var res []int
    var stack = []*TreeNode{root}
    for len(stack) != 0 {
        pop := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append([]int{pop.Val}, res...)
        if pop.Left != nil {
            stack = append(stack, pop.Left)
        }
        if pop.Right != nil {
            stack = append(stack, pop.Right)
        }
    }
    
    return res
}

