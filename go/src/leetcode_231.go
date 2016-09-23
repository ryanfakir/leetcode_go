package src

func isPowerOfTwo(n int) (res bool) {
    if n < 1 {
        return false
    }
    res = true
    for n > 1 {
        if n%2 != 0 {
            res = false
            break
        }
        n = n/2
    }
    return
}