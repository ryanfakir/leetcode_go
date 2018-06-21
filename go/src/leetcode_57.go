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


func insert(intervals []Interval, newInterval Interval) []Interval {
    var res []Interval
    start, end := newInterval.Start, newInterval.End
    var add bool
    for _, v := range intervals {
        if add {
            res = append(res, v)
            continue
        }
        if start > v.End {
            res = append(res, v)
            continue
        }
        if end < v.Start {
            add = true
            res = append(res, Interval{start, end})
            res = append(res, v)
            continue
        }
        start = int(math.Min(float64(start), float64(v.Start)))
        end = int(math.Max(float64(end), float64(v.End)))
    }
    if !add {
        res = append(res, Interval{start, end})
    } 
    return res
}
