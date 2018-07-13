func rangeBitwiseAnd(m int, n int) int {
    var level uint
    for m != n {
        m >>= 1
        n >>= 1
        level++
    }
    return m<<level
}
