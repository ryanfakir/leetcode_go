package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	tmp := res
	var carry int
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			l1 = &ListNode{Val: 0}
		}
		if l2 == nil {
			l2 = &ListNode{Val: 0}
		}
		tmp.Next = &ListNode{}
		tmp.Next.Val = (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10
		tmp = tmp.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	return res.Next
}
