package src

func isValidSudoku(board [][]byte) bool {
    for i:=0; i<9; i++ {
        matrix := make([]bool, 10)
        for j:=0;j<9; j++ {
            if !checker(matrix, board[i][j]) {
                return false
            }
        }
    }
    for i:=0; i<9; i++ {
        matrix := make([]bool, 10)
        for j:=0; j<9;j++ {
            if !checker(matrix, board[j][i]) {
                return false
            }
        }
    }
    for i:=0; i<9; i=i+3 {
        for j:=0; j<9; j=j+3 {
            matrix := make([]bool, 10)
           for k:=0;k<9;k++ {
               if !checker(matrix, board[i+k/3][j + k%3]) {
                   return false
               }
           }
        }
    }
    return true
}

func checker(matrix []bool, element byte) bool {
    if element == '.' {
        return true
    }
    index := element - '0'
    if matrix[index]  {
        return false
    } else {
        matrix[index] = true
    }
    return true
}
