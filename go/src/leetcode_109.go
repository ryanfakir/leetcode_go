func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {return nil}
    if head.Next == nil {return &TreeNode{Val: head.Val}}
    s, f, prev := head, head, &ListNode{-1, head}
    for f != nil && f.Next != nil {
        fmt.Println(f)
        prev = prev.Next
        s = s.Next
        f = f.Next.Next
    }
    prev.Next = nil
    node := &TreeNode{Val: s.Val}
    node.Left = sortedListToBST(head)
    node.Right = sortedListToBST(s.Next)
    return node
}
