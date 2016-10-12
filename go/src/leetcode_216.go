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