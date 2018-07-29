package main

func findSubsequences(nums []int) [][]int {
	return helper(nums, 0, nil, nil)
}

func helper(nums []int, start int, arr []int, res [][]int) [][]int {
	if len(arr) >= 2 {
		temp := make([]int, len(arr))
		copy(temp, arr)
		res = append(res, temp)
	}
	m := make(map[int]bool)
	for i := start; i < len(nums); i++ {
		if (len(arr) > 0 && arr[len(arr)-1] > nums[i]) || (m[nums[i]]) {
			continue
		}
		m[nums[i]] = true
		arr = append(arr, nums[i])
		res = helper(nums, i+1, arr, res)
		arr = arr[:len(arr)-1]
	}
	return res
}

var res [][]int
func findSubsequences(nums []int) [][]int {
    res = nil
    dfs(nums, nil, 0)
    return res
}

func dfs(nums, tmp []int, level int) {
    if len(tmp) >= 2 {
        el := make([]int, len(tmp))
        copy(el, tmp)
        res = append(res, el)
    }
    if level == len(nums) {return}
    levelDict := make(map[int]bool)
    for i := level; i < len(nums); i++ {
        if  !levelDict[nums[i]] && (len(tmp) == 0 || tmp[len(tmp)-1] <= nums[i] ) {
            levelDict[nums[i]] = true
            tmp = append(tmp, nums[i])
            dfs(nums, tmp, i+1)
            tmp = tmp[:len(tmp)-1]
        }
    }
}


