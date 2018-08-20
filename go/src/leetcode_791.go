func customSortString(S string, T string) string {
    dict := make(map[rune]int)
    for _, v := range T {
        dict[v]++
    }
    var res string
    for _, v := range S {
        if dict[v] != 0 {
            for i:= 0; i < dict[v]; i++ {
                res += string(v)
            }
            dict[v] = 0
        }
    }
    for k, v := range dict {
        for i:= 0; i < v; i++ {
            res += string(k)
        }
    }
    return res
}
