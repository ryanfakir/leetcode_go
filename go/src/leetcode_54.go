func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {return nil}
    xs, xe, ys, ye := 0, len(matrix)-1, 0, len(matrix[0])-1
    var res []int
    fmt.Println(xs, xe, ys, ye)
    for xs <= xe && ys <= ye {
        i, j := xs, ys
        for j <= ye {
            fmt.Println("1[", i, ",", j, "]" )
            res = append(res, matrix[i][j])
            j++
        }
        xs++
        i, j = xs, ye
        for i <= xe  {
            fmt.Println("2[", i, ",", j, "]" )
            res =append(res, matrix[i][j])
            i++
        }
        ye--
        i, j = xe, ye
        for j >= ys && ys <= ye && xs <= xe {
            fmt.Println("3[", i, ",", j, "]" )
            res = append(res, matrix[i][j])
            j--
        }
        xe--
        i, j = xe, ys
        for i >= xs && xs <= xe && ys <= ye{
            fmt.Println("4[", i, ",", j, "]" )
            res = append(res, matrix[i][j])
            i--
        }
        ys++
    }
    return res
}
