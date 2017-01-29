package main

func longestPalindrome(s string) int {
	m := make([]int, 256)
	for _, v := range s {
		m[v]++
	}
	var res int
	var single bool
	for _, v := range m {
		if v%2 == 0 {
			res += v
			continue
		}
		single = true
		if v > 2 {
			res += v - 1
		}
	}
	if single {
		return res + 1
	} else {
		return res
	}
}
