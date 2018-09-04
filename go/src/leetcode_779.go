func kthGrammar(N int, K int) int {
    if N == 1 {return 0}
    if K <= 1 << uint(N-2) {
        return kthGrammar(N-1, K)
    }
    return kthGrammar(N-1, K - (1 << uint(N-2))) ^1
}
