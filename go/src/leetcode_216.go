package src

func combinationSum3(k int, n int) [][]int {
    res := make([][]int, 0)
    cur := make([]int, 0)
    dfs(&res, cur, 1, k, n)
    return res
}

func dfs(res *[][]int, cur []int, level int, k int,sum int) {
    if sum < 0 {
        return
    }
    if sum == 0 && len(cur) == k {
        newCur := make([]int, len(cur))
        copy(newCur, cur)
        *res = append(*res, newCur)
        return
    }
    for i:= level; i<=9;i++ {
        cur = append(cur, i)
        dfs(res, cur, i+1, k, sum-i)
        cur = cur[:len(cur)-1]
    }
}


var res [][]int


func combinationSum3(k int, n int) [][]int {
    res = nil
    dfs(nil, 1, k, n, 0)
    return res
}


func dfs(tmp []int, level, k,n, sum int)  {
    if len(tmp)  == k {
        if sum == n {
            arr := make([]int, len(tmp))
            copy(arr, tmp)
            res = append(res, arr)
        }
        return
    }
    for i:= level; i < 10; i++ {
        tmp = append(tmp, i)
        dfs(tmp, i+1, k,n, sum + i)
        tmp = tmp[:len(tmp)-1]
    }
    
}
