package main

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return helper(0, nil, nil, nums)
}

func helper(visited int, temp []int, res [][]int, nums []int) [][]int {
	if len(nums) == len(temp) {
		dst := make([]int, len(temp))
		copy(dst, temp)
		res = append(res, dst)
		return res
	}
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if visited&(1<<uint(i)) == 0 && !m[nums[i]] {
			m[nums[i]] = true
			org := visited
			visited = visited | (1 << uint(i))
			temp = append(temp, nums[i])
			res = helper(visited, temp, res, nums)
			temp = temp[:len(temp)-1]
			visited = org
		}
	}
	return res
}
