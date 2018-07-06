func longestConsecutive(nums []int) int {
    var res int
    dict := make(map[int]int)
    for _, v := range nums {
        if dict[v] == 0 {
            left, right := dict[v-1], dict[v+1]
            sum := left + right + 1
            dict[v]= sum
            res = int(math.Max(float64(res), float64(sum)))
            dict[v-left] = sum
            dict[v+ right] = sum
        }
    }
    return res
}
