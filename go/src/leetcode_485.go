package main

import "math"

func findMaxConsecutiveOnes(nums []int) int {
	var cnt, res int
	for _, v := range nums {
		if v == 0 {
			res = int(math.Max(float64(res), float64(cnt)))
			cnt = 0
			continue
		}
		cnt++
		res = int(math.Max(float64(res), float64(cnt)))
	}
	return res
}
