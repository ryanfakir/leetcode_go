func nthSuperUglyNumber(n int, primes []int) int {
    p  := primes
    pos := make([]int, len(p))
    res := []int{1}
    for i:= 1; i < n; i++ {
        var next = 1 << 31 -1 
        for j:= 0; j < len(pos); j++ {
            next = int(math.Min(float64(next), float64(res[pos[j]]* p[j])))
        }
        for j:= 0; j < len(pos); j++ {
            if next == res[pos[j]] * p[j] {
                pos[j]++
            }
        }
        res = append(res, next)
    }
    return res[n-1]
}
