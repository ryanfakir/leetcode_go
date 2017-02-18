package main

func longestIncreasingPath(matrix [][]int) int {
	if matrix == nil || len(matrix) == 0 {
		return 0
	}
	row, col := len(matrix), len(matrix[0])
	dp := make([][]int, row)
	for i, _ := range dp {
		dp[i] = make([]int, col)
	}
	var res int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			local := helper(matrix, dp, i, j)
			if local > res {
				res = local
			}
		}
	}
	return res
}

func helper(matrix [][]int, dp [][]int, i, j int) int {
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	var dir = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, 1}, []int{0, -1}}
	local := 1
	for _, v := range dir {
		x, y := i+v[0], j+v[1]
		if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[0]) || matrix[i][j] >= matrix[x][y] {
			continue
		}
		max := helper(matrix, dp, x, y) + 1
		if local < max {
			local = max
		}
	}
	dp[i][j] = local
	return local
}
