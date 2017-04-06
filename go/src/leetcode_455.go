package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var res, childIndex int
	for _, v := range s {
		if v >= g[childIndex] {
			res++
			childIndex++
			if childIndex >= len(g) {
				break
			}
		}
	}
	return res
}
