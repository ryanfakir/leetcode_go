package main

func subsets(nums []int) [][]int {
	return helper(nums, nil, nil, 0)
}

func helper(nums []int, temp []int, res [][]int, level int) [][]int {
	el := make([]int, len(temp))
	copy(el, temp)
	res = append(res, el)
	for i := level; i < len(nums); i++ {
		temp = append(temp, nums[i])
		res = helper(nums, temp, res, i+1)
		temp = temp[:len(temp)-1]
	}
	return res
}
