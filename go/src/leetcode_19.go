package src

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    zero := &ListNode{0, head}
    temp := zero
    tail := zero
    count := 0
    for tail != nil {
        tail = tail.Next
        if count > n {
            temp = temp.Next
        }
        count++
    }
    temp.Next = temp.Next.Next
    return zero.Next
}
