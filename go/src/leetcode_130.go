package main

func solve(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	row, col := len(board), len(board[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 || j == 0 || i == row-1 || j == col-1 {
				if board[i][j] == 'O' {
					helper(board, i, j, row, col)
				}
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] != 'V' {
				board[i][j] = 'X'
			} else {
				board[i][j] = 'O'
			}
		}
	}
}

func helper(board [][]byte, x, y, row, col int) {
	if x < 0 || x >= row || y < 0 || y >= col || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'V'
	helper(board, x-1, y, row, col)
	helper(board, x+1, y, row, col)
	helper(board, x, y-1, row, col)
	helper(board, x, y+1, row, col)
}

func solve(board [][]byte)  {
    for i:= 0; i < len(board); i++ {
        for j:= 0; j < len(board[0]); j++ {
            if i == 0 || i == len(board)-1 || j == 0 || j == len(board[0])-1 {
                if board[i][j] == 'O' {dfs(board, i, j)}
            }
        }
    }
    for i := 0; i < len(board); i++ {
        for j:=0; j<len(board[0]); j++ {
            if board[i][j] == 'O' {
                board[i][j] = 'X'
            }
            if board[i][j] == '?' {
                board[i][j] = 'O'
            }
        }
    }
}

func dfs(board [][]byte, i, j int) {
    if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == 'X' || board[i][j] == '?'{return}
    board[i][j] = '?'
    dfs(board, i-1, j)
    dfs(board, i+1, j)
    dfs(board, i, j-1)
    dfs(board, i, j+1)
}
