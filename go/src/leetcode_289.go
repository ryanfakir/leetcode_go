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

func gameOfLife(board [][]int)  {
    b := board
    if len(b) == 0 || len(b[0]) == 0 {return}
    for i:=0; i < len(b); i++ {
        for j:=0; j< len(b[0]);j++ {
            sum := count(b, i, j)
            if sum < 2 || sum >3 { b[i][j] = b[i][j] & 1}
            if sum == 3  {
                b[i][j] = b[i][j] | 2
            }
            if sum == 2 {
                if b[i][j] == 1 {b[i][j] = 3}
            }
        }
    }
    for i:= 0; i < len(b); i++ {
        for j:=0; j< len(b[0]); j++ {
            b[i][j] >>= 1
        }
    }
}


func count(b [][]int, x, y int) int {
    var res int
    for i := x-1; i<= x+1; i++ {
        for j:= y-1; j<= y+1; j++ {
            if i < 0 || j <0 || i >= len(b) || j >= len(b[0]) { continue}
            res += b[i][j] & 1
        }
    }
    res -= b[x][y] & 1
    return res
}
