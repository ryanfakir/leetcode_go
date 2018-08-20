func numTilings(N int) int {
    if N == 0 {return 0}
    if N == 1 {return 1}
    mod := int(math.Pow(10, 9) + 7)
    dp := make([]int, N+1)
    dp[0], dp[1], dp[2] = 1, 1,  2
    for i:= 3; i<=N; i++ {
        dp[i] = (dp[i-1] * 2 + dp[i-3])%mod
    }
    return dp[N]
}
