func oddEvenList(head *ListNode) *ListNode {
    if head == nil {return nil}
    odd := head
    even := head.Next
    tail := even
    for odd != nil && even != nil {
        odd.Next = even.Next
        if odd.Next == nil {
            break
        }
        even.Next = odd.Next.Next
        odd = odd.Next
        even = even.Next
    }
    odd.Next = tail
    return head
}
