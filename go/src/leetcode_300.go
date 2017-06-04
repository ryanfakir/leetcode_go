package main

import "math"

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i, _ := range dp {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		max := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				max = int(math.Max(float64(max), float64(dp[j]+1)))
			}
		}
		dp[i] = max
	}
	var res int
	for _, v := range dp {
		res = int(math.Max(float64(res), float64(v)))
	}
	return res
}
