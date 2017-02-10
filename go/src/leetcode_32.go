package main

import "math"

func longestValidParentheses(s string) int {
	var q []int
	var res, start int
	for i, v := range s {
		if v == '(' {
			q = append(q, i)
		}
		if v == ')' {
			if len(q) > 0 {
				q = q[:len(q)-1]
				if len(q) == 0 {
					res = int(math.Max(float64(i-start+1), float64(res)))
				} else {
					res = int(math.Max(float64(i-q[len(q)-1]), float64(res)))
				}
			} else {
				start = i + 1
			}
		}
	}
	return res
}
