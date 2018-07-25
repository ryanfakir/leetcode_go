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

var res int 

func longestIncreasingPath(matrix [][]int) int {
    res = 1
    if len(matrix) == 0 || len(matrix[0]) == 0 {return 0}
    row, col := len(matrix), len(matrix[0])
    dp := make([][]int, row)
    for i:= range dp {
        dp[i] = make([]int, col)
    }
    for i:= 0; i < row; i++ {
        for j:=0; j< col; j++ {
            res = int(math.Max(float64(res), float64(dfs(matrix, dp, -1 << 31, i, j))))
        }
    }
    return res
}

func dfs(matrix,dp [][]int, val, x, y int) int {
    if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[0]) || val >= matrix[x][y] {return 0}
    if dp[x][y] != 0 {return dp[x][y]}
    dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    var local = 1
    for _ ,v  := range dir {
        local = int(math.Max(float64(local), float64(1 + dfs(matrix, dp, matrix[x][y], x + v[0], y + v[1]))))
    }
    dp[x][y] = local
    return local
} 
