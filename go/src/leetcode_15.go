package main

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := len(nums) - 1; j > 1; j-- {
			if j < len(nums)-1 && nums[j] == nums[j+1] {
				continue
			}
			k := i + 1
			for k > i && k < j {
				sum := nums[i] + nums[j] + nums[k]
				if sum == 0 {
					res = append(res, []int{nums[i], nums[k], nums[j]})
					break
				} else if sum > 0 {
					break
				} else {
					k++
				}
			}
		}
	}
	return res
}
