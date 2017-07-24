package main

import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	var res int
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			sum := nums[i] + nums[j]
			left, right := j+1, len(nums)
			for left < right {
				mid := (right-left)/2 + left
				if nums[mid] < sum {
					left = mid + 1
				} else {
					right = mid
				}
			}
			res += right - j - 1
		}
	}
	return res
}
