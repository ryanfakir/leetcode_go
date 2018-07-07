package src

func preorderTraversal(root *TreeNode) []int {
	temp := root
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	for temp != nil || len(stack) != 0 {
		for temp != nil {
			stack = append(stack, temp)
			res = append(res, temp.Val)
			temp = temp.Left
		}
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		temp = pop.Right
	}
	return res
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}


func preorderTraversal(root *TreeNode) []int {
    var q []*TreeNode
    var res []int
    for len(q) > 0 || root != nil {
        for root != nil {
            res = append(res, root.Val)
            q = append(q, root)
            root = root.Left
        }
        pop := q[len(q)-1]
        q = q[:len(q)-1]
        root = pop.Right
    }
    return res
}
