func reorderList(head *ListNode)  {
    if head == nil {return} 
    f, s := head, head
    for f != nil && f.Next != nil {
        f = f.Next.Next
        s = s.Next
    }
    sec := s.Next
    s.Next = nil
    var prev *ListNode
    for sec != nil {
        tmp := sec.Next
        sec.Next = prev
        prev = sec
        sec = tmp
    }
    for head != nil && prev != nil {
        save := head.Next
        head.Next = prev
        prev = prev.Next
        head.Next.Next = save
        head = save
    }
}
