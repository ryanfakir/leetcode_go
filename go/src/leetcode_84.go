package main

import "math"

func largestRectangleArea(heights []int) int {
	if heights == nil {
		return 0
	}
	col := len(heights)
	dp := make([][]int, col)
	for i, _ := range dp {
		dp[i] = make([]int, col)
	}
	var res int
	for i := 0; i < col; i++ {
		for j := i; j < col; j++ {
			if i == j {
				dp[i][j] = heights[i]
			} else {
				if dp[i][j-1] < heights[j] {
					dp[i][j] = dp[i][j-1]
				} else {
					dp[i][j] = heights[j]
				}
			}
			res = int(math.Max(float64(res), float64((j-i+1)*dp[i][j])))
		}
	}
	return res

}
