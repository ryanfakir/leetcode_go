package main

func findCircleNum(M [][]int) int {
	row := len(M)
	var res int
	for i := 0; i < row; i++ {
		if M[i][i] == 1 {
			res++
			helper(M, i)
		}

	}
	return res
}

func helper(M [][]int, x int) {
	col := len(M[0])
	M[x][x] = 0
	for i := 0; i < col; i++ {
		if M[x][i] == 1 {
			M[x][i] = 0
			M[i][x] = 0
			helper(M, i)
		}
	}
}
