package main

func combinationSum2(candidates []int, target int) [][]int {
	return helper(nil, nil, target, 0, candidates, 0)
}

func helper(res [][]int, temp []int, target int, val int, nums []int, level int) [][]int {
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
		if i > level && nums[i] == nums[i-1] {
			continue
		}
		temp = append(temp, nums[i])
		val += nums[i]
		res = helper(res, temp, target, val, nums, i+1)
		temp = temp[:len(temp)-1]
		val -= nums[i]
	}
	return res
}
