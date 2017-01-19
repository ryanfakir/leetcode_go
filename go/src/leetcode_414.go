package main

import "math"

func thirdMax(nums []int) int {
	var max1, max2, max3 int = math.MinInt64, math.MinInt64, math.MinInt64
	for _, v := range nums {
		if v == max1 || v == max2 || v == max3 {
			continue
		}
		if v > max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}
	}
	if max3 == math.MinInt64 {
		return max1
	} else {
		return max3
	}
}
