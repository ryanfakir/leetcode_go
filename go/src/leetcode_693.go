func hasAlternatingBits(n int) bool {
    n ^= (n / 4)
    return n & (n-1) == 0 
}
