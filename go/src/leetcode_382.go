package src

import "math/rand"

type Solution struct {
    head *ListNode
}


/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
    s := Solution{head}
    s.head = head
    return s
}


/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
    dummy := this.head
    len := 1
    res := dummy.Val
    for dummy != nil {
        if rand.Intn(len) == 0 {
            res = dummy.Val
        }
        len++
        dummy = dummy.Next
    }
    return res
}
