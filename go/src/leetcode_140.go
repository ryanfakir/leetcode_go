package main

import "fmt"

func wordBreak(s string, wordDict []string) []string {
	dp := make([]bool, len(s))
	for i, _ := range dp {
		dp[i] = true
	}
	dict := make(map[string]bool)
	for _, v := range wordDict {
		dict[v] = true
	}
	fmt.Println(dict)
	res := helper(nil, dp, dict, s, 0, "")
	fmt.Println(dp)
	return res
}

func helper(res []string, dp []bool, wordDict map[string]bool, s string, level int, temp string) []string {
	if level == len(s) {
		res = append(res, temp)
		return res
	}
	if !dp[level] {
		return res
	}
	for i := level + 1; i <= len(s); i++ {
		if wordDict[s[level:i]] {
			size := len(res)
			org := temp
			if len(temp) == 0 {
				temp += s[level:i]
			} else {
				temp += " " + s[level:i]
			}
			res = helper(res, dp, wordDict, s, i, temp)
			if len(res) == size {
				dp[i-1] = false
			}
			temp = org
		}
	}
	return res
}
