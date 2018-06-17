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

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {return nil}
    if len(lists) == 1 {return lists[0]}
    
    
    for len(lists) != 1 {
        
        var newList []*ListNode
        for len(lists) >= 2 {
            tmp := merge(lists[0], lists[1])
            newList = append(newList, tmp)
            lists =lists[2:]
        }
        if len(lists) == 1 {
            newList = append(newList, lists[0])
        }
        lists = newList
    }
    return lists[0]
}


func merge(l1, l2 *ListNode) *ListNode {
    var tmp *ListNode = &ListNode{}
    var res = tmp
    for l1 != nil && l2 != nil {
        if l1.Val >= l2.Val {
            res.Next = &ListNode{l2.Val, nil}
            l2 = l2.Next
        } else {
            res.Next = &ListNode{l1.Val, nil}
            l1 = l1.Next
        }
        res = res.Next
    }
    for l1 != nil {
        res.Next =  &ListNode{l1.Val, nil}
        res = res.Next
        l1 = l1.Next
    }
    for l2 != nil {
        res.Next =  &ListNode{l2.Val, nil}
        res = res.Next
        l2 = l2.Next
    }
    fmt.Println(tmp.Next)
    return tmp.Next
    
}
