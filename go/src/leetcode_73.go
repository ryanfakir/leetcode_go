func setZeroes(matrix [][]int)  {
    m, n := len(matrix), len(matrix[0])
    var h, v bool
    for i:=0 ; i < m; i++ {
        if matrix[i][0] == 0 {
            v = true
        }
    }
    for i:=0 ; i < n; i++ {
        if matrix[0][i] == 0 {
            h = true
        }
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }
    for i:=1 ; i < m; i++ {
        if matrix[i][0] == 0 {
            var j int
            for j < n {
                matrix[i][j] = 0
                j++
            }
        }
    }
    for i:=1 ; i < n; i++ {
        if matrix[0][i] == 0 {
            var j int
            for j < m {
                matrix[j][i] = 0
                j++
            }
        }
    }
    if h {
        for i:=0 ; i < n; i++ {
            matrix[0][i] = 0
        }
    }
    if v {
        for i:=0 ; i < m; i++ {
            matrix[i][0] = 0
        }
    }
}
