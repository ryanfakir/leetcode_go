package main

import "math"

func findMaxLength(nums []int) int {
	var sum, res int
	dict := make(map[int]int)
	dict[0] = -1
	for i, v := range nums {
		if v == 0 {
			sum += -1
		} else {
			sum += 1
		}
		if v, ok := dict[sum]; !ok {
			dict[sum] = i
		} else {
			res = int(math.Max(float64(res), float64(i-v)))
		}
	}
	return res
}
