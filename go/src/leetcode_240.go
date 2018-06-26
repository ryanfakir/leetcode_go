package main

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	x, y := row-1, 0
	for {
		if matrix[x][y] > target {
			x--
		} else if matrix[x][y] < target {
			y++
		} else if x < 0 || y >= col {
			break
		} else {
			return true
		}
	}
	return false
}


func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {return false}
    m, n := len(matrix), len(matrix[0])
    start, end := m -1, 0
    for start >=0 && end < n {
        val := matrix[start][end]
        if val > target {
            start--
        } else if val < target {
            end++
        } else {
            return true
        }
    }
    return false
}
