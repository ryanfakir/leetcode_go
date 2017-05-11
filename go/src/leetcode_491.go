package main

func findSubsequences(nums []int) [][]int {
	return helper(nums, 0, nil, nil)
}

func helper(nums []int, start int, arr []int, res [][]int) [][]int {
	if len(arr) >= 2 {
		temp := make([]int, len(arr))
		copy(temp, arr)
		res = append(res, temp)
	}
	m := make(map[int]bool)
	for i := start; i < len(nums); i++ {
		if (len(arr) > 0 && arr[len(arr)-1] > nums[i]) || (m[nums[i]]) {
			continue
		}
		m[nums[i]] = true
		arr = append(arr, nums[i])
		res = helper(nums, i+1, arr, res)
		arr = arr[:len(arr)-1]
	}
	return res
}
