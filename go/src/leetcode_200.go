package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	row := len(grid)
	col := len(grid[0])
	visited := make([][]bool, row)
	for k, _ := range visited {
		visited[k] = make([]bool, col)
	}
	var count int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				helper(i, j, visited, row, col, grid)
				count++
			}
		}
	}
	return count
}

func helper(i int, j int, visited [][]bool, row int, col int, arr [][]byte) {
	if i < 0 || j < 0 || i >= row || j >= col || arr[i][j] == '0' || visited[i][j] {
		return
	}
	visited[i][j] = true
	helper(i+1, j, visited, row, col, arr)
	helper(i-1, j, visited, row, col, arr)
	helper(i, j-1, visited, row, col, arr)
	helper(i, j+1, visited, row, col, arr)
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	row := len(grid)
	col := len(grid[0])
	// visited := make([][]bool, row)
	// for k, _ := range visited {
	//     visited[k] = make([]bool, col)
	// }
	var count int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				helper(i, j, row, col, grid)
				count++
			}
		}
	}
	return count
}

func helper(i int, j int, row int, col int, arr [][]byte) {
	if i < 0 || j < 0 || i >= row || j >= col || arr[i][j] == '0' {
		return
	}
	// visited[i][j] = true
	arr[i][j] = '0'
	helper(i+1, j, row, col, arr)
	helper(i-1, j, row, col, arr)
	helper(i, j-1, row, col, arr)
	helper(i, j+1, row, col, arr)
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	var res int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				helper(grid, i, j)
				res++
			}
		}
	}
	return res
}

func helper(grid [][]byte, x, y int) {
	var dir = [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}
	grid[x][y] = '0'
	for _, v := range dir {
		a, b := x+v[0], y+v[1]
		if a < 0 || b < 0 || a >= len(grid) || b >= len(grid[0]) || grid[a][b] == '0' {
			continue
		}
		helper(grid, a, b)
	}
}
