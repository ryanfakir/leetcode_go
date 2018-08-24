func alienOrder(words []string) string {
    graph := make(map[byte][]byte)
    ind := make(map[byte]int)
    for _, word := range words {
        for _, w := range word {
            ind[byte(w)] = 0
        }
    }
    for i:= 1; i < len(words); i++ {
        mn := int(math.Min(float64(len(words[i])), float64(len(words[i-1]))))
        for j:=0 ; j < mn; j++ {
            if words[i][j] != words[i-1][j] {
                graph[words[i-1][j]] = append(graph[words[i-1][j]], words[i][j])
                ind[words[i][j]]++
                break
            }
        }
    }
    var q []byte
    var res string
    for k, v := range ind {
        if v == 0 {q = append(q, k)}
    }
    for len(q) > 0 {
        pop := q[0]
        q = q[1:]
        res += string(pop)
        for _, arr := range graph[pop] {
            ind[arr]--
            if ind[arr] == 0 {q = append(q, arr)}
        }
    }
    if len(res) != len(ind) {return ""}
    return res
}
