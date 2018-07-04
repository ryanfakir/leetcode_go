func minimumTotal(triangle [][]int) int {
    t := triangle
    if len(t) == 0 || len(t[0]) == 0 {return 0}
    m := len(t)
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, i+1)
    }
    
    dp[0][0] = t[0][0]
    for i := 1; i < m; i++ {
        dp[i][0] += dp[i-1][0] + t[i][0]
        dp[i][i] += dp[i-1][i-1] + t[i][i]
    }
    for i:=2; i< m; i++ {
        for j:= 1; j < i; j++ {
            dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i-1][j-1]))) + t[i][j]
        }
    }
    var res = 1 << 32
    for j:=0; j< m; j++ {
        res = int(math.Min(float64(res), float64(dp[m-1][j])))
    }
    return res
}
