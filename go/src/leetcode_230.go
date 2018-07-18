package src

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	temp := root
	for len(stack) != 0 || temp != nil {
		for temp != nil {
			stack = append(stack, temp)
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

var count int

func kthSmallest(root *TreeNode, k int) int {
	count = k
	return helper(root)
}

func helper(root *TreeNode) int {
	if root == nil {
		return -1
	}
	val := helper(root.Left)
	if count == 0 {
		return val
	}
	count--
	if count == 0 {
		return root.Val
	}
	return helper(root.Right)
}
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if k == 1 {
			return root.Val
		}
		k--
		root = root.Right
	}
	return -1
}

func kthSmallest(root *TreeNode, k int) int {
    var q []*TreeNode
    var cnt int
    for len(q) > 0 || root != nil {
        for root != nil {
            q = append(q, root)
            root = root.Left
        }
        pop := q[len(q)-1]
        q  = q[:len(q)-1]
        cnt++
        if cnt == k {return pop.Val}
        root = pop.Right
    }
    return 0
    
}
