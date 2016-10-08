package src

func countBits(num int) []int {
    res := make([]int, num+1)
    for i:=1; i<num+1; i++ {
        res[i] = res[i&(i-1)] + 1
    }
    return res
}
