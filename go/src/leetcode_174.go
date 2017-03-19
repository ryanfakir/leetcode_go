package main

import (
	"fmt"
	"math"
)

func calculateMinimumHP(dungeon [][]int) int {
	if dungeon == nil || len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 0
	}
	row, col := len(dungeon), len(dungeon[0])
	dp := make([][]int, row)
	for i, _ := range dp {
		dp[i] = make([]int, col)
	}

	if dungeon[0][0] < 0 {
		dp[0][0] = 1 - dungeon[0][0]
	} else {
		dp[0][0] = 1
	}
	for i := 1; i < row; i++ {
		if dungeon[i][0] > 0 {
			dp[i][0] = dp[i-1][0]
		} else {
			dp[i][0] = dp[i-1][0] - dungeon[i][0]
		}
	}
	for i := 1; i < col; i++ {
		if dungeon[0][i] > 0 {
			dp[0][i] = dp[0][i-1]
		} else {
			dp[0][i] = dp[0][i-1] - dungeon[0][i]
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if dungeon[i][j] > 0 {
				dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1])))
			} else {
				dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) - dungeon[i][j]
			}
		}
	}
	fmt.Println(dp)
	return dp[row-1][col-1]
}
