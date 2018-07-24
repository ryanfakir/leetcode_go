func removeDuplicateLetters(s string) string {
    dict := make(map[rune]int)
    for _, v := range s {
        dict[v]++
    }
    res := "0"
    visited := make(map[rune]bool)
    for _, v := range s {
        dict[v]--
        if visited[v] {continue}
        for res[len(res)-1] > byte(v) && dict[rune(res[len(res)-1])] > 0 {
            visited[rune(res[len(res)-1])] = false
            res = res[:len(res)-1]
        }
        res += string(v)
        visited[v] = true
    }
    return res[1:]
}
