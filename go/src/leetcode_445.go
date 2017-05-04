package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1, s2 []int
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	var carry int
	var prev *ListNode = &ListNode{}
	i, j := len(s1)-1, len(s2)-1
	for i >= 0 || j >= 0 {
		var sum int
		if i >= 0 {
			sum += s1[i]
		}
		if j >= 0 {
			sum += s2[j]
		}
		sum += carry
		now := prev
		(*now).Val = sum % 10
		prev = &ListNode{}
		(*prev).Next = now
		carry = sum / 10
		i--
		j--
	}
	if carry == 1 {
		prev.Val = 1
		return prev
	} else {
		return prev.Next
	}
}
