var res [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
    res = nil
    dfs(0, len(graph), []int{0}, graph, make(map[int]bool))
    return res
}


func dfs(cur, size int, tmp []int, graph [][]int, visited map[int]bool) {
    if size -1 == cur {
        el := make([]int, len(tmp))
        copy(el, tmp)
        res = append(res, el)
        return
    }
    for _, v := range graph[cur] {
        if visited[v] {continue}
        visited[v] = true
        tmp = append(tmp, v)
        dfs(v, size, tmp, graph, visited)
        visited[v] = false
        tmp = tmp[:len(tmp)-1]
    }
    return 
}
