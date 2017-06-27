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
