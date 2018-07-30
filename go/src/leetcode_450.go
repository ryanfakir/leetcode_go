package main

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil || root.Right == nil {
			if root.Left != nil {
				root = root.Left
			} else {
				root = root.Right
			}
		} else {
			temp := root.Right
			for temp.Left != nil {
				temp = temp.Left
			}
			root.Val = temp.Val
			root.Right = deleteNode(root.Right, root.Val)
		}
	}
	return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil { return nil }
    if root.Val > key {
        root.Left = deleteNode(root.Left, key)
    } else if root.Val < key {
        root.Right = deleteNode(root.Right, key)
    } else {
        if root.Left  == nil || root.Right == nil {
            if root.Left != nil {
                root = root.Left
            } else {
                root = root.Right
            }
        } else {
            tmp := root.Left
            var val int
            for tmp != nil {
                val = tmp.Val
                tmp = tmp.Right
            }
            root.Val = val
            root.Left = deleteNode(root.Left, val)
        }
    }
    return root
}
