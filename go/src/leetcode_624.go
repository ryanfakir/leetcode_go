func maxDistance(arrays [][]int) int {
    var res int
    start, end := arrays[0][0], arrays[0][len(arrays[0])-1]
    for i:= 1; i < len(arrays); i++ {
        res = int(math.Max(float64(res), math.Max(math.Abs(float64(end -arrays[i][0])), math.Abs(float64(arrays[i][len(arrays[i])-1] - start)))))
        start = int(math.Min(float64(arrays[i][0]), float64(start)))
        end = int(math.Max(float64(end), float64(arrays[i][len(arrays[i]) - 1])))
    }
    return res
}
