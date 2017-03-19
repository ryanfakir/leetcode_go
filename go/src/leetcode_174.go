package main

import "math"

func calculateMinimumHP(dungeon [][]int) int {
	if dungeon == nil || len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 0
	}
	row, col := len(dungeon), len(dungeon[0])
	dp := make([][]int, row)
	for i, _ := range dp {
		dp[i] = make([]int, col)
	}
	dp[row-1][col-1] = int(math.Max(float64(1), float64(1-dungeon[row-1][col-1])))
	for i := row - 2; i >= 0; i-- {
		dp[i][col-1] = int(math.Max(float64(1), float64(0-dungeon[i][col-1]+dp[i+1][col-1])))
	}
	for i := col - 2; i >= 0; i-- {
		dp[row-1][i] = int(math.Max(float64(1), float64(0-dungeon[row-1][i]+dp[row-1][i+1])))
	}
	for i := row - 2; i >= 0; i-- {
		for j := col - 2; j >= 0; j-- {
			dp[i][j] = int(math.Max(float64(1), float64(0-dungeon[i][j]+int(math.Min(float64(dp[i][j+1]), float64(dp[i+1][j]))))))
		}
	}
	return dp[0][0]
}
