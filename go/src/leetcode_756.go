func pyramidTransition(bottom string, allowed []string) bool {
    dict := make(map[string][]string)
    for _, a := range allowed {
        dict[a[:2]] = append(dict[a[:2]], a[2:])
    }
    return dfs(bottom, "", 0, dict)
}

func dfs(b, tmp string, start int,  dict map[string][]string) bool {
    if len(b) == 1 {return true}
    if start == len(b)-1 {return dfs(tmp, "", 0, dict)}
    el := b[start: start+ 2]
    for _, g := range dict[el] {
        if dfs(b, tmp + g, start + 1, dict) {
            return true
        }
    }
    return false
}
