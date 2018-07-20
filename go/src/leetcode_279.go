func numSquares(n int) int {
    dp := make([]int, n+1)
    for i := range dp {
        dp[i] = 1<<7-1
    }
    dp[0] = 0
    for i:= 0 ; i<= n; i++ {
        for j:= 1; j * j + i<=n; j++ {
            dp[j * j + i] = int(math.Min(float64(dp[j*j+ i]), float64(dp[i] + 1)))
        } 
    }
    return dp[n]
}
