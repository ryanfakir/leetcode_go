package main

import "math"

func minCut(s string) int {
	dp := make([]int, len(s))
	dp[0] = 0
	for i := 1; i < len(s); i++ {
		cut := 1 << 16
		for j := 0; j <= i; j++ {
			if valid(s, j, i) {
				if 0 == j {
					cut = 0
				} else {
					cut = int(math.Min(float64(cut), float64(dp[j-1]+1)))
				}
			}
		}
		dp[i] = cut
	}
	//fmt.Println(dp)
	return dp[len(s)-1]
}

func valid(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
