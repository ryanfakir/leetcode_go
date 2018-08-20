func numMatchingSubseq(S string, words []string) int {
    dict := make(map[rune][]int)
    for i, v := range S {
        dict[v] = append(dict[v], i)
    }
    var res int
    for _, word := range words {
        var cur int = -1
        var add = true
        for _, w := range word {
            local := cur
            for _, v := range dict[w] {
                if v > local {
                    local = v
                    break
                } 
            }
            if local == cur {
                add =false
                break
            }
            cur = local
        }
        if add {
            res++
        }
    }
    return res
}
