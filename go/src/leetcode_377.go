package main

func combinationSum4(nums []int, target int) int {
	return helper(nums, target, 0, make(map[int]int))
}

func helper(nums []int, target int, val int, m map[int]int) int {
	if target == val {
		return 1
	}
	if target < val {
		return 0
	}
	if val, ok := m[val]; ok {
		return m[val]
	}
	var cnt int
	for i := 0; i < len(nums); i++ {
		val += nums[i]
		cnt += helper(nums, target, val, m)
		val -= nums[i]
	}
	m[val] = cnt
	return cnt
}
