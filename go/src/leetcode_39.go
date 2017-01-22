package main

func combinationSum(candidates []int, target int) [][]int {
	return helper(nil, nil, candidates, 0, target, 0)
}

func helper(res [][]int, temp, nums []int, val, target, level int) [][]int {
	if val > target {
		return res
	}
	if val == target {
		el := make([]int, len(temp))
		copy(el, temp)
		res = append(res, el)
		return res
	}
	for i := level; i < len(nums); i++ {
		val += nums[i]
		temp = append(temp, nums[i])
		res = helper(res, temp, nums, val, target, i)
		temp = temp[:len(temp)-1]
		val -= nums[i]
	}
	return res
}
