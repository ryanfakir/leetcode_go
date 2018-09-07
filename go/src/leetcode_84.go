package main

import (
	"fmt"
	"math"
)

func largestRectangleArea(heights []int) int {
	if heights == nil {
		return 0
	}
	col := len(heights)
	dp := make([][]int, col)
	for i, _ := range dp {
		dp[i] = make([]int, col)
	}
	var res int
	for i := 0; i < col; i++ {
		for j := i; j < col; j++ {
			if i == j {
				dp[i][j] = heights[i]
			} else {
				if dp[i][j-1] < heights[j] {
					dp[i][j] = dp[i][j-1]
				} else {
					dp[i][j] = heights[j]
				}
			}
			res = int(math.Max(float64(res), float64((j-i+1)*dp[i][j])))
		}
	}
	fmt.Println(heights)
	return res

}


func largestRectangleArea(heights []int) int {
    h := heights
    h = append(h, 0)
    var s []int
    var res, i int
    for i < len(h) {
        if len(s) == 0 || s[len(s)-1] <= h[i] {
            s = append(s, i)
            i++
        } else {
            var l int
            pop := s[len(s)-1]
            s = s[:len(s)-1]
            if len(s) == 0 {
                l = i
            } else {
                l = i - pop 
            }
            l = i- pop
            res = int(math.Max(float64(res), float64(l * h[pop])))
        }
    }
    return res
}
