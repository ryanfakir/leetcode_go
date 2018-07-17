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
func calculate(s string) int {
	var stack []int
	var val int
	var op = '+'
	for i, v := range s {
		if v >= '0' && v <= '9' {
			digit, _ := strconv.Atoi(string(v))
			val = val*10 + digit
		}
		if v == '*' || v == '+' || v == '-' || v == '/' || i == len(s)-1 {
			if op == '+' {
				stack = append(stack, val)
			}
			if op == '-' {
				stack = append(stack, -val)
			}
			if op == '*' {
				temp := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, temp*val)
			}
			if op == '/' {
				temp := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, temp/val)
			}
			op = v
			val = 0
		}
	}
	var res int
	for _, v := range stack {
		res += v
	}
	return res
}

func calculate(s string) int {
    var stack []int
    var op = '+'
    var num int
    for i , v := range s {
        if v >= '0' && v <= '9' {
            num = num * 10 + int(v - '0')
        if v == '+' || v == '-' || v == '*' || v == '/' || i == len(s)-1 {
            fmt.Println(num)
            if op == '+' {stack = append(stack, num)}
            if op == '-' {stack = append(stack, -num)}
            if op == '*' {
                 pop := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                stack = append(stack, pop * num)
            }
            if op == '/' {
                 pop := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                stack = append(stack, pop/ num)
            }
            op = v
            num = 0
        }
    }
    var res int
    for _ ,v  := range stack {
        res  += v
    }
    return res
}
