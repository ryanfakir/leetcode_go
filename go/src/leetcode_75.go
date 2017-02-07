package main

func sortColors(nums []int) {
	var start, end = 0, len(nums) - 1
	for i := 0; i <= end; i++ {
		if nums[i] == 0 {
			nums[i] = nums[start]
			nums[start] = 0
			start++
		} else if nums[i] == 2 {
			nums[i] = nums[end]
			nums[end] = 2
			end--
			i--
		}
	}
}
