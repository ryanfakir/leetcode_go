package main

func gameOfLife(board [][]int) {
	row, col := len(board), len(board[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			num := count(board, i, j)
			if board[i][j] == 1 {
				if num == 2 || num == 3 {
					board[i][j] = 3
				}
			} else {
				if num == 3 {
					board[i][j] = 2
				}
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			board[i][j] >>= 1
		}
	}
}

func count(board [][]int, x, y int) int {
	var res int
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i >= 0 && j >= 0 && i < len(board) && j < len(board[0]) {
				res += board[i][j] & 1
			}
		}
	}
	res -= board[x][y] & 1
	return res
}
