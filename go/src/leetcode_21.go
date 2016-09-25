package src

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    placeHolder := new(ListNode)
    placeHolder.Val = 0
    temp := placeHolder
    for l1 != nil && l2 != nil {
        if l1.Val >= l2.Val {
            temp.Next = &ListNode{l2.Val, nil}
            l2 = l2.Next
        } else {
            temp.Next = &ListNode{l1.Val, nil}
            l1 = l1.Next
        }
        temp = temp.Next
    }
    for l1 != nil {
        temp.Next = &ListNode{l1.Val, nil}
        l1 = l1.Next
        temp = temp.Next
    }
    for l2 != nil {
        temp.Next = &ListNode{l2.Val, nil}
        l2 = l2.Next
        temp = temp.Next
    }
    return placeHolder.Next
}
