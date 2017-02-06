package main

import "strconv"

func calculate(s string) int {
	var q []int
	var d int
	var op byte = '+'
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			temp, _ := strconv.Atoi(string(s[i]))
			d = d*10 + temp
		}
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' || i == len(s)-1 {
			if op == '+' {
				q = append(q, d)
			}
			if op == '-' {
				q = append(q, -d)
			}
			if op == '*' {
				pop := q[len(q)-1]
				q = q[:len(q)-1]
				q = append(q, pop*d)
			}
			if op == '/' {
				pop := q[len(q)-1]
				q = q[:len(q)-1]
				q = append(q, pop/d)
			}
			op = s[i]
			d = 0
		}

	}
	var res int
	for _, v := range q {
		res += v
	}
	return res
}
