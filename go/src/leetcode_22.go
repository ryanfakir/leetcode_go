package src

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
        dfs(res, s + "(", left-1, right)
    }
    if right > 0 && left < right {
        dfs(res, s+")", left, right-1)
    }
}
