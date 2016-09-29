package src

func getRow(rowIndex int) []int {
    res := make([]int, 0)
    curlvl := 0
    for curlvl <= rowIndex {
        cur := make([]int,0)
        for j :=0; j<curlvl+1;j++ {
            if j==0 || j==curlvl {
                cur = append(cur, 1)
            } else {
                cur = append(cur, res[j-1]+res[j])
            }
        }
        res = cur
        curlvl++
    }
    return res
}
