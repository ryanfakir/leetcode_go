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

func pyramidTransition(bottom string, allowed []string) bool {
    dict := make(map[string][]string)
    for _ , a :=range allowed {
        dict[a[:2]] = append(dict[a[:2]], a[2:])
    }
    return dfs(bottom, "", dict)
}

func dfs(low, up string, dict map[string][]string) bool {
    if len(up) == 1 && len(low) == 2{return true}
    if len(up) == len(low) -1 {return dfs(up, "", dict)}
    for _, v := range dict[low[len(up): len(up)+2]] {
        if dfs(low, up + v, dict) {return true}  
    }
    return false
}
