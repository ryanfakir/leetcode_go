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

func combine(n int, k int) [][]int {
    return dfs(nil, nil, n, k, 1)
}


func dfs(res [][]int, tmp []int, n, k, level int) [][]int {
    if len(tmp) == k {
        el := make([]int, len(tmp))
        copy(el, tmp)
        res = append(res, el)
        return res
    }
    for i:= level; i <= n; i++ {
        tmp = append(tmp, i)
        res = dfs(res, tmp, n, k, i+1)
        tmp = tmp[:len(tmp)-1]
    }
    return res
}
