package main

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
