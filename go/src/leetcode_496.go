package main

func nextGreaterElement(findNums []int, nums []int) []int {
	dict := make(map[int]int)
	for i, v := range nums {
		dict[v] = i
	}
	res := make([]int, len(findNums))
	for i, v := range findNums {
		start := dict[v]
		res[i] = -1
		for start = start + 1; start < len(nums); start++ {
			if nums[start] > v {
				res[i] = nums[start]
				break
			}
		}
	}
	return res
}
