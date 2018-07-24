func maxProduct(words []string) int {
    dict := make([]int, len(words))
    var res int
    for i := 0; i< len(words); i++ {
        for _, v := range words[i] {
            dict[i] |= 1 << uint(v- 'a')
        }
        for j := 0; j < i; j++ {
            if (dict[i] & dict[j]) == 0 {
                res = int(math.Max(float64(res), float64(len(words[i]) * len(words[j]))))  
            }
        }
    }
    return res
}
