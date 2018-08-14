func accountsMerge(accounts [][]string) [][]string {
    dict := make(map[string][]int)
    for i, a := range accounts {
        for index, v := range a {
            if index == 0 {continue}
            dict[v] = append(dict[v], i)
        }
    }
    visited := make(map[int]bool)
    var res [][]string
    for i:= 0; i < len(accounts); i++ {
        if visited[i] { continue}
        var q = []int{i}
        set := make(map[string]bool)
        visited[i] = true
        for len(q) > 0 {
            pop := q[0]
            q = q[1:]
            visited[pop] = true
            for i, v := range accounts[pop] {
                if i == 0 {continue}
                set[v] = true
            }
            for k := range set {
                for _, user := range dict[k] {
                    if visited[user] {continue}
                    visited[user] = true
                    q = append(q, user)
                }
            }
        }
        var tmp []string
        for k := range set {
            tmp = append(tmp, k)
        }
        sort.Strings(tmp)
        el := append([]string{accounts[i][0]}, tmp...)
        res = append(res, el)
    }
    return res
}
