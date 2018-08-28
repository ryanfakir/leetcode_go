func isBipartite(graph [][]int) bool {
    color := make([]int, len(graph))
    for i := 0; i< len(graph); i++ {
        if color[i] != 0 {continue}
        var q = []int{i}
        color[i] = 1
        for len(q) > 0 {
            pop := q[0]
            q = q[1:]
            for _, v := range graph[pop] {
                if color[v] == 0 {
                    color[v] = -color[pop]
                    q = append(q, v)
                } else {
                    if color[v] != -color[pop] {return false} 
                }
            }
        }
    }
    return true
}


func isBipartite(graph [][]int) bool {
    color := make([]int, len(graph))
    for i:= 0; i < len(graph); i++ {
        if color[i] == 0 && !dfs(graph, 1, i, color) {
            return false
        }
    }
    return true
}

func dfs(graph [][]int, color, cur int, arr []int) bool {
    if arr[cur] != 0 { return arr[cur] == color}
    if arr[cur] == 0 { arr[cur] = color}
    for _ , v := range graph[cur] {
        if !dfs(graph, -color, v, arr) {return false}
    }
    return true
}
