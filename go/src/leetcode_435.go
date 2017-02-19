package main

import "sort"

func eraseOverlapIntervals(intervals []Interval) int {
	sort.Sort(Col(intervals))
	var last, res int
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start < intervals[last].End {
			res++
			if intervals[last].End > intervals[i].End {
				last = i
			}
		} else {
			last = i
		}
	}
	return res
}

type Col []Interval

func (c Col) Len() int           { return len(c) }
func (c Col) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Col) Less(i, j int) bool { return c[i].Start < c[j].Start }
