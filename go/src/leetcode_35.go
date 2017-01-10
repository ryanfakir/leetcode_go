package src

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)/2
		if target > nums[mid] {
			left = mid
		} else if target < nums[mid] {
			right = mid
		} else {
			return mid
		}
	}
	if nums[left] == target || nums[right] == target {
		if nums[left] == target {
			return left
		} else {
			return right
		}
	} else if nums[left] > target {
		return left
	} else if nums[right] < target {
		return right + 1
	} else {
		return right
	}

}
