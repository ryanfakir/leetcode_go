package main

import "math"

func insert(intervals []Interval, newInterval Interval) []Interval {
	var res []Interval
	var add bool
	for _, v := range intervals {
		if v.End < newInterval.Start {
			res = append(res, v)
		} else if v.Start > newInterval.End {
			if !add {
				res = append(res, newInterval)
			}
			res = append(res, v)
			add = true
		} else {
			newInterval.Start = int(math.Min(float64(newInterval.Start), float64(v.Start)))
			newInterval.End = int(math.Max(float64(newInterval.End), float64(v.End)))
		}
	}
	if !add {
		res = append(res, newInterval)
	}
	return res
}
