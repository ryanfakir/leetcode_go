package main

import "strconv"

func decodeString(s string) string {
	return helper(s, 0, len(s)-1)
}

func helper(s string, i, j int) string {
	var res string
	var cnt, l int
	for left := i; left <= j; left++ {
		if s[left] >= 'a' && s[left] <= 'z' {
			res += string(s[left])
		}
		if s[left] >= '0' && s[left] <= '9' {
			val, _ := strconv.Atoi(string(s[left]))
			cnt = cnt*10 + val
		}
		if s[left] == '[' {
			l++
			x := left + 1
			for l > 0 {
				left++
				if s[left] == '[' {
					l++
				}
				if s[left] == ']' {
					l--
				}
			}
			y := left - 1
			temp := helper(s, x, y)
			for cnt > 0 {
				res += temp
				cnt--
			}
		}
	}
	return res
}
