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

func generateParenthesis(n int) []string {
	return helper(nil, n, n, "")
}

func helper(res []string, left, right int, temp string) []string {
	if left > right {
		return res
	}
	if left == 0 && right == 0 {
		res = append(res, temp)
		return res
	}
	if left > 0 {
		res = helper(res, left-1, right, temp+"(")
	}
	if right > 0 {
		res = helper(res, left, right-1, temp+")")
	}
	return res
}

func generateParenthesis(n int) []string {
    return helper(n, n , nil, "")    
}

func helper(left, right int, res []string, tmp string) []string {
    if left < 0 || right < 0 {
        return res
    }
    if left == 0 && right == 0 {
        res = append(res, tmp)
        return res
    }    
    if right < left {return res}
    res = helper(left -1, right, res, tmp + "(")
    res = helper(left, right-1, res, tmp+ ")")    
    return res
    
}
