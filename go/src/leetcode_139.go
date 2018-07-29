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

func wordBreak(s string, wordDict []string) bool {
	l := len(s)
	if l == 0 {
		return true
	}
	for i := 0; i < l; i++ {
		if contains(wordDict, s[:i+1]) {
			if wordBreak(s[i+1:], wordDict) {
				return true
			}
		}
	}
	return false
}

func contains(wordDict []string, target string) bool {
	for _, v := range wordDict {
		if v == target {
			return true
		}
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
    if len(s) == 0 {return true}
    dict := make(map[string]bool)
    for _ , word:= range wordDict {
        dict[word] = true
    }
    dp := make([]bool, len(s))
    for i:= 0; i < len(dp); i++ {
        for j := -1; j < i; j++ {
            if (j == -1 || dp[j]) && dict[s[j + 1: i + 1]] {
                dp[i] = true
            }
        }
    }
    return dp[len(s)-1]
}
