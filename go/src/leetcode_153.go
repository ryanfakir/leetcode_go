package main

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[left] && nums[mid] > nums[right] {
			left = mid
		} else if nums[mid] < nums[left] && nums[mid] < nums[right] {
			right = mid
		} else {
			return nums[left]
		}
	}
	if nums[left] < nums[right] {
		return nums[left]
	} else {
		return nums[right]
	}
}
