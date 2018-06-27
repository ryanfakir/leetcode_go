package main

func subsets(nums []int) [][]int {
	return helper(nums, nil, nil, 0)
}

func helper(nums []int, temp []int, res [][]int, level int) [][]int {
	el := make([]int, len(temp))
	copy(el, temp)
	res = append(res, el)
	for i := level; i < len(nums); i++ {
		temp = append(temp, nums[i])
		res = helper(nums, temp, res, i+1)
		temp = temp[:len(temp)-1]
	}
	return res
}

func subsets(nums []int) [][]int {
    return dfs(nil, nil, nums, 0)
}

func dfs(res [][]int, temp, nums []int, level int) [][]int {
    real := make([]int, len(temp))
    copy(real, temp)
    res = append(res, real)
    for i := level; i < len(nums); i++ {
        temp = append(temp, nums[i])
        res = dfs(res, temp, nums, i+1)
        temp = temp[:len(temp)-1]
    }
    return res
}
