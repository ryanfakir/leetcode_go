package main

import "math"

func merge(intervals []Interval) []Interval {
	var res []Interval
	for _, v := range intervals {
		var tempRes []Interval
		for _, temp := range res {
			if v.Start > temp.End || v.End < temp.Start {
				tempRes = append(tempRes, temp)
				continue
			}
			left := math.Min(float64(v.Start), float64(temp.Start))
			right := math.Max(float64(v.End), float64(temp.End))
			v = Interval{int(left), int(right)}
		}
		tempRes = append(tempRes, v)
		res = tempRes
	}
	return res
}
