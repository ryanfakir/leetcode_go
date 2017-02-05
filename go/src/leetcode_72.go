package main

import "math"

func minDistance(word1 string, word2 string) int {
	row, col := len(word1), len(word2)
	dp := make([][]int, row+1)
	for i, _ := range dp {
		dp[i] = make([]int, col+1)
	}
	for i := 0; i <= row; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= col; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = int(math.Min(float64(dp[i-1][j-1]), math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + 1)
			}
		}
	}
	return dp[row][col]
}
