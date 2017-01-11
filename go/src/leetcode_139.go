package main

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && contains(s[j:i], wordDict) {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func contains(target string, dict []string) bool {
	for _, v := range dict {
		if v == target {
			return true
		}
	}
	return false
}
