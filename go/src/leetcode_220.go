package main

import "math"

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	m := make(map[int]int)
	if t < 0 {
		return false
	}
	start := 0
	for i, v := range nums {
		bucket := (v - math.MinInt8) / (t + 1)
		if i-start > k {
			m[(nums[start]-math.MinInt8)/(t+1)] = 0
			start++
		}
		if (m[bucket] != 0 && int(math.Abs(float64(nums[i]-nums[m[bucket]-1]))) <= t) ||
			(m[bucket+1] != 0 && int(math.Abs(float64(nums[i]-nums[m[bucket+1]-1]))) <= t) ||
			(m[bucket-1] != 0 && int(math.Abs(float64(nums[i]-nums[m[bucket-1]-1]))) <= t) {
			return true
		}
		m[bucket] = i + 1
	}
	return false
}
