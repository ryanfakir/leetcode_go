func numSubarrayBoundedMax(A []int, L int, R int) int {
    var res int
    for i := 0 ; i < len(A); i++ {
        if A[i] > R {continue}
        var curMax = -1 << 15
        for j:= i ; j < len(A); j++ {
            curMax = int(math.Max(float64(curMax), float64(A[j])))
            if curMax > R {break}
            if curMax >= L {res++}
        }
    }
    return res
}
