package main

func findKthLargest(nums []int, k int) int {
	k = len(nums) - k
	start, end := 0, len(nums)-1
	for {
		index := helper(nums, start, end)
		if index > k {
			end = index - 1
		} else if index < k {
			start = index + 1
		} else {
			break
		}
	}
	return nums[k]
}

func helper(nums []int, start, end int) int {
	pvIndex := end
	pv := nums[end]
	end--
	for start <= end {
		for start <= end && nums[start] < pv {
			start++
		}
		for start <= end && nums[end] > pv {
			end--
		}
		if start <= end {
			temp := nums[start]
			nums[start] = nums[end]
			nums[end] = temp
			start++
			end--
		}
	}
	temp := nums[pvIndex]
	nums[pvIndex] = nums[start]
	nums[start] = temp
	return start
}
