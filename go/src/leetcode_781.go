func numRabbits(answers []int) int {
    dict := make(map[int]int)
    var res int
    for _, v := range answers {
        if dict[v+1] == 0 {
            dict[v+1] = v
            res += v + 1
        } else {
            dict[v+1]--
        }
    }
    return res
}
