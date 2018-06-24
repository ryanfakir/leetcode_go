package src

func uniquePaths(m int, n int) int {
    dp := make([][]int, m)
    for i, _ := range dp {
        col := make([]int, n)
        dp[i] = col
    }
    for i := 0; i<m; i++ {
        dp[i][0] = 1
    }

    for i :=0; i<n; i++ {
        dp[0][i] = 1
    }

    for i:=1; i<m; i++ {
        for j:=1; j<n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[m-1][n-1]
}


func uniquePaths(m int, n int) int {
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, m)
    }
    dp[0][0] = 1
    for i:= 0; i < n; i++ {
        dp[i][0] = 1
    }
    for j:= 0; j < m; j++ {
        dp[0][j] = 1
    }
    for i:= 1; i < n; i++ {
        for j := 1; j < m; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[n-1][m-1]
}
