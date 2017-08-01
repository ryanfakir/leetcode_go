package main

func longestPalindrome(s string) string {
	var res string
	for i, _ := range s {
		left, right := i, i
		odd := helper(left, right, s)
		even := helper(left, right+1, s)
		if len(odd) >= len(res) {
			res = odd
		}
		if len(even) >= len(res) {
			res = even
		}
	}
	return res
}

func helper(left int, right int, s string) (res string) {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			if len(s[left:right+1]) >= len(res) {
				res = s[left : right+1]
			}
			left--
			right++
			continue
		}
	}
	return
}
