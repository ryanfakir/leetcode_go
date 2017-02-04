package src

import "math"

func maxProfit(prices []int) int {
	min := 0
	max := 0
	for i, e := range prices {
		if i == 0 {
			min = e
			continue
		}
		min = int(math.Min(float64(min), float64(e)))
		max = int(math.Max(float64(max), float64(e-min)))
	}
	return max
}

func maxProfit(prices []int) int {
	min := math.MaxInt16
	max := 0
	for _, e := range prices {
		min = int(math.Min(float64(min), float64(e)))
		max = int(math.Max(float64(max), float64(e-min)))
	}
	return max
}
