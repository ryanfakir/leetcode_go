package main

import "strconv"

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s))
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		for j := i - 1; j <= i; j++ {
			if valid(s[j : i+1]) {
				if j-1 < 0 {
					dp[i] += 1
				} else {
					dp[i] += dp[j-1]
				}
			}
		}
	}
	return dp[len(s)-1]
}

func valid(s string) bool {
	if s[0] == '0' {
		return false
	}
	val, _ := strconv.Atoi(s)
	if val > 26 {
		return false
	}
	return true
}

func numDecodings(s string) int {
    if len(s) == 0  || s[0] == '0'{
        return 0
    }
    dp := make([]int, len(s))
    dp[0] = 1
    for i:= 1; i< len(s); i++ {
        for j := i-2; j < i; j++ {
            if valid(s[j+1:i+1]) {
                if j < 0 {
                    dp[i] += 1
                } else {
                    dp[i] += dp[j]
                }
            }
        }
    }
    fmt.Println(dp)
    return dp[len(s)-1]
}

func valid(s string) bool {
    if s[0] == '0' {
        return false
    } 
    val, _ := strconv.Atoi(s)
    if val > 26 {
        return false
    }
    return true
}
