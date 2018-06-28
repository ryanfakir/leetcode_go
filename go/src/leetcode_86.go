func partition(head *ListNode, x int) *ListNode {
    var left, right  = &ListNode{-1, nil}, &ListNode{-1, nil}
    pl, pr := left, right
    for head != nil {
        if head.Val >= x {
            right.Next = &ListNode{head.Val, nil}
            right = right.Next
        } else {
            left.Next = &ListNode{head.Val, nil}
            left = left.Next
        }
        head = head.Next
    }
    if left.Val == -1 {
        return pr.Next
    } else {
        left.Next = pr.Next
        return pl.Next
    }
}
