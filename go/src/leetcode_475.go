package main

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	var res int
	for _, v := range houses {
		length := find(v, heaters)
		res = int(math.Max(float64(res), float64(length)))
	}
	return res
}

func find(v int, heater []int) int {
	start, end := 0, len(heater)-1
	if start == end {
		if math.Abs(float64(v-heater[start])) > math.Abs(float64(heater[end]-v)) {
			return int(math.Abs(float64(heater[end] - v)))
		} else {
			return int(math.Abs(float64(v - heater[start])))
		}
	}
	for start+1 < end {
		mid := start + (end-start)/2
		if heater[mid] > v {
			end = mid
		} else if heater[mid] < v {
			start = mid
		} else {
			return 0
		}
	}
	if math.Abs(float64(v-heater[start])) > math.Abs(float64(heater[end]-v)) {
		return int(math.Abs(float64(heater[end] - v)))
	} else {
		return int(math.Abs(float64(v - heater[start])))
	}
}
