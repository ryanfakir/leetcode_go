package main

import "sort"

func wiggleSort(nums []int) {
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	start, end := (len(nums)+1)/2, len(nums)
	for i := 0; i < len(nums); i++ {
		if i&1 == 1 {
			end--
			nums[i] = temp[end]

		} else {
			start--
			nums[i] = temp[start]

		}
	}
}
