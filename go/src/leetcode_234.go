package src

func isPalindrome(head *ListNode) bool {
    if head == nil {
        return true
    }
    mid := findMid(head)
    tail := reverse(mid.Next)
    for tail != nil{
        if head.Val != tail.Val {
            return false
        }
        head = head.Next
        tail = tail.Next
    }
    return true
}

func findMid(head *ListNode) *ListNode {
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}

func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        temp := head.Next
        head.Next = prev
        prev = head
        head = temp
    }
    return prev
}
