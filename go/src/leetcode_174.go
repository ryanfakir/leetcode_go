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

func calculateMinimumHP(dungeon [][]int) int {
    m, n := len(dungeon), len(dungeon[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    dp[m-1][n-1] = int(math.Max(1.0, float64(1-dungeon[m-1][n-1])))
    for i := m-2; i >= 0 ; i-- {
        dp[i][n-1] = int(math.Max(1.0, float64(dp[i+1][n-1] - dungeon[i][n-1])))
    }
    for i:= n-2; i>=0; i-- {
        dp[m-1][i] = int(math.Max(1.0, float64(dp[m-1][i+1] - dungeon[m-1][i])))
    }
    for i := m-2; i >=0; i-- {
        for j:= n-2; j>=0; j-- {
            dp[i][j] = int(math.Max(1.0, math.Min(float64(dp[i+1][j]), float64(dp[i][j+1])) - float64(dungeon[i][j])))
        }
    }
    return dp[0][0]
}
