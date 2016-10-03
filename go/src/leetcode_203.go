package src

func removeElements(head *ListNode, val int) *ListNode {
    temp:= new(ListNode)
    temp.Next = head
    front := temp
    for temp.Next != nil {
        if temp.Next.Val == val {
            temp.Next = temp.Next.Next;
        } else {
            temp = temp.Next
        }
    }
    return front.Next
}
