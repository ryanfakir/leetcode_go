package src
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for ;len(queue) != 0; {
		// offer
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			// push
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		temp := node.Left
		node.Left = node.Right
		node.Right = temp
	}
	return root
}