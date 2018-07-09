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

func findMin(nums []int) int {
    if len(nums) == 0 {return 0}
    left, right := 0, len(nums)-1
    res := nums[0]
    for left <= right {
        var tmp int
        mid := (left + right)/2
        if nums[mid] < nums[left] {
            right = mid
            tmp = nums[mid]
        } else if nums[mid] > nums[right] {
            left = mid + 1
            tmp = nums[right]
        } else {
            tmp = nums[left]
            left++
        }
        res = int(math.Min(float64(res), float64(tmp)))
    }
    return res
}
