package main

func updateMatrix(matrix [][]int) [][]int {
	row, col := len(matrix), len(matrix[0])
	res, dp := make([][]int, row), make([][]int, row)
	for i, _ := range res {
		res[i] = make([]int, col)
	}
	for i, _ := range dp {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			dp[i][j] = -1
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] != 0 {
				res[i][j] = bfs(matrix, i, j, row, col)
			}
		}
	}
	return res
}

func bfs(matrix [][]int, i, j, row, col int) int {
	var queue [][]int
	visited := make([][]bool, row)
	for o, _ := range visited {
		visited[o] = make([]bool, col)
	}
	queue = append(queue, []int{i, j})
	var level int
	for len(queue) > 0 {
		num := len(queue)
		for i := 0; i < num; i++ {
			pop := queue[0]
			queue = queue[1:]
			if matrix[pop[0]][pop[1]] == 0 {
				return level
			}
			visited[pop[0]][pop[1]] = true
			if valid(pop[0]+1, pop[1], row, col) && !visited[pop[0]+1][pop[1]] {
				queue = append(queue, []int{pop[0] + 1, pop[1]})
			}
			if valid(pop[0]-1, pop[1], row, col) && !visited[pop[0]-1][pop[1]] {
				queue = append(queue, []int{pop[0] - 1, pop[1]})
			}
			if valid(pop[0], pop[1]+1, row, col) && !visited[pop[0]][pop[1]+1] {
				queue = append(queue, []int{pop[0], pop[1] + 1})
			}
			if valid(pop[0], pop[1]-1, row, col) && !visited[pop[0]][pop[1]-1] {
				queue = append(queue, []int{pop[0], pop[1] - 1})
			}
		}
		level++
	}
	return level
}

func valid(i, j, row, col int) bool {
	if i < 0 || j < 0 || i >= row || j >= col {
		return false
	}
	return true
}
