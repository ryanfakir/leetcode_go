package src

import (
	"strconv"
)

func diffWaysToCompute(input string) []int {
	res := make([]int, 0)
	for i, el := range input {
		if el == '+' || el == '-' || el == '*' {
			left := diffWaysToCompute(input[:i])
			right := diffWaysToCompute(input[i+1:])
			for _, l := range left {
				for _, r := range right {
					switch el {
					case '+':
						res = append(res, l+r)
						break
					case '-':
						res = append(res, l-r)
						break
					case '*':
						res = append(res, l*r)
						break
					}
				}
			}
		}
	}

	if len(res) == 0 {
		i, _ := strconv.Atoi(input)
		res = append(res, i)
	}

	return res
}
