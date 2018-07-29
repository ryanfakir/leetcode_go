package main

func wordBreak(s string, wordDict []string) []string {
	dp := make([]bool, len(s))
	for i, _ := range dp {
		dp[i] = true
	}
	dict := make(map[string]bool)
	for _, v := range wordDict {
		dict[v] = true
	}
	return helper(nil, dp, dict, s, 0, "")
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
				dp[i] = false
			}
			temp = org
		}
	}
	return res
}

var res []string

func wordBreak(s string, wordDict []string) []string {
    res = nil
    dict := make(map[string]bool)
    for _ , word := range wordDict {
        dict[word] = true
    }
    dp := make([]bool, len(s))
    for i := range dp {
        dp[i] = true
    }
    dfs(s, "", 0, dict, dp)
    return res
}

func dfs(s, tmp string, level int, dict map[string]bool, dp []bool) {
    if level == len(s) {
        res = append(res, tmp)
        return
    }
    if !dp[level] {return}
    size := len(res)
    for i:= level; i < len(s); i++ {
        if dict[s[level:i+1]] {
            var input string
            if len(tmp) == 0 {
                input = s[level: i+1]
            } else {
                input = tmp + " " + s[level: i+1]
            }
            dfs(s, input, i+1, dict, dp)
        }
    }
    if size == len(res) {dp[level] = false}
}


