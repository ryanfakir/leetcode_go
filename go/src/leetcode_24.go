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

func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{0, head}
    prev := dummy
    first:= head
    for first != nil && first.Next != nil {
        sec := first.Next
        first.Next = sec.Next
        sec.Next = first
        prev.Next = sec
        prev = first
        first = first.Next
    }
    return dummy.Next
}
