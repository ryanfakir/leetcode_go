package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	row, col := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, row)
	for i, _ := range dp {
		dp[i] = make([]int, col)
	}
	if obstacleGrid[0][0] != 1 {
		dp[0][0] = 1
	}
	for i := 1; i < row; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = dp[i-1][0]
	}
	for i := 1; i < col; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = dp[0][i-1]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[row-1][col-1]
}
