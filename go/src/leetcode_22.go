package src

import "bytes"

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	dfs(&res, "", n, n)
	return res
}

func dfs(res *[]string, s string, left int, right int) {
	if left == 0 && right == 0 {
		*res = append(*res, s)
		return
	}
	if left > 0 {
		dfs(res, s+"(", left-1, right)
	}
	if right > 0 && left < right {
		dfs(res, s+")", left, right-1)
	}
}

func generateParenthesis(n int) []string {
	var b bytes.Buffer
	return helper(nil, n, n, b)
}

func helper(res []string, left int, right int, temp bytes.Buffer) []string {
	if left == 0 && right == 0 {
		res = append(res, temp.String())
	}
	if left > 0 {
		temp.WriteString("(")
		l := temp.Len()
		res = helper(res, left-1, right, temp)
		temp.Truncate(l - 1)
	}
	if right > left {
		temp.WriteString(")")
		l := temp.Len()
		res = helper(res, left, right-1, temp)
		temp.Truncate(l - 1)
	}
	return res
}
