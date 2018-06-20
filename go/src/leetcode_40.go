package main

func combinationSum2(candidates []int, target int) [][]int {
	return helper(nil, nil, target, 0, candidates, 0)
}

func helper(res [][]int, temp []int, target int, val int, nums []int, level int) [][]int {
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
		if i > level && nums[i] == nums[i-1] {
			continue
		}
		temp = append(temp, nums[i])
		val += nums[i]
		res = helper(res, temp, target, val, nums, i+1)
		temp = temp[:len(temp)-1]
		val -= nums[i]
	}
	return res
}


func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    return dfs(nil, candidates, nil, target, 0, 0)
}


func dfs(res [][]int, nums, list []int, target, tmp, start int) [][]int {{
    if tmp > target {
        return res
    }
    if tmp == target {
        tList := make([]int, len(list))
        copy(tList, list)
        res = append(res, tList)
        return res
    }
    dedup := make(map[int]bool)
    for i := start; i < len(nums); i++ {
        if dedup[nums[i]] {continue}
        dedup[nums[i]] = true
        list = append(list, nums[i])
        res = dfs(res, nums, list, target, tmp + nums[i], i+1)
        list = list[:len(list)-1]
    }
    return res
}}
