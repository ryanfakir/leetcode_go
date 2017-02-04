package src

import "math"

func maxProfit(prices []int) int {
	res := 0
	for i, _ := range prices {
		if i == len(prices)-1 {
			break
		}
		if prices[i+1]-prices[i] >= 0 {
			res += prices[i+1] - prices[i]
		}
	}
	return res
}
func maxProfit(prices []int) int {
	prev := math.MaxInt16
	var res int
	for _, v := range prices {
		if v > prev {
			res += v - prev
		}
		prev = v
	}
	return res
}
