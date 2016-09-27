package src

func generate(numRows int) [][]int {
    res := make([][]int, 0)
    for i:=0; i<numRows;i++ {
        level := make([]int, 0)
        start := 0
        for start < i + 1 {
            if start == 0 {
                level = append(level, 1)
            } else if start == i {
                level = append(level, 1)
            } else {
                prev := res[i-1]
                level = append(level, prev[start-1] + prev[start])
            }
            start++
        }
        res = append(res, level)
    }
    return res
}
