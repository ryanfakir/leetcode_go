package main

func twoSum(nums []int, target int) []int {
	var res []int
	m := make(map[int]int)
	for k, v := range nums {
		if val, ok := m[target-v]; ok {
			res = append(res, []int{k, val}...)
			return res
		}
		m[v] = k
	}
	return res 
}
