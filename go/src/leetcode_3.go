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
