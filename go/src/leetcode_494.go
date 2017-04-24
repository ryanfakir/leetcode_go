package main

func findTargetSumWays(nums []int, S int) int {
	return helper(nums, 0, S)
}

func helper(nums []int, level, sum int) int {
	if level == len(nums) {
		if sum == 0 {
			return 1
		}
		return 0
	}
	var res int
	res += helper(nums, level+1, sum-nums[level])
	res += helper(nums, level+1, sum+nums[level])
	return res
}
