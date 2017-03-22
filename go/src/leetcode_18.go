package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := len(nums) - 1; j > 2; j-- {
			if j < len(nums)-1 && nums[j] == nums[j+1] {
				continue
			}
			left, right := i+1, j-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					res = append(res, []int{nums[i], nums[left], nums[right], nums[j]})
					left++
					right--
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if sum < target {
					left++
				} else if sum > target {
					right--
				}
			}
		}
	}
	return res
}
