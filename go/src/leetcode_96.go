package src

func numTrees(n int) int {
    res := make([]int, n+1)
    res[0], res[1] = 1, 1
    for i := 2; i<=n; i++ {
        for j:= 0; j<i; j++ {
            res[i] += res[j]*res[i-j-1]
        }
    }
    return res[n]
}


func numTrees(n int) int {
    dp := make([]int, n+1)
    dp[0] = 1
    for i:= 1; i<=n; i++ {
        for j:=0; j< i; j++ {
            dp[i] += dp[j] * dp[i-j-1]
        }
    }
    return dp[n]
}
