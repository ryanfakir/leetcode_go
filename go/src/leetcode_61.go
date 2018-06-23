package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	oldHead := head
	var len int = 1
	for head.Next != nil {
		head = head.Next
		len++
	}
	k = k % len
	fir, sec := oldHead, oldHead
	for k > 0 {
		sec = sec.Next
		k--
	}
	for sec.Next != nil {
		fir = fir.Next
		sec = sec.Next
	}
	var newHead *ListNode
	if fir.Next != nil {
		newHead = fir.Next
	}
	fir.Next = nil

	if newHead == nil {
		return oldHead
	}
	sec.Next = oldHead
	return newHead
}

func rotateRight(head *ListNode, k int) *ListNode {
    var length int
    dummy := &ListNode{0, head}
    tmp := head
    for tmp != nil {
        tmp = tmp.Next
        length++
    }
    if length == 0 {return nil}    
    k = k % length
    if k == 0 {return head}
    s, f := dummy, dummy
    for k > 0 {
        f = f.Next
        k--
    }
    for f.Next != nil {
        f = f.Next
        s = s.Next
    }
    f.Next = head
    head = s.Next
    s.Next = nil
    return head
}
