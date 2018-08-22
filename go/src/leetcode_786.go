func kthSmallestPrimeFraction(A []int, K int) []int {
    left, right := 0.0, 1.0
    p, q := 0, 1
    for {
        mid := (left + right)/2
        fmt.Println(mid)
        var cnt int
        for i, j := 0, 0; i < len(A); i++ {
            for j < len(A) && float64(A[i]) > mid * float64(A[j]) {j++}
            cnt += len(A) - j
            if j < len(A) && p* A[i] <  A[j] * q {
                p, q = A[i], A[j]
            }
        }
        fmt.Println(cnt)
        if cnt == K {
            return []int{p, q}
        } else if cnt < K {
            left = mid
        } else {
            right = mid
        }
    }
    return nil
}
