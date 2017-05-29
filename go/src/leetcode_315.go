package main

func countSmaller(nums []int) []int {
	var temp []int
	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		left, right := 0, len(temp)
		for left < right {
			mid := left + (right-left)/2
			if temp[mid] > nums[i] {
				right = mid
			} else {
				left = mid + 1
			}
		}
		res[i] = right
		temp = append(temp[:right], append([]int{nums[i]}, temp[right:]...)...)
	}
	return res
}
