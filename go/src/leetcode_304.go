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


type NumMatrix struct {
    m [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    m := make([][]int, len(matrix))
    for  i:= range m {
        m[i] = make([]int, len(matrix[0]))
    }
    for i := 0; i < len(matrix); i++ {
        for j:= 0; j < len(matrix[0]); j++ {
            if i == 0 && j == 0 {
                m[i][j] = matrix[0][0]
            } else if i == 0 {
                m[i][j] = m[i][j-1] + matrix[i][j] 
            } else if j == 0 {
                m[i][j] = m[i-1][j] + matrix[i][j]
            } else {
                m[i][j] = m[i-1][j] + m[i][j-1] - m[i-1][j-1] + matrix[i][j]
            }
            
        }
    }
    return NumMatrix{m}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    if row1 == 0 && col1 == 0 {
        return this.m[row2][col2]
    } else if row1 == 0 {
        return this.m[row2][col2] - this.m[row2][col1-1] 
    } else if col1 == 0 {
        return this.m[row2][col2] - this.m[row1-1][col2]
    }
    return this.m[row2][col2] - this.m[row2][col1-1] - this.m[row1-1][col2] + this.m[row1-1][col1 -1]
}

