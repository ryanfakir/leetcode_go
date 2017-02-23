package main

import "sort"

func meetingRoom(arr []Interval) bool {
	sort.Sort(Col(arr))
	for i := 1; i < len(arr); i++ {
		if arr[i].Start < arr[i-1].End {
			return false
		}
	}
	return true
}

type Col []Interval

func (c Col) Len() int           { return len(c) }
func (c Col) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Col) Less(i, j int) bool { return c[i].Start < c[j].End }
