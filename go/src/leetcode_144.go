package src

func preorderTraversal(root *TreeNode) []int {
    temp := root
    stack := make([]*TreeNode, 0)
    res := make([]int, 0)
    for temp != nil || len(stack) != 0  {
        for temp != nil {
            stack = append(stack, temp)
            res = append(res, temp.Val)
            temp = temp.Left
        }
        pop:= stack[len(stack) -1]
        stack = stack[:len(stack)-1]
        temp = pop.Right
    }
    return res
}
