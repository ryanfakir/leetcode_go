package src

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fir := head
	for fir != nil {
		sec := fir.Next
		for sec != nil {
			if sec.Val == fir.Val {
				fir.Next = sec.Next
				sec = sec.Next
			} else {
				break
			}
		}
		fir = sec
	}
	return head
}
