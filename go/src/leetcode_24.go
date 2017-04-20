package src

func swapPairs(head *ListNode) *ListNode {
	placeHolder := &ListNode{0, nil}
	placeHolder.Next = head
	temp := placeHolder
	for temp != nil && temp.Next != nil && temp.Next.Next != nil {
		fir := temp.Next
		sec := temp.Next.Next.Next
		temp.Next = temp.Next.Next
		temp.Next.Next = fir
		temp.Next.Next.Next = sec
		temp = fir
	}
	return placeHolder.Next

}
