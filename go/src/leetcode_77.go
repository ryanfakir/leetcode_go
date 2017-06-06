package main

func combine(n int, k int) [][]int {
	return helper(nil, nil, k, n, 1)
}

func helper(res [][]int, temp []int, k, n, level int) [][]int {
	if len(temp) == k {
		arr := make([]int, len(temp))
		copy(arr, temp)
		res = append(res, arr)
		return res
	}
	for i := level; i <= n; i++ {
		temp = append(temp, i)
		res = helper(res, temp, k, n, i+1)
		temp = temp[:len(temp)-1]
	}
	return res
}
