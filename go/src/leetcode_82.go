func deleteDuplicates(head *ListNode) *ListNode {
    dummy := &ListNode{0, head}
    cur:= head
    check := dummy
    for cur != nil && cur.Next != nil {
        if cur.Val == cur.Next.Val {
            for cur != nil && cur.Next != nil && cur.Val == cur.Next.Val {
                cur = cur.Next
            }
            cur = cur.Next
            check.Next = cur
        } else {
            cur = cur.Next
            check = check.Next
        }
    }
    return dummy.Next
}
