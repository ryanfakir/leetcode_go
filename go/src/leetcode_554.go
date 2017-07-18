package main

import (
	"math"
)

func leastBricks(wall [][]int) int {
	m := make(map[int]int)
	var layer int
	for _, v := range wall {
		var sum int
		for i, val := range v {
			if i == len(v)-1 {
				continue
			}
			sum += val
			m[sum]++
			layer = int(math.Max(float64(layer), float64(m[sum])))
		}
	}
	return len(wall) - layer
}
