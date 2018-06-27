package main

func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 {
		return false
	}
	row, col := len(board), len(board[0])
	visited := make([][]bool, row)
	for i, _ := range visited {
		visited[i] = make([]bool, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if helper(board, visited, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func helper(board [][]byte, visited [][]bool, word string, x, y, index int) bool {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || visited[x][y] {
		return false
	}
	visited[x][y] = true
	if word[index] == board[x][y] {
		index++
		if index == len(word) {
			return true
		}
		if helper(board, visited, word, x+1, y, index) || helper(board, visited, word, x-1, y, index) || helper(board, visited, word, x, y+1, index) || helper(board, visited, word, x, y-1, index) {
			return true
		}
	}
	// reset visited
	visited[x][y] = false
	return false
}

func exist(board [][]byte, word string) bool {
    row, col := len(board), len(board[0])
    dict := make([][]bool, row)
    for i := range dict {
        dict[i] = make([]bool, col)
    }
    for i:=0; i< row; i++ {
        for j:=0; j < col; j++ {
            if dfs(board, dict, word, i, j, 0) {return true}
        }
    }
    return false
}

func dfs(board [][]byte, visited [][]bool, word string, x, y, index int) bool {
    if index == len(word) {return true}
    row, col := len(board), len(board[0])
    if x < 0 || x >= row || y < 0 || y >= col || visited[x][y] || word[index] != board[x][y] {return false}
    visited[x][y] = true
    res := dfs(board, visited, word, x+1 , y, index +1 ) || dfs(board, visited, word, x-1 , y, index+1)  || dfs(board, visited, word, x , y+1, index+1)  || dfs(board, visited, word, x , y-1, index +1) 
    visited[x][y] =false
    return res
}
