func kthSmallestPrimeFraction(A []int, K int) []int {
    left, right := 0.0, 1.0
    for {
        mid := (left + right)/2
        var cnt int
        p, q := 0, 1
        for i, j := 0, 0; i < len(A); i++ {
            for j < len(A) && float64(A[i]) > mid * float64(A[j]) {j++}
            cnt += len(A) - j
            if j < len(A) && p* A[j] <  A[i] * q {
                p, q = A[i], A[j]
            }
        }
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
