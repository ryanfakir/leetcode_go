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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{0, head}
    prev := dummy
    s, f := head, head
    for n > 0 && f != nil{
        f = f.Next
        n--
    }
    for f != nil {
        prev = prev.Next
        f = f.Next
        s = s.Next
    }
    prev.Next = s.Next
    return dummy.Next
}
