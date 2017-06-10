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

func dfs(root *TreeNode, level *[]int, curlvl int, maxlvl int) {
	if root == nil || curlvl > maxlvl {
		return
	}
	if curlvl == maxlvl {
		*level = append(*level, root.Val)
		return
	}
	dfs(root.Left, level, curlvl+1, maxlvl)
	dfs(root.Right, level, curlvl+1, maxlvl)
}
func levelOrder(root *TreeNode) [][]int {
	return helper(root, 0, nil)
}

func helper(root *TreeNode, level int, res [][]int) [][]int {
	if root == nil {
		return res
	}
	if len(res) < level+1 {
		res = append(res, []int(nil))
	}
	res[level] = append(res[level], root.Val)
	if root.Left != nil {
		res = helper(root.Left, level+1, res)
	}
	if root.Right != nil {
		res = helper(root.Right, level+1, res)
	}
	return res
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		size := len(stack)
		var temp []int
		for i := 0; i < size; i++ {
			pop := stack[0]
			stack = stack[1:]
			temp = append(temp, pop.Val)
			if pop.Left != nil {
				stack = append(stack, pop.Left)
			}
			if pop.Right != nil {
				stack = append(stack, pop.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}
