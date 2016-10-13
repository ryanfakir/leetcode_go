package src

func numTrees(n int) int {
    res := make([]int, n+1)
    res[0], res[1] = 1, 1
    for i := 2; i<=n; i++ {
        for j:= 0; j<i; j++ {
            res[i] += res[j]*res[i-j-1]
        }
    }
    return res[n]
}