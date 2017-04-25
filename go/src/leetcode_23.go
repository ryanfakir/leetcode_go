package main

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return part(lists, 0, len(lists))
}

func part(list []*ListNode, start, end int) *ListNode {
	if start+1 > len(list) {
		return nil
	}
	if start == end {
		return list[start]
	}
	mid := (end-start)/2 + start
	return merge(part(list, start, mid), part(list, mid+1, end))
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	
	if l2.Val < l1.Val {
		l2.Next = merge(l1, l2.Next)
		return l2
	} else {
		l1.Next = merge(l1.Next, l2)
		return l1
	}
}
