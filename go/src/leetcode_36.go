package src

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		matrix := make([]bool, 10)
		for j := 0; j < 9; j++ {
			if !checker(matrix, board[i][j]) {
				return false
			}
		}
	}
	for i := 0; i < 9; i++ {
		matrix := make([]bool, 10)
		for j := 0; j < 9; j++ {
			if !checker(matrix, board[j][i]) {
				return false
			}
		}
	}
	for i := 0; i < 9; i = i + 3 {
		for j := 0; j < 9; j = j + 3 {
			matrix := make([]bool, 10)
			for k := 0; k < 9; k++ {
				if !checker(matrix, board[i+k/3][j+k%3]) {
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
	if matrix[index] {
		return false
	} else {
		matrix[index] = true
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		col := make(map[int]bool)
		row := make(map[int]bool)
		cube := make(map[int]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if col[int(board[i][j]-'0')] {
					return false
				}
				col[int(board[i][j]-'0')] = true
			}
			if board[j][i] != '.' {
				if row[int(board[j][i]-'0')] {
					return false
				}
				row[int(board[j][i]-'0')] = true
			}
			if board[i/3*3+j/3][i%3*3+j%3] != '.' {
				if cube[int(board[i/3*3+j/3][i%3*3+j%3])] {
					return false
				}
				cube[int(board[i/3*3+j/3][i%3*3+j%3])] = true
			}
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
    if len(board) == 0 || len(board[0]) == 0 {
        return false
    }
    for i := 0 ; i < len(board); i++ {
        h, v, c := make(map[byte]bool), make(map[byte]bool), make(map[byte]bool)
        for j:= 0; j < len(board[0]); j++ {
            if board[i][j] != byte('.') {
                if !h[board[i][j]] {
                    h[board[i][j]] = true
                } else {
                    return false
                }
            }
            if (board[j][i]) != byte('.') {
                if !v[board[j][i]] {
                    v[board[j][i]] = true
                } else {
                    return false
                }
            }
            if (board[i/3 * 3 + j/3][i%3 * 3 + j%3]) != byte('.') {
                if !c[board[i/3 * 3 + j/3][i%3 * 3 + j%3]] {
                    c[board[i/3 * 3 + j/3][i%3 * 3 + j%3]] = true
                } else {
                    return false
                }
            }
            
        }
    }
    return true
}
