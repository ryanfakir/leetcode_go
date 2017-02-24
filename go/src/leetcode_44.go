package main

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i, _ := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i < n; i++ {
		dp[0][i+1] = dp[0][i] && (p[i] == '*')
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[j] == '*' {
				dp[i+1][j+1] = dp[i][j+1] || dp[i+1][j]
			} else {
				dp[i+1][j+1] = dp[i][j] && (p[j] == s[i] || p[j] == '?')
			}
		}
	}
	return dp[m][n]
}
