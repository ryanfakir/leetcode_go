package main

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i, _ := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i < n; i++ {
		dp[0][i+1] = i > 0 && p[i] == '*' && dp[0][i-1]
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[j] != '*' {
				dp[i+1][j+1] = dp[i][j] && (p[j] == '.' || p[j] == s[i])
			} else {
				dp[i+1][j+1] = (j > 0 && (p[j-1] == s[i] || p[j-1] == '.') && dp[i][j+1]) || dp[i+1][j] || (j > 0 && dp[i+1][j-1])
			}
		}
	}
	return dp[m][n]
}
