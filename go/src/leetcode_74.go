package main

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	start, end := 0, row*col-1
	for start+1 < end {
		mid := (end-start)/2 + start
		x, y := mid/col, mid%col
		if matrix[x][y] < target {
			start = mid
		} else if matrix[x][y] > target {
			end = mid
		} else {
			return true
		}
	}
	if matrix[start/col][start%col] == target || matrix[end/col][end%col] == target {
		return true
	} else {
		return false
	}
}

func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {return false}
    m, n := len(matrix), len(matrix[0])
    left ,right := 0, m * n - 1
    for left <= right {
        
        mid := (left + right)/2
        if matrix[mid/n][mid%n] == target {
            return true
        } else if matrix[mid/n][mid%n] > target {
            right = mid -1
        } else {
            left = mid +1
        }
    }
    return false
}
