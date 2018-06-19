package main

func combinationSum(candidates []int, target int) [][]int {
	return helper(nil, nil, candidates, 0, target, 0)
}

func helper(res [][]int, temp, nums []int, val, target, level int) [][]int {
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
		val += nums[i]
		temp = append(temp, nums[i])
		res = helper(res, temp, nums, val, target, i)
		temp = temp[:len(temp)-1]
		val -= nums[i]
	}
	return res
}


func combinationSum(candidates []int, target int) [][]int {
    return dfs(nil, 0, target, 0, candidates, nil)
}


func dfs(res [][]int,start, target, tmp int, nums, list []int) [][]int {
    if tmp > target {
        return res
    }
    if tmp == target {
        c := make([]int, len(list))
        copy(c , list)
        res = append(res, c)
        return res
    }
    for i:= start; i < len(nums); i++ {
        list = append(list, nums[i])
        res = dfs(res, i, target, tmp + nums[i], nums, list)
        list = list[:len(list)-1]
    }
    return res
}
