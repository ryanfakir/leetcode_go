package main

func moveZeroes(nums []int) {
	var start int
	first, last := 0, len(nums)-1
	for first <= last {
		if nums[first] != 0 {
			temp := nums[first]
			nums[first] = nums[start]
			nums[start] = temp
			start++
		}
		first++
	}
}
