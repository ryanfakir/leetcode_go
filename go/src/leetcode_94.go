package src

func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return make([]int, 0)
    }
    stack := make([]*TreeNode, 0)
    temp := root
    res := make([]int, 0)
    for len(stack) != 0 || temp != nil {
        for temp != nil {
            stack = append(stack, temp)
            temp = temp.Left
        }
        pop := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, pop.Val)
        temp = pop.Right
    }
    return res
}

func inorderTraversal(root *TreeNode) []int {
    if root == nil {return nil}
    var q []*TreeNode
    var res []int
    for len(q) >0 || root != nil {
        for root != nil {
            q = append(q, root)
            root = root.Left
        }
        pop := q[len(q)-1]
        q = q[:len(q)-1]
        res = append(res, pop.Val)
        root = pop.Right
    }
    return res
}

