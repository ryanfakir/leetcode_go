package main

func minPatches(nums []int, n int) int {
	var res, i int
	miss := 1
	for miss <= n {
		if i < len(nums) && nums[i] <= miss {
			miss += nums[i]
			i++
		} else {
			miss += miss
			res++
		}
	}
	return res
}
