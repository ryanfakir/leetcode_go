package main

import "math"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	var res int
	for left < right {
		var h int
		w := right - left
		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}
		res = int(math.Max(float64(h*w), float64(res)))
	}
	return res
}
