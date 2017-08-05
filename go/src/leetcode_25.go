package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{-1, head}
	cur := head
	pre := dummy
	var i int
	for cur != nil {
		i++
		if i%k == 0 {
			pre = reverser(pre, cur.Next)
			cur = pre.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

func reverser(pre *ListNode, final *ListNode) *ListNode {
	fir := pre.Next
	sec := fir.Next
	for sec != final {
		fir.Next = sec.Next
		sec.Next = pre.Next
		pre.Next = sec
		sec = fir.Next
	}
	return fir
}
