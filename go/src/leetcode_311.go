package main

func multiply(a, b [][]int) [][]int {
	res := make([][]int, len(a))
	for i := range res {
		res[i] = make([]int, len(b[0]))
	}
	for i := 0; i < len(res); i++ {
		for k := 0; k < len(a[0]); k++ {
			if a[i][k] != 0 {
				for j := 0; j < len(res[0]); j++ {
					if b[k][j] != 0 {
						res[i][j] += a[i][k] * b[k][j]
					}
				}
			}
		}
	}
	return res
}
