package src

func permute(nums []int) [][]int {
	return dfs(nums, 0, nil, nil)
}

func dfs(nums []int, visited int, res [][]int, temp []int) [][]int {
	if len(temp) == len(nums) {
		dict := make([]int, len(temp))
		copy(dict, temp)
		res = append(res, dict)
		return res
	}
	for i := 0; i < len(nums); i++ {
		if visited&(1<<uint(i)) == 0 {
			// fmt.Println(visited)
			original := visited
			visited = visited | (1 << uint(i))
			temp = append(temp, nums[i])
			res = dfs(nums, visited, res, temp)
			visited = original
			temp = temp[:len(temp)-1]
		}
	}
	return res
}

func permute(nums []int) [][]int {
    return dfs(nil, nil, nums, make(map[int]bool))
}

func dfs(res [][]int, tmp, nums []int, dict map[int]bool) [][]int {
    if len(tmp) == len(nums) {
        coper := make([]int, len(tmp))
        copy(coper, tmp)
        res = append(res, coper)
        return res
    }
    for i := 0; i < len(nums); i++ {
        if dict[nums[i]] {continue}
        dict[nums[i]] = true
        tmp = append(tmp, nums[i])
        res = dfs(res, tmp, nums, dict)
        tmp = tmp[:len(tmp)-1]
        dict[nums[i]] = false
    }
    return res
}
