func maximalSquare(matrix [][]byte) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {return 0}
    m, n := len(matrix), len(matrix[0])
    dp := make([][]int, m)
    for i:= range dp {
        dp[i] = make([]int, n)
    }
    var max int
    for i:= 0; i< m; i++ {
        for j:=0; j < n ; j++ {
            if i == 0 || j ==0 {dp[i][j] = int(matrix[i][j] - matrix[0][0])} else {
                dp[i][j] = int(matrix[i][j] - '0') * ( int(math.Min(float64(dp[i-1][j-1]), math.Min(float64(dp[i-1][j]), float64(dp[i][j-1])))) + 1)
            }
            max = int(math.Max(float64(max), float64(dp[i][j])))
        }
    }
    return max * max
}

