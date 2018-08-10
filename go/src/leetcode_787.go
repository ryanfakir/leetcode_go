func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
    dict := make(map[int][][]int)
    for _, f := range flights {
        dict[f[0]] = append(dict[f[0]], []int{f[1], f[2]})
    }
    var res = 1 << 63 -1
    var q = [][]int{[]int{src, 0}}
    for len(q) > 0 {
        size:= len(q)
        for i:= 0; i < size; i++ {
            pop := q[0]
            q = q[1:]
            if pop[0] == dst {res = int(math.Min(float64(res), float64(pop[1])))}
            for _, v := range dict[pop[0]] {
                if v[1] + pop[1] > res {continue}
                q = append(q, []int{v[0], v[1] + pop[1]})
            }
        } 
        if K < 0 {break}
        K--
    }
    if res == 1 << 63 -1 {return -1}
    return res
}

var res int 

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
    res = 1 << 63 -1
    dict := make(map[int][][]int)
    for _, f := range flights {
        dict[f[0]] = append(dict[f[0]], []int{f[1], f[2]})
    }
    visited := make(map[int]bool)
    visited[src] = true
    dfs(dict, src, dst, 0, K,visited)
    if res == 1 << 63 -1 {return -1}
    return res
}


func dfs(dict map[int][][]int, cur, dst, sum, k int, visited map[int]bool) {
    if sum > res {return}
    if cur == dst {
        res = int(math.Min(float64(res), float64(sum)))
        return
    }
    if k < 0 {return}
    for _, v := range dict[cur] {
        if !visited[v[0]] {
            visited[v[0]] = true
            dfs(dict, v[0], dst, sum+ v[1], k-1, visited)
        }
        visited[v[0]] = false
    }
}
