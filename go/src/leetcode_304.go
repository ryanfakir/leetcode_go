package main

type NumMatrix struct {
	Dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}
	row, col := len(matrix), len(matrix[0])
	sub := make([][]int, row+1)
	for i, _ := range sub {
		sub[i] = make([]int, col+1)
	}
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			sub[i][j] = sub[i-1][j] + sub[i][j-1] - sub[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{Dp: sub}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.Dp[row2+1][col2+1] - this.Dp[row2+1][col1] - this.Dp[row1][col2+1] + this.Dp[row1][col1]
}
