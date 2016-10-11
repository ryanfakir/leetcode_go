package src

func kthSmallest(root *TreeNode, k int) int {
    stack := make([]*TreeNode, 0)
    temp := root;
    for len(stack) != 0 || temp != nil {
        for temp!= nil {
            stack= append(stack, temp)
            temp = temp.Left
        }
        pop := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        k--
        if k == 0 {
            return pop.Val
        }
        temp = pop.Right
    }
    return -1
}
