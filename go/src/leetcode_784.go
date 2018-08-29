var res []string

func letterCasePermutation(S string) []string {
    res = nil
    dfs(0, S, "")
    return res
}

func dfs(cur int, s, tmp string) {
    if len(s) == cur {
        res = append(res, tmp)
        return 
    } 
    i := cur
    if s[i] >= '0' && s[i] <= '9' {
        tmp += string(s[i])
        dfs(i+1, s, tmp)
    } else {
        ltmp := tmp + strings.ToLower(string(s[i])) 
        dfs(i+1, s, ltmp)
        utmp := tmp + strings.ToUpper(string(s[i]))
        dfs(i+1, s, utmp)
    }
        
}
