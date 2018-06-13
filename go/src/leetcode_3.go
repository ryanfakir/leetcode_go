package main

import "math"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := make([]int, 256)
	var i, j int
	res := 1
	for ; j < len(s); j++ {
		for m[s[j]] != 0 {
			m[s[i]] = 0
			i++
		}
		m[s[j]]++
		res = int(math.Max(float64(res), float64(j-i+1)))
	}
	return res
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var res int
	dict := make(map[byte]bool)
	var left, right int
	for right < len(s) {
		if !dict[s[right]] {
			dict[s[right]] = true
			res = int(math.Max(float64(res), float64(right-left+1)))
			right++
			continue
		}
		for s[left] != s[right] {
			dict[s[left]] = false
			left++
		}
		left++
		right++
	}
	return res
}
