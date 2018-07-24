func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i := range dp {
        dp[i] = 1 << 31 -1
    }
    dp[0] = 0
    for i:= 1 ; i <= amount; i++ {
        for j:= 0; j < len(coins); j++ {
            if i - coins[j] >= 0  && dp[i- coins[j]] != 1 << 31-1 {
                dp[i] = int(math.Min(float64(dp[i]), float64(1 + dp[i- coins[j]])))
            }
        }
    }
    if dp[amount] == 1 << 31 -1 {
        return -1
    } else {
        return dp[amount]
    }
}
