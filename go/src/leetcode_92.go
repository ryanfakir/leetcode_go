package main

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}
	pre := &ListNode{}
	pre.Next = head
	start := pre
	for i := m; i > 1; i-- {
		pre = pre.Next
	}
	fir := pre.Next
	sec := fir.Next
	for i := n - m; i > 0; i-- {
		fir.Next = sec.Next
		sec.Next = pre.Next
		pre.Next = sec

		sec = fir.Next
	}
	return start.Next
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
    dummy := &ListNode{-1, head}
    prev := dummy
    for i := 1 ; i < m; i++ {
        prev = prev.Next
    }
    first, sec := prev.Next, prev.Next.Next
    for sec != nil && m < n {
        first.Next = sec.Next
        sec.Next = prev.Next
        prev.Next = sec
        sec = first.Next
        m++
    }
    return dummy.Next
}
