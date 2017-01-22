package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := helper(nums, nil, nil, 0, 0)
	return res
}

func helper(nums []int, temp []int, res [][]int, visited int, level int) [][]int {
	seg := make([]int, len(temp))
	copy(seg, temp)
	res = append(res, seg)
	m := make(map[int]bool)
	for i := level; i < len(nums); i++ {
		if !m[nums[i]] && (visited&(1<<uint(i)) == 0) {
			m[nums[i]] = true
			orginal := visited
			visited = visited | (1 << uint(i))
			temp = append(temp, nums[i])
			res = helper(nums, temp, res, visited, i+1)
			visited = orginal
			temp = temp[:len(temp)-1]
		}
	}
	return res
}
