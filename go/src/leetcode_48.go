func rotate(matrix [][]int)  {
    for i, j := 0, len(matrix)-1; i < j; i , j  = i+1, j - 1 {
        tmp := matrix[j]
        matrix[j] = matrix[i]
        matrix[i] = tmp
    }
    for i:= 0; i < len(matrix); i++ {
        for j := i +1; j < len(matrix[0]); j++ {
            tmp := matrix[i][j]
            matrix[i][j] = matrix[j][i]
            matrix[j][i] = tmp
        }
    }
}
