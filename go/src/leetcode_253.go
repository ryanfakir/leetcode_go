package main

import (
	"math"
	"sort"
)

func minMeetingRooms(input []Interval) int {
	sort.Sort(Col(input))
	var temp []Interval
	for _, v := range input {
		var node Interval
		var mid []Interval
		for _, tmp := range temp {
			if tmp.Start < v.Start && v.Start < tmp.End {
				node = Interval{int(math.Min(float64(tmp.Start), float64(v.Start), int(math.Max(float64(tmp.End), float64(v.End)))))}
			} else {
				mid = append(mid, tmp)
			}
		}
		temp = append(temp, node)
		temp = append(temp, mid...)

	}
	return len(temp)

}

type Col []Interval

func (c Col) Len() int      { return len(c) }
func (c Col) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c Col) Less(i, j int) { return c[i].Start < c[j].Start && c[i].End < c[j].End }
