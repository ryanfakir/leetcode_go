func generateMatrix(n int) [][]int {
    if n == 0 {return nil}
    res := make([][]int, n)
    for i := range res {
        res[i] = make([]int, n)
    }
    xs, xe, ys, ye := 0, len(res)-1, 0, len(res[0])-1
    num := 1
    for xs <= xe && ys <= ye {
        i, j := xs, ys
        for j <= ye {
            res[i][j] = num
            num++
            j++
        }
        xs++
        i, j = xs, ye
        for i <= xe  {
            res[i][j] = num
            num++
            i++
        }
        ye--
        i, j = xe, ye
        for j >= ys && ys <= ye && xs <= xe {
            res[i][j] = num
            num++
            j--
        }
        xe--
        i, j = xe, ys
        for i >= xs && xs <= xe && ys <= ye{
            res[i][j] = num
            num++
            i--
        }
        ys++
    }
    return res
}
